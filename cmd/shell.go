package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start an interactive shell with completions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shell called")
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
