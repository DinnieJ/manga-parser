package types

type MangaInfo struct {
	Thumbnail       string
	NumberOfChapter int32
	Name            string
	Description     string
	Authors         []string
	Genres          []string
}

type Chapter struct {
	Index      int32   `json:"-"`
	Name       string  `json:"chapterTitle"`
	ChapterUrl string  `json:"chapterUrl"`
	TotalPage  int32   `json:"totalPage"`
	Pages      []*Page `json:"pages"`
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
