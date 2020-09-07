package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

var pulsarURL = "pulsar://192.168.0.11:6650"

func main() {
	// create pulsar client.
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               pulsarURL,
		ConnectionTimeout: 20 * time.Second,
		OperationTimeout:  10 * time.Second,
		// Max number of connections to a single broker that will kept in the pool. (Default: 1 connection)
		MaxConnectionsPerBroker: 10,
	})

	if err != nil {
		log.Println("create client error: ", err)
		return
	}

	defer client.Close()

	reader, err := client.CreateReader(pulsar.ReaderOptions{
		Topic:          "my-topic",
		StartMessageID: pulsar.EarliestMessageID(),
	})

	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	for reader.HasNext() {
		msg, err := reader.Next(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n", msg.ID(), string(msg.Payload()))
	}

}
