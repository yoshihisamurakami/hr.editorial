package service

import (
	"fmt"
	"model"
)

//
type TopAnalizerInterface interface {
	GetEditorialLinks() []string
}

//
type ContentsAnalizerInterface interface {
	GetArticleContents(string) model.EditorialInfo
}

// Crawl ...
func Crawl(topAnalizer TopAnalizerInterface, contentsAnalizer ContentsAnalizerInterface) {
	latestLinks := topAnalizer.GetEditorialLinks()
	for i := 0; i < len(latestLinks); i++ {
		editorialInfo := contentsAnalizer.GetArticleContents(latestLinks[i])
		if editorialInfo.Count() == 0 {
			editorialInfo.Insert()
			fmt.Println(editorialInfo.Url + " Insert!")
		} else {
			fmt.Println(editorialInfo.Url + " (Update..)")
		}
	}
}
