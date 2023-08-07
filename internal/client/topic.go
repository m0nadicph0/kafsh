package client

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/m0nadicph0/kafsh/internal/models"
	"log"
)

func (k *kafkaClient) ListTopics() ([]*models.Topic, error) {

	result := make([]*models.Topic, 0)

	admin, err := sarama.NewClusterAdmin(k.brokers, k.config)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := admin.Close()
		if err != nil {
			log.Println("Error closing client:", err)
		}
	}()

	topics, err := admin.ListTopics()
	if err != nil {
		return nil, err
	}

	for name, topic := range topics {
		result = append(result, &models.Topic{
			Name:       name,
			Partitions: topic.NumPartitions,
			Replicas:   topic.ReplicationFactor,
		})
	}

	return result, nil
}

func (k *kafkaClient) CreateTopic(name string, partitions int32, replication int16) error {
	admin, err := sarama.NewClusterAdmin(k.brokers, k.config)
	if err != nil {
		return fmt.Errorf("failed to create admin client: %v", err)
	}
	defer func() {
		if err := admin.Close(); err != nil {
			log.Printf("Error closing admin client: %v", err)
		}
	}()
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     partitions,
		ReplicationFactor: replication,
	}

	err = admin.CreateTopic(name, topicDetail, false)
	if err != nil {
		return fmt.Errorf("failed to create topic %s: %v", name, err)
	}
	return nil
}
