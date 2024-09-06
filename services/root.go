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
	Genres          []string
}

type Chapter struct {
	Index      int32  `json:"-"`
	Name       string `json:"chapterTitle"`
	ChapterUrl string `json:"chapterUrl"`
	TotalPage  int32  `json:"totalPage"`
	Pages      []*Page
}

type Page struct {
	Index         int32  `json:"index"`
	ImageUrl      string `json:"imageUrl"`
	PageImageData []byte `json:"-"`
}

type BookData struct {
	Info     *MangaInfo
	Chapters []Chapter
}

type BookDataParseJson struct {
	Title    string    `json:"title"`
	Cover    string    `json:"cover"`
	Authors  []string  `json:"authors"`
	Chapters []Chapter `json:"chapters"`
}

type ParserStruct struct {
	Service *selenium.Service
	Driver  selenium.WebDriver
}

type ParserService interface {
	InitInstance()
	KillService() error
	GetInfo(url string) *MangaInfo
	GetListChapter(url string) []Chapter
	ParseData(url string, start int32, end int32) *BookDataParseJson
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
	chromeCaps := chrome.Capabilities{Args: []string{
		"--blink-settings=imagesEnabled=false",
		"--headless",
		"--no-sandbox",
		"--disable-dev-shm-usage"},
	}
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
