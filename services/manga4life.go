package services

import (
	"github.com/tebeka/selenium"
)

type Manga4LifeParserService struct {
	service *selenium.Service
	driver  selenium.WebDriver
}

func (m *Manga4LifeParserService) InitInstance() {
	if driver, service, err := NewDriver(); err == nil {
		m.driver = driver
		m.service = service
	} else {
		panic(err)
	}
}

func (m *Manga4LifeParserService) GetInfo(url string) *MangaInfo {
	defer m.service.Stop()
	if err := m.driver.Get(url); err != nil {
		panic(err)
	}
	info := &MangaInfo{}
	if nameEl, err := m.driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div:nth-child(1) > div.col-md-9.col-sm-8.top-5 > ul > li:nth-child(1)"); err == nil {
		if name, err := nameEl.Text(); err == nil {
			info.Name = name
		}
	}

	if descriptionEl, err := m.driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div:nth-child(1) > div.col-md-9.col-sm-8.top-5 > ul > li:nth-child(11) > div"); err == nil {
		if description, err := descriptionEl.Text(); err == nil {
			info.Description = description
		}
	}

	if authorsEl, err := m.driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div:nth-child(1) > div.col-md-9.col-sm-8.top-5 > ul > li:nth-child(4)"); err == nil {
		authorTagsEl, _ := authorsEl.FindElements(selenium.ByTagName, "a")
		for _, e := range authorTagsEl {
			if authorName, err := e.Text(); err == nil {
				info.Authors = append(info.Authors, authorName)
			}
		}
	}

	if thumbnailEl, err := m.driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div:nth-child(1) > div.col-md-3.col-sm-4.col-3.top-5 > img"); err == nil {
		thumbnailUrl, _ := thumbnailEl.GetAttribute("src")
		info.Thumbnail = thumbnailUrl
	}

	if showAllChapterEl, err := m.driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div.list-group.top-10.bottom-5.ng-scope > div"); err == nil {
		showAllChapterEl.Click()
	}

	if chapterBoxEl, err := m.driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div.list-group.top-10.bottom-5.ng-scope"); err == nil {
		listChapterEls, _ := chapterBoxEl.FindElements(selenium.ByTagName, "a")
		info.NumberOfChapter = int32(len(listChapterEls))
	}

	return info
}

func (m *Manga4LifeParserService) GetListChapter(url string) []Chapter {
	return []Chapter{}
}

func (m *Manga4LifeParserService) ParseData(url string, start int32, end int32) *BookData {
	return nil
}
