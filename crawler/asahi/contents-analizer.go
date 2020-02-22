package asahi

import (
	"fmt"
	"io/ioutil"
	"model"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const IsContentDownloadMode = true
const ContentsPageFile = "./asahi/html/contents.html"
const MediaIdAsahi = 3

func (c ContentsAnalizer) GetArticleContents(url string) model.EditorialInfo {
	if IsContentDownloadMode == true {
		c.Doc = c.getContentsDocByDownload(url)
	} else {
		c.Doc = c.getContentsDocFromFile()
	}
	e := model.EditorialInfo{}
	e.Url = url
	e.Date = c.getDate()
	e.MediaId = MediaIdAsahi
	e.Title = c.getTitle()
	e.Body = c.getContents()
	return e
}

func (c ContentsAnalizer) getContentsDocByDownload(url string) *goquery.Document {
	time.Sleep(1 * time.Second)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	return doc
}

func (c ContentsAnalizer) getContentsDocFromFile() *goquery.Document {
	fileInfos, _ := ioutil.ReadFile(ContentsPageFile)
	stringReader := strings.NewReader(string(fileInfos))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		panic(err)
	}
	return doc
}

func (c ContentsAnalizer) getTitle() string {
	title := c.Doc.Find("h1").Text()
	title = strings.TrimSpace(title)
	title = strings.Replace(title, "（社説）", "", 1)
	return title
}

func (c ContentsAnalizer) getDate() string {
	dateText := c.Doc.Find(".UpdateDate > time").Text()
	rep := regexp.MustCompile("(\\d+)年(\\d+)月(\\d+)日")
	result := rep.FindStringSubmatch(dateText)
	month, _ := strconv.Atoi(result[2])
	day, _ := strconv.Atoi(result[3])
	return fmt.Sprintf("%s-%02d-%02d", result[1], month, day)
}

func (c ContentsAnalizer) getContents() string {
	var body []string
	paragraphes := c.Doc.Find(".ArticleText > p")
	paragraphes.Each(func(index int, paragraph *goquery.Selection) {
		line := strings.TrimSpace(paragraph.Text())
		if len(line) > 0 {
			body = append(body, line)
		}
	})
	return strings.Join(body, "\n")
}
