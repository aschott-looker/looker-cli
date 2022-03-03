package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "looker-cli",
	Short: "Command line interface for interacting with a Looker instance.",
	Long: "Command line interface for interacting with a Looker instance.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
