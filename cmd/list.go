package cmd

import (
	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:   "list-chapter",
	Short: "Show all chapter of manga",
	Long:  "Show all chapter of manga",
	Run: func(cmd *cobra.Command, args []string) {
		g_Module.GetListChapter(f_url)
	},
}
