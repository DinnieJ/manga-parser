package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:   "list-chapter",
	Short: "Show all chapter of manga",
	Long:  "Show all chapter of manga",
	Run: func(cmd *cobra.Command, args []string) {
		defer g_Module.KillService()
		chapters := g_Module.GetListChapter(f_url)
		if len(chapters) == 0 {
			fmt.Println("No chapters found")
		}
		for idx, chapter := range chapters {
			customPrintLn(strconv.Itoa(idx)+": ", chapter.Name)
		}
	},
}
