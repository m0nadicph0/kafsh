package cmd

import (
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/printer"
	"github.com/spf13/cobra"
	"os"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Describe and List nodes",
}

var nodeLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List nodes in a cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		kCli := client.NewKafkaClient(getKafkaConfig(), []string{"localhost:9092"})
		nodes, err := kCli.ListNodes()
		if err != nil {
			return err
		}

		printer.NewPrinter(printer.TabPrinter, os.Stdout).PrintNodes(nodes)
		return nil
	},
}

func init() {
	nodeCmd.AddCommand(nodeLsCmd)
	rootCmd.AddCommand(nodeCmd)
}
