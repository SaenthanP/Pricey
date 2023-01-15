package asyncmessaging

import (
	"encoding/json"
	"fmt"
	"linkservice/model"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type AsyncMessageClient struct {
	producer      *kafka.Producer
	delivery_chan chan kafka.Event
}

func NewAsyncMessageClient() *AsyncMessageClient {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "linkservice",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Producer could not be created: %v", err)
	}
	delivery_chan := make(chan kafka.Event, 10000)

	return &AsyncMessageClient{p, delivery_chan}
}

func (messageClient AsyncMessageClient) CallScrape(linkToScrape model.Link) {
	topic := "scrape"
	linkJson, err := json.Marshal(linkToScrape)
	if err != nil {
		log.Fatal(err)
	}

	err = messageClient.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          linkJson},
		messageClient.delivery_chan,
	)

	if err != nil {
		log.Fatalf("Something went wrong trying to send scrape message %v", err)
	}

	log.Println("Successfully published message")
}
