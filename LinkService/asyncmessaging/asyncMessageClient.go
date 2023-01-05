package asyncmessaging

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type AsyncMessageClient struct {
	producer *kafka.Producer
}

func NewAsyncMessageClient() *AsyncMessageClient {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "linkservice",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Producer could not be created: %v", err)
	}
	return &AsyncMessageClient{p}
}

func (messageClient AsyncMessageClient) CallScrape() {
	topic := "test"
	delivery_chan := make(chan kafka.Event, 10000)

	err := messageClient.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          []byte("val")},
		delivery_chan,
	)

	if err != nil {
		log.Fatalf("Something went wrong trying to send scrape message %v", err)
	}

	log.Println("Successfully published message")
}
