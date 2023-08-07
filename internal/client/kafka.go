package client

import (
	"github.com/Shopify/sarama"
	"github.com/m0nadicph0/kafsh/internal/models"
)

type KafkaClient interface {
	ListTopics() ([]*models.Topic, error)
	CreateTopic(name string, partitions int32, replication int16) error
	DeleteTopic(name string) error
	DescribeTopic(name string) (*models.TopicDesc, error)
}

type kafkaClient struct {
	config  *sarama.Config
	brokers []string
}

func NewKafkaClient(config *sarama.Config, brokers []string) KafkaClient {
	return &kafkaClient{config: config, brokers: brokers}
}
