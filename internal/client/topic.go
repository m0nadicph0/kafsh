package client

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/m0nadicph0/kafsh/internal/models"
	"log"
	"strings"
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

func (k *kafkaClient) DeleteTopic(name string) error {
	admin, err := sarama.NewClusterAdmin(k.brokers, k.config)
	if err != nil {
		return fmt.Errorf("failed to create admin client: %v", err)
	}
	defer func() {
		if err := admin.Close(); err != nil {
			log.Printf("Error closing admin client: %v", err)
		}
	}()

	err = admin.DeleteTopic(name)
	if err != nil {
		return fmt.Errorf("failed to delete topic %s: %v", name, err)
	}

	return nil
}

func (k *kafkaClient) DescribeTopic(name string) (*models.TopicDesc, error) {
	admin, err := sarama.NewClusterAdmin(k.brokers, k.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create admin client: %v", err)
	}
	defer func() {
		if err := admin.Close(); err != nil {
			log.Printf("Error closing admin client: %v", err)
		}
	}()
	result, err := admin.DescribeTopics([]string{name})
	if err != nil {
		return nil, err
	}

	metadata := result[0]

	cfg, err := admin.DescribeConfig(sarama.ConfigResource{
		Type: sarama.TopicResource,
		Name: name,
	})

	return &models.TopicDesc{
		Name:       metadata.Name,
		Internal:   metadata.IsInternal,
		Compacted:  isCompacted(cfg),
		Partitions: getPartitionDetails(metadata),
		Config:     getPartitionConfigs(cfg),
	}, nil
}

func getPartitionConfigs(cfg []sarama.ConfigEntry) []*models.PartitionConfig {
	result := make([]*models.PartitionConfig, 0)
	for _, entry := range cfg {
		result = append(result, &models.PartitionConfig{
			Name:      entry.Name,
			Value:     entry.Value,
			ReadOnly:  entry.ReadOnly,
			Sensitive: entry.Sensitive,
		})
	}
	return result
}

func getPartitionDetails(metadata *sarama.TopicMetadata) []*models.PartitionDetail {
	partitionDetails := make([]*models.PartitionDetail, 0)
	for _, p := range metadata.Partitions {
		partitionDetails = append(partitionDetails, &models.PartitionDetail{
			Partition: p.ID,
			Leader:    p.Leader,
			Replicas:  p.Replicas,
			ISR:       p.Isr,
		})
	}
	return partitionDetails
}

func isCompacted(cfg []sarama.ConfigEntry) bool {
	for _, e := range cfg {
		if e.Name == "cleanup.policy" {
			for _, setting := range strings.Split(e.Value, ",") {
				if setting == "compact" {
					return true
				}
			}
		}
	}
	return false
}
