package cmd

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"net/http"
	"os"

	"github.com/dolmen-go/kittyimg"
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

var InfoCommand = &cobra.Command{
	Use:   "info",
	Short: "Get manga information from url",
	Long:  "Get manga information from url",
	Run: func(cmd *cobra.Command, args []string) {
		info := g_Module.GetInfo(f_url)
		displayImgTerm(info.Thumbnail)
		fmt.Println(info)
	},
}
