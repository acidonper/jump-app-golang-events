package event

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	kafka "github.com/segmentio/kafka-go"
    "github.com/segmentio/ksuid"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func Consumer() {
	// get kafka reader using environment variables.
	kafkaURL := os.Getenv("KAFKA_HOST")
	topic := os.Getenv("KAFKA_TOPIC")
    groupID := ksuid.New()

	reader := getKafkaReader(kafkaURL, topic, groupID.String())

	defer reader.Close()

	fmt.Printf("Start consuming - Kafka server: %s , Topic: %s & GroupID: %s\n", kafkaURL, topic, groupID)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Message at topic:%v partition:%v offset:%v %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
