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

var Modules = map[string]services.ParserService{
	"manga4life": &services.Manga4LifeParserService{},
}

var rootCmd = &cobra.Command{
	Use:   "manga-parser",
	Short: "Manga data parser services",
	Long:  `Manga parser service module, used to parse a manga url to JSON file`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		moduleFlag := cmd.PersistentFlags().Lookup("module")
		if moduleFlag == nil {
			return fmt.Errorf("module not provided\nSupported modules: %s", reflect.ValueOf(Modules).MapKeys())
		}
		if !enabledModule[moduleFlag.Value.String()] {
			return fmt.Errorf("module not supported\nSupported modules: %s", reflect.ValueOf(Modules).MapKeys())
		}
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
	rootCmd.PersistentFlags().StringP("url", "u", "", "Manga url string")
	rootCmd.PersistentFlags().StringP("module", "m", "", "Parser module")
	rootCmd.MarkPersistentFlagRequired("url")
	rootCmd.MarkPersistentFlagRequired("module")
	parseCommandPreRun(rootCmd, []*cobra.Command{ListCommand, ParseCommand, InfoCommand})
}
