package util

import "github.com/m0nadicph0/kafsh/internal/models"

func GetTopicNames(topics []*models.Topic) []string {
	result := make([]string, 0)
	for _, topic := range topics {
		result = append(result, topic.Name)
	}
	return result
}
