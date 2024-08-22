package cmd

import (
	"fmt"
	"reflect"

	"github.com/DinnieJ/Manga-Parser/services"
	"github.com/spf13/cobra"
)

var enabledModule = map[string]bool{
	"manga4life": true,
	"mangadex":   false,
}

var (
	f_module string
	f_url    string
)

var modules = map[string]services.ParserService{
	"manga4life":     &services.Manga4LifeParserService{},
	"myreadingmanga": nil,
}

var g_Module services.ParserService

var rootCmd = &cobra.Command{
	Use:   "manga-parser",
	Short: "Manga data parser services",
	Long:  `Manga parser service module, used to parse a manga url to JSON file`,
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if f_module == "" {
			return fmt.Errorf("module not provided\nSupported modules: %s", reflect.ValueOf(modules).MapKeys())
		}
		if !enabledModule[f_module] {
			return fmt.Errorf("module not supported\nSupported modules: %s", reflect.ValueOf(modules).MapKeys())
		}
		g_Module = modules[f_module]
		g_Module.InitInstance()
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fucked")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func parseCommandPreRun(rootCmd *cobra.Command, cmds []*cobra.Command) {
	for _, element := range cmds {
		element.PreRunE = rootCmd.PreRunE
		rootCmd.AddCommand(element)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&f_url, "url", "u", "", "Manga url string")
	rootCmd.PersistentFlags().StringVarP(&f_module, "module", "m", "", "Parser module")
	rootCmd.MarkPersistentFlagRequired("url")
	rootCmd.MarkPersistentFlagRequired("module")
	rootCmd.CompletionOptions = cobra.CompletionOptions{DisableDefaultCmd: true}
	parseCommandPreRun(rootCmd, []*cobra.Command{ListCommand, ParseCommand, InfoCommand})

}
