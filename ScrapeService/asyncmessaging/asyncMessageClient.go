package asyncmessaging

import (
	"fmt"
	"log"
	"os"
	"scrapeservice/model"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var _consumer *kafka.Consumer
var _workerPool *model.WorkerPool

func NewAsyncMessageClient(workerPool *model.WorkerPool) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest"})

	if err != nil {
		log.Fatalf("Consumer could not be created: %v", err)
	}

	err = consumer.Subscribe("test", nil)

	if err != nil {
		log.Fatalf("Consumer could not be created: %v", err)
	}

	_consumer = consumer
	_workerPool = workerPool

	log.Println("Started consumer, and subscribed to topic: test")

	go ConsumerPolling()
}

func ConsumerPolling() {
	log.Println("Polling Topic")

	for {
		ev := _consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Printf("RECEIVED %s\n", string(e.Value))
			job := model.Job{JobType: string(e.Value), Executor: nil}
			_workerPool.Jobs <- job
			// application-specific processing
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)

		}
	}
}
