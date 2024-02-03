package app

import (
	"log"
	"strings"

	"github.com/IBM/sarama"
)

func StartConsumer() {
	consumer, err := sarama.NewConsumer(strings.Split("localhost:9092", ","), sarama.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	partition, err := consumer.ConsumePartition("topic1", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range partition.Messages() {
		log.Printf("Consumed message offset %s\n", msg.Value)
	}
}

func StartProductor() (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(strings.Split("localhost:9092", ","), nil)
	if err != nil {
		return nil, err
	}
	return producer, nil
}
