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
		toggle, _ := cmd.Flags().GetBool("toggle")
		fmt.Println("Toggle set to ", toggle)
		queryId, _ := cmd.Flags().GetInt64("query_id")
		fmt.Println("query_id set to: ", queryId)
		fields, _ := cmd.Flags().GetString("fields")
		fmt.Println("fields set to: ", fields)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	helloCmd.Flags().Int64("query_id", 0, "Id of query")
	cobra.MarkFlagRequired(helloCmd.Flags(), "query_id")
	helloCmd.Flags().String("fields", "", "Requested fields.")
}
