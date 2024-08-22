package cmd

import "github.com/spf13/cobra"

var m_StartIdx, m_EndIdx int

var ParseCommand = &cobra.Command{
	Use:   "parse",
	Short: "Parsing books to JSON format",
	Long:  "Parsing books to JSON format",
	Run: func(cmd *cobra.Command, args []string) {
		g_Module.ParseData(f_url, int32(m_StartIdx), int32(m_EndIdx))
	},
}

func init() {
	ParseCommand.Flags().IntVar(&m_StartIdx, "start", -1, "Start chapter")
	ParseCommand.Flags().IntVar(&m_EndIdx, "end", -1, "End chapter")
}
