package mainichi

import (
	"fmt"
	"model"
)

type TopAnalizer struct {
	model.TopAnalizer
}

type ContentsAnalizer struct {
	model.ContentsAnalizer
}

func Crawl() {
	topAnalizer := TopAnalizer{}
	latestLinks := topAnalizer.GetEditorialLinks()
	for i := 0; i < len(latestLinks); i++ {
		contentsAnalizer := ContentsAnalizer{}
		editorialInfo := contentsAnalizer.GetArticleContents(latestLinks[i])
		//fmt.Println(editorialInfo)
		if editorialInfo.Count() == 0 {
			editorialInfo.Insert()
			fmt.Println(editorialInfo.Url + " Insert!")
		} else {
			fmt.Println(editorialInfo.Url + " Update!")
		}
	}
}
