package cmd

import (
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/printer"
	"os"

	"github.com/spf13/cobra"
)

var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List groups",
	RunE: func(cmd *cobra.Command, args []string) error {
		kCli := client.NewKafkaClient(getKafkaConfig(), []string{"localhost:9092"})
		groups, err := kCli.ListGroups()
		if err != nil {
			return err
		}

		printer.NewPrinter(printer.TabPrinter, os.Stdout).PrintGroups(groups)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(groupsCmd)
}
