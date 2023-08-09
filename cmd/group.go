package cmd

import (
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/printer"
	"github.com/spf13/cobra"
	"os"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Display information about consumer groups",
}

var groupLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List groups",
	RunE:  groupsCmd.RunE,
}

var groupDescCmd = &cobra.Command{
	Use:   "describe GROUP",
	Short: "Describe consumer group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		kCli := client.NewKafkaClient(getKafkaConfig(), []string{"localhost:9092"})
		group, err := kCli.DescribeGroup(args[0])
		if err != nil {
			return err
		}

		printer.NewPrinter(printer.TabPrinter, os.Stdout).PrintGroupDescription(group)
		return nil
	},
}

func init() {
	groupCmd.AddCommand(groupLsCmd)
	groupCmd.AddCommand(groupDescCmd)
	rootCmd.AddCommand(groupCmd)
}
