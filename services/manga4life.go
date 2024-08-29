package services

import (
	"fmt"
	"time"

	"github.com/DinnieJ/godash"
	"github.com/tebeka/selenium"
)

type Manga4LifeParserService struct {
	ParserStruct
}

func (m *Manga4LifeParserService) InitInstance() {
	if driver, service, err := NewDriver(); err == nil {
		m.Driver = driver
		m.Service = service
	} else {
		panic(err)
	}
}

func (m *Manga4LifeParserService) GetInfo(url string) *MangaInfo {
	if err := m.Driver.Get(url); err != nil {
		panic(err)
	}
	waitLoadErr := m.Driver.WaitWithTimeoutAndInterval(func(wd selenium.WebDriver) (bool, error) {
		if _, err := wd.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div"); err != nil {
			return false, err
		}
		return true, nil
	}, 30*time.Second, 5*time.Second)
	if waitLoadErr != nil {
		panic(waitLoadErr)
	}
	info := &MangaInfo{}
	infoBoxEl := godash.Must(m.Driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div:nth-child(1) > div.col-md-9.col-sm-8.top-5 > ul"))
	listInfoEls := godash.Must(infoBoxEl.FindElements(selenium.ByTagName, "li"))
	infoElementByLabel := make(map[string]selenium.WebElement)
	for _, element := range listInfoEls {
		if labelEl, err := element.FindElement(selenium.ByClassName, "mlabel"); err == nil {
			infoElementByLabel[godash.Must(labelEl.Text())] = element
		}
	}
	if nameEl, err := m.Driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div:nth-child(1) > div.col-md-9.col-sm-8.top-5 > ul > li:nth-child(1)"); err == nil {
		if name, err := nameEl.Text(); err == nil {
			info.Name = name
		}
	}

	if alterNameEl := infoElementByLabel["Alternate Name(s):"]; alterNameEl != nil {
		if alterName, err := alterNameEl.Text(); err == nil {
			info.Name += fmt.Sprintf(" (%s)", alterName[len("Alternate Name(s): "):])
		}
	}
	if descriptionEl, err := infoElementByLabel["Description:"].FindElement(selenium.ByTagName, "div"); err == nil {
		if description, err := descriptionEl.Text(); err == nil {
			info.Description = description
		}
	}

	if authorTagsEl, err := infoElementByLabel["Author(s):"].FindElements(selenium.ByTagName, "a"); err == nil {
		for _, e := range authorTagsEl {
			if authorName, err := e.Text(); err == nil {
				info.Authors = append(info.Authors, authorName)
			}
		}
	}

	if genreTagsEl, err := infoElementByLabel["Genre(s):"].FindElements(selenium.ByTagName, "a"); err == nil {
		for _, e := range genreTagsEl {
			if g, err := e.Text(); err == nil {
				info.Genres = append(info.Genres, g)
			}
		}
	}

	if thumbnailEl, err := m.Driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div:nth-child(1) > div.col-md-3.col-sm-4.col-3.top-5 > img"); err == nil {
		thumbnailUrl, _ := thumbnailEl.GetAttribute("src")
		info.Thumbnail = thumbnailUrl
	}

	if showAllChapterEl, err := m.Driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div.list-group.top-10.bottom-5.ng-scope > div"); err == nil {
		showAllChapterEl.Click()
	}

	if chapterBoxEl, err := m.Driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div.list-group.top-10.bottom-5.ng-scope"); err == nil {
		listChapterEls, _ := chapterBoxEl.FindElements(selenium.ByTagName, "a")
		info.NumberOfChapter = int32(len(listChapterEls))
	}

	return info
}

func (m *Manga4LifeParserService) GetListChapter(url string) []Chapter {
	if err := m.Driver.Get(url); err != nil {
		panic(err)
	}

	waitLoadErr := m.Driver.WaitWithTimeoutAndInterval(func(wd selenium.WebDriver) (bool, error) {
		if _, err := wd.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div"); err != nil {
			return false, err
		}
		return true, nil
	}, 30*time.Second, 5*time.Second)
	if waitLoadErr != nil {
		panic(waitLoadErr)
	}

	if showAllChapterEl, err := m.Driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div.list-group.top-10.bottom-5.ng-scope > div"); err == nil {
		showAllChapterEl.Click()
	}

	chapters := []Chapter{}
	const ROOT_URL string = "https://manga4life.com"
	if chapterBoxEl, err := m.Driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div.list-group.top-10.bottom-5.ng-scope"); err == nil {
		listChapterEls, _ := chapterBoxEl.FindElements(selenium.ByTagName, "a")
		for i := len(listChapterEls) - 1; i >= 0; i-- {
			chapterUrl, _ := listChapterEls[i].GetAttribute("href")
			chapterNameEl, _ := listChapterEls[i].FindElement(selenium.ByTagName, "span")
			chapterName, _ := chapterNameEl.Text()
			chapters = append(chapters, Chapter{Index: int32(len(listChapterEls) - i + 1), Name: chapterName, ChapterUrl: fmt.Sprintf("%s%s", ROOT_URL, chapterUrl)})
		}
	}
	return chapters
}

func (m *Manga4LifeParserService) ParseData(url string, start int32, end int32) *BookData {
	return nil
}

func (m *Manga4LifeParserService) KillService() error {
	return m.Service.Stop()
}
