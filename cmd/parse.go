package cmd

import "github.com/spf13/cobra"

var ParseCommand = &cobra.Command{
	Use:   "parse",
	Short: "Parsing books to JSON format",
	Long:  "Parsing books to JSON format",
	Run:   func(cmd *cobra.Command, args []string) {},
}
