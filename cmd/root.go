package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tamaAPI",
	Short: "An API to find the tamas around you",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the root command. Use 'app help' for more information.")
	},
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    // Add subcommands here
    rootCmd.AddCommand(radiusCmd)
}
