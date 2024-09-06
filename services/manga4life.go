package services

import (
	"fmt"
	"strings"
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
	var infoBoxEl selenium.WebElement
	waitLoadErr := m.Driver.WaitWithTimeoutAndInterval(func(wd selenium.WebDriver) (bool, error) {
		if el, err := wd.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div:nth-child(1) > div.col-md-9.col-sm-8.top-5 > ul"); err != nil {
			return false, nil
		} else {
			infoBoxEl = el
			return true, nil
		}

	}, 30*time.Second, 5*time.Second)
	if waitLoadErr != nil || infoBoxEl == nil {
		panic(waitLoadErr)
	}
	info := &MangaInfo{}
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
		if _, err := wd.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div.list-group.top-10.bottom-5.ng-scope"); err != nil {
			return false, nil
		}
		return true, nil
	}, 30*time.Second, 3*time.Second)
	if waitLoadErr != nil {
		panic(waitLoadErr)
	}

	if showAllChapterEl, err := m.Driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div.list-group.top-10.bottom-5.ng-scope > div"); err == nil {
		showAllChapterEl.Click()
	}

	chapters := []Chapter{}
	if chapterBoxEl, err := m.Driver.FindElement(selenium.ByCSSSelector, "body > div.container.MainContainer > div > div > div > div > div.list-group.top-10.bottom-5.ng-scope"); err == nil {
		listChapterEls, _ := chapterBoxEl.FindElements(selenium.ByTagName, "a")
		for i := len(listChapterEls) - 1; i >= 0; i-- {
			chapterUrl, _ := listChapterEls[i].GetAttribute("href")
			chapterNameEl, _ := listChapterEls[i].FindElement(selenium.ByTagName, "span")
			chapterName, _ := chapterNameEl.Text()
			chapters = append(chapters, Chapter{Index: int32(len(listChapterEls) - i), Name: chapterName, ChapterUrl: fmt.Sprintf("%s%s", "", chapterUrl)})
		}
	}
	return chapters
}

func (m *Manga4LifeParserService) getChapterPages(url string) ([]*Page, error) {
	listPages := []*Page{}
	if err := m.Driver.Get(url); err != nil {
		return listPages, fmt.Errorf("get chapter %s failed", url)
	}

	waitLoadErr := m.Driver.WaitWithTimeoutAndInterval(func(wd selenium.WebDriver) (bool, error) {
		if el, err := wd.FindElement(selenium.ByCSSSelector, "body > div.MainContainer > div:nth-child(1) > div"); err != nil {
			return false, nil
		} else {
			if listButtonBoxElements, err := el.FindElements(selenium.ByCSSSelector, ".Column.col-lg-2.col-6"); err == nil {
				if len(listButtonBoxElements) == 4 {
					if fullPageBtn, err := listButtonBoxElements[2].FindElement(selenium.ByTagName, "button"); err == nil {
						if txt, _ := listButtonBoxElements[2].Text(); strings.ToLower(strings.TrimSpace(txt)) == "long strip" {
							fullPageBtn.Click()
						}
					} else {
						fmt.Println("FUCK")
					}
				}
			} else {
				panic(err)
			}
		}
		return true, nil
	}, 30*time.Second, 3*time.Second)

	if waitLoadErr != nil {
		return listPages, waitLoadErr
	}

	if pageLayoutElement, err := m.Driver.FindElement(selenium.ByID, "TopPage"); err == nil {
		if imgElements, err := pageLayoutElement.FindElements(selenium.ByClassName, "img-fluid"); err == nil {
			for idx, el := range imgElements {
				src, _ := el.GetAttribute("src")
				listPages = append(listPages, &Page{ImageUrl: src, Index: int32(idx) + 1, PageImageData: nil})
			}
		}

	}
	return listPages, nil
}
func (m *Manga4LifeParserService) ParseData(url string, start int32, end int32) *BookDataParseJson {
	info := m.GetInfo(url)
	chapters := m.GetListChapter(url)
	startIdx := godash.If(start == -1, 0, start)
	endIdx := godash.If(end == -1, int32(len(chapters)-1), end)
	for i := startIdx; i <= endIdx; i++ {
		pages, err := m.getChapterPages(chapters[i].ChapterUrl)
		if err != nil {
			panic(err)
		}
		chapters[i].Pages = pages
		chapters[i].TotalPage = int32(len(pages))
	}
	return &BookDataParseJson{
		Title:    info.Name,
		Authors:  info.Authors,
		Cover:    info.Thumbnail,
		Chapters: chapters[startIdx : endIdx+1],
	}
}

func (m *Manga4LifeParserService) KillService() error {
	return m.Service.Stop()
}
