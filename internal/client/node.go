package client

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/m0nadicph0/kafsh/internal/models"
	"log"
)

func (k *kafkaClient) ListNodes() ([]*models.Node, error) {
	admin, err := sarama.NewClusterAdmin(k.brokers, k.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create admin client: %v", err)
	}
	defer func() {
		if err := admin.Close(); err != nil {
			log.Printf("Error closing admin client: %v", err)
		}
	}()

	brokers, ctlrID, err := admin.DescribeCluster()
	if err != nil {
		return nil, fmt.Errorf("failed to describe cluster: %v", err)
	}

	nodes := make([]*models.Node, 0)

	for _, broker := range brokers {
		nodes = append(nodes, &models.Node{
			ID:           broker.ID(),
			Address:      broker.Addr(),
			IsController: broker.ID() == ctlrID,
		})
	}

	return nodes, nil
}
