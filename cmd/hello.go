package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "hello",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a color argument")
		}
		if args[0] == "world" {
			return nil
		}
		return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0], "is a valid color")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
