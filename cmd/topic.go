package cmd

import (
	"fmt"
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/printer"
	"github.com/spf13/cobra"
	"os"
)

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Create and describe topics",
}

var lsCmd = &cobra.Command{
	Use:   "ls",
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

/*
add-config  Add config key/value pair to topic
  delete      Delete a topic
  describe    Describe topic
*/

var createCmd = &cobra.Command{
	Use:   "create TOPIC",
	Short: "Create a topic",

	RunE: func(cmd *cobra.Command, args []string) error {
		partitions, err := cmd.Flags().GetInt32("partitions")
		if err != nil {
			return err
		}
		replicas, err := cmd.Flags().GetInt16("replicas")
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return fmt.Errorf("insufficient arguments: topic name")
		}

		kCli := client.NewKafkaClient(getKafkaConfig(), []string{"localhost:9092"})
		err = kCli.CreateTopic(args[0], partitions, replicas)
		if err != nil {
			return err
		}

		fmt.Printf("created topic %s with partitions=%d, replicas=%d\n", args[0], partitions, replicas)
		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a topic",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	topicCmd.AddCommand(lsCmd)

	createCmd.Flags().Int32P("partitions", "p", 1, "Number of partitions")
	createCmd.Flags().Int16P("replicas", "r", 1, "Number of replicas")
	topicCmd.AddCommand(createCmd)

	rootCmd.AddCommand(topicCmd)

}
