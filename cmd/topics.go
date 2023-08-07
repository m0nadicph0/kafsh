package cmd

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/printer"
	"github.com/spf13/cobra"
	"os"
)

var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "List topics",
	Run: func(cmd *cobra.Command, args []string) {
		kCli := client.NewKafkaClient(getKafkaConfig(), []string{"localhost:9092"})
		topics, err := kCli.ListTopics()
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			return
		}

		printer.NewPrinter(printer.TabPrinter, os.Stdout).PrintTopics(topics)

	},
}

func getKafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.ClientID = "kafsh"
	return config
}

func init() {
	rootCmd.AddCommand(topicsCmd)
}
