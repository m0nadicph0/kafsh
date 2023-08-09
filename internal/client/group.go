package client

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/m0nadicph0/kafsh/internal/models"
	"log"
)

func (k *kafkaClient) ListGroups() ([]*models.Group, error) {
	admin, err := sarama.NewClusterAdmin(k.brokers, k.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create admin client: %v", err)
	}
	defer func() {
		if err := admin.Close(); err != nil {
			log.Printf("Error closing admin client: %v", err)
		}
	}()

	groups, err := admin.ListConsumerGroups()
	if err != nil {
		return nil, fmt.Errorf("failed to describe cluster: %v", err)
	}
	groupNames := getGroupNames(groups)

	groupDescs, err := admin.DescribeConsumerGroups(groupNames)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Group, 0)

	for _, desc := range groupDescs {
		result = append(result, &models.Group{
			Name:      desc.GroupId,
			State:     desc.State,
			Consumers: len(desc.Members),
		})
	}

	return result, nil
}

func getGroupNames(m map[string]string) []string {
	result := make([]string, 0)
	for name := range m {
		result = append(result, name)
	}
	return result
}
