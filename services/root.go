package services

import (
	"fmt"

	"github.com/DinnieJ/godash"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type MangaInfo struct {
	Thumbnail       string
	NumberOfChapter int32
	Name            string
	Description     string
	Authors         []string
}

type Chapter struct {
	Index      int32
	Name       string
	ChapterUrl int32
	TotalPage  int32
	Pages      []Page
}

type Page struct {
	Index         int32
	ImageUrl      string
	PageImageData []byte
}

type BookData struct {
	Info     *MangaInfo
	Chapters []Chapter
}

type ParserService interface {
	InitInstance()
	GetInfo(url string) *MangaInfo
	GetListChapter(url string) []Chapter
	ParseData(url string, start int32, end int32) *BookData
}

func StartChromeDriverService(driverSrc string, driverPort int) *selenium.Service {
	driverSrcDef := godash.If(driverSrc != "", driverSrc, "chromedriver-linux64/chromedriver")
	driverPortDef := godash.If(driverPort != 0, driverPort, 4444)
	service, err := selenium.NewChromeDriverService(driverSrcDef, driverPortDef)
	if err != nil {
		panic(err)
	}
	return service
}

func NewDriver() (selenium.WebDriver, *selenium.Service, error) {
	service := StartChromeDriverService("./chromedriver", 4444)
	caps := selenium.Capabilities{}
	chromeCaps := chrome.Capabilities{Args: []string{"--headless"}}
	// chromeCaps.AddExtension("")
	caps.AddChrome(chromeCaps)

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	if err := driver.MaximizeWindow(""); err != nil {
		return nil, nil, err
	}

	return driver, service, nil
}
