package cmd

import (
	"github.com/spf13/cobra"
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

func init() {
	groupCmd.AddCommand(groupLsCmd)
	rootCmd.AddCommand(groupCmd)
}
