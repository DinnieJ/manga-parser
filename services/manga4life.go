package services

import "github.com/tebeka/selenium"

type Manga4LifeParserService struct {
	driver selenium.WebDriver
}

func (m *Manga4LifeParserService) InitInstance() {
	if driver, err := NewDriver(); err == nil {
		m.driver = driver
	} else {
		panic(err)
	}
}

func (m *Manga4LifeParserService) GetInfo(url string) *MangaInfo {
	if err := m.driver.Get(url); err != nil {
		panic(err)
	}
	return &MangaInfo{}
}

func (m *Manga4LifeParserService) GetListChapter(url string) []Chapter {
	return []Chapter{}
}

func (m *Manga4LifeParserService) ParseData(url string, start int32, end int32) *BookData {
	return nil
}
