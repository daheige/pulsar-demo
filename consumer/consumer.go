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

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})

	if err != nil {
		log.Println("create consumer error: ", err)
		return
	}

	defer consumer.Close()

	var msg pulsar.Message
	for {
		msg, err = consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// receive msg
		fmt.Printf("Received message msgId: %#v -- content: '%s'\n", msg.ID(), string(msg.Payload()))

		// ack message
		consumer.Ack(msg)
	}
}

/**
INFO[0000] Connecting to broker                          remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] TCP connection established                    local_addr="192.168.0.8:54288" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connection is ready                           local_addr="192.168.0.8:54288" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connecting to broker                          remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] TCP connection established                    local_addr="192.168.0.8:54289" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connection is ready                           local_addr="192.168.0.8:54289" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connecting to broker                          remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] TCP connection established                    local_addr="192.168.0.8:54290" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connection is ready                           local_addr="192.168.0.8:54290" remote_addr="pulsar://192.168.0.11:6650"
INFO[0000] Connected consumer                            consumerID=1 name=gcdll subscription=my-sub topic="persistent://public/default/my-topic"
INFO[0000] Created consumer                              consumerID=1 name=gcdll subscription=my-sub topic="persistent://public/default/my-topic"
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:110, entryID:0, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc00024c000), receivedTime:time.Time{wall:0xbfcd49312c751e68, ext:174329244, loc:(*time.Location)(0x4cc7fa0)}} -- content: 'hello'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:110, entryID:1, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc00024c000), receivedTime:time.Time{wall:0xbfcd49312c757070, ext:174350384, loc:(*time.Location)(0x4cc7fa0)}} -- content: 'hello'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:110, entryID:2, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc00024c000), receivedTime:time.Time{wall:0xbfcd49312c758010, ext:174354939, loc:(*time.Location)(0x4cc7fa0)}} -- content: 'hello'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:110, entryID:4, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc00024c000), receivedTime:time.Time{wall:0xbfcd49399a4ba4c8, ext:33869442107, loc:(*time.Location)(0x4cc7fa0)}} -- content: 'hello'
*/
