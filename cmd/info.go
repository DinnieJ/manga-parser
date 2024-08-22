package cmd

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dolmen-go/kittyimg"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func displayImgTerm(url string) {
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != 200 {
		return
	}

	defer resp.Body.Close()
	// bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)
	m, _, err := image.Decode(resp.Body)
	if err != nil {
		return
	}

	if err := kittyimg.Fprint(os.Stdout, m); err != nil {
		panic(err)
	}

}

func customPrintLn(title string, value string) {
	titleSprint := color.New(color.FgCyan, color.Bold).SprintfFunc()
	valueSprint := color.New(color.FgHiMagenta).SprintfFunc()

	fmt.Printf("%s%s\n", titleSprint("%s: ", title), valueSprint(value))
}

var InfoCommand = &cobra.Command{
	Use:   "info",
	Short: "Get manga information from url",
	Long:  "Get manga information from url",
	Run: func(cmd *cobra.Command, args []string) {
		info := g_Module.GetInfo(f_url)
		displayImgTerm(info.Thumbnail)
		customPrintLn("Title", info.Name)
		customPrintLn("Authors", strings.Join(info.Authors, ", "))
		customPrintLn("Description", info.Description)
		customPrintLn("Total chapters", strconv.FormatInt(int64(info.NumberOfChapter), 10))
	},
}
