package cmd

import "github.com/spf13/cobra"

var InfoCommand = &cobra.Command{
	Use:   "info",
	Short: "Get manga information from url",
	Long:  "Get manga information from url",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
