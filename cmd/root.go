package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafsh",
	Short: "Kafka command-line client",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	brokers    string
	cluster    string
	configFile string
	verbose    bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&brokers, "brokers", "b", "", "comma separated list of broker ip:port pairs")
	rootCmd.PersistentFlags().StringVarP(&cluster, "cluster", "c", "", "set a temporary current cluster")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.kafsh/config)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")
}
