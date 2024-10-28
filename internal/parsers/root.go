package services

import (
	"fmt"

	"github.com/DinnieJ/Manga-Parser/pkg/types"
	"github.com/DinnieJ/godash"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type ParserStruct struct {
	Service *selenium.Service
	Driver  selenium.WebDriver
}

type ParserService interface {
	InitInstance()
	KillService() error
	GetInfo(url string) *types.MangaInfo
	GetListChapter(url string) []types.Chapter
	ParseData(url string, start int32, end int32) *types.BookDataParseJson
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
	service := StartChromeDriverService("./assets/chromedriver", 4444)
	caps := selenium.Capabilities{}
	chromeCaps := chrome.Capabilities{Args: []string{
		"--blink-settings=imagesEnabled=false",
		// "--headless",
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
