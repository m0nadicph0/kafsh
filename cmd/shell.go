package cmd

import (
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/shell"
	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start an interactive shell with completions",
	RunE: func(cmd *cobra.Command, args []string) error {
		return shell.Start(client.NewKafkaClient(getKafkaConfig(), []string{"localhost:9092"}))
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}
