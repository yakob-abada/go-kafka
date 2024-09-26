package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"strconv"
)

const KafkaServerAddress = "localhost:9092"
const KafkaTopic = "notifications"

func createProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{KafkaServerAddress}, config)

	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}
	return producer, nil
}

func main() {
	producer, err := createProducer()
	if err != nil {
		log.Fatal(err)
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(1)),
		Value: sarama.StringEncoder("notification_text"),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return
	}
}
