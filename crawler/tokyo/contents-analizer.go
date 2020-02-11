package tokyo

import (
	"fmt"
	"io/ioutil"
	"model"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const IsContentDownloadMode = true
const EditorialTopDir = "https://www.tokyo-np.co.jp/article/column/editorial/"
const ContentsAnalizerFile = "./tokyo/html/contents.html"
const MediaIdTokyo = 2

func (c ContentsAnalizer) GetArticleContents(url string) model.EditorialInfo {
	if IsContentDownloadMode == true {
		c.Doc = c.getContentsDocByDownload(EditorialTopDir + url)
	} else {
		c.Doc = c.getContentsDocFromFile()
	}

	e := model.EditorialInfo{}
	e.Url = url
	e.Date = c.getDate()
	e.MediaId = MediaIdTokyo
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
	fileInfos, _ := ioutil.ReadFile(ContentsAnalizerFile)
	stringReader := strings.NewReader(string(fileInfos))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		panic(err)
	}
	return doc
}

func (c ContentsAnalizer) getTitle() string {
	titleSjis := c.Doc.Find("h1").Text()
	return sjisToUtf8(titleSjis)
}

func (c ContentsAnalizer) getDate() string {
	dateTextSjis := c.Doc.Find(".data").Text()
	date := sjisToUtf8(dateTextSjis)

	rep := regexp.MustCompile("(\\d+)年(\\d+)月(\\d+)日")
	result := rep.FindStringSubmatch(date)
	month, _ := strconv.Atoi(result[2])
	day, _ := strconv.Atoi(result[3])
	return fmt.Sprintf("%s-%02d-%02d", result[1], month, day)
}

func (c ContentsAnalizer) getContents() string {
	var body []string
	paragraphes := c.Doc.Find(".News-textarea > .Text > p")
	paragraphes.Each(func(index int, paragraph *goquery.Selection) {
		para := sjisToUtf8(paragraph.Text())
		line := strings.TrimSpace(para)
		if len(line) > 0 {
			body = append(body, line)
		}
	})
	return strings.Join(body, "\n")
}

func sjisToUtf8(sjisStr string) string {
	reader := strings.NewReader(sjisStr)
	rio := transform.NewReader(reader, japanese.ShiftJIS.NewDecoder())
	ret, _ := ioutil.ReadAll(rio)
	return string(ret)
}
