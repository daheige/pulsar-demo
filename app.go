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

	// create a producer.
	var producer pulsar.Producer
	producer, err = client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})

	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}

	fmt.Println("Published message")
}

/**
INFO[0000] Connecting to broker                          remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] TCP connection established                    local_addr="192.168.0.8:54344" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connection is ready                           local_addr="192.168.0.8:54344" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connecting to broker                          remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] TCP connection established                    local_addr="192.168.0.8:54345" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connection is ready                           local_addr="192.168.0.8:54345" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connecting to broker                          remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] TCP connection established                    local_addr="192.168.0.8:54346" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connection is ready                           local_addr="192.168.0.8:54346" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Created producer                              cnx="192.168.0.8:54346 -> 192.168.0.11:6650" producerID=1 producer_name=standalone-0-9 topic="persistent://public/default/my-topic"
Published message
INFO[0000] Closing producer                              producerID=1 producer_name=standalone-0-9 topic="persistent://public/default/my-topic"
INFO[0000] Closed producer                               producerID=1 producer_name=standalone-0-9 topic="persistent://public/default/my-topic"
*/
