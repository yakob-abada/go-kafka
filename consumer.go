package main

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

const ConsumerGroup = "notifications-group"
const KafkaServerAdd = "localhost:9092"
const KafkaTopic1 = "notifications"

type NotificationStore struct {
}

func (c NotificationStore) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c NotificationStore) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c NotificationStore) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("message: %s", message.Value)

		session.MarkMessage(message, "")
	}

	return nil
}

func createConsumerGroup() (sarama.ConsumerGroup, error) {
	c, err := sarama.NewConsumerGroup([]string{KafkaServerAdd}, ConsumerGroup, sarama.NewConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return c, nil
}

func main() {
	fmt.Printf("Kafka CONSUMER (Group: %s) 👥📥 ", ConsumerGroup)

	consumer, err := createConsumerGroup()
	if err != nil {
		log.Fatalf("failed to create consumer group: %v", err)
	}

	defer consumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	err = consumer.Consume(ctx, []string{KafkaTopic1}, NotificationStore{})
	defer cancel()

	if err != nil {
		log.Fatalf("failed to start consumer group: %v", err)
	}
}
