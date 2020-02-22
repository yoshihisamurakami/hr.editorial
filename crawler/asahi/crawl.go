package asahi

import (
	"model"
	"service"
)

type TopAnalizer struct {
	model.TopAnalizer
}

type ContentsAnalizer struct {
	model.ContentsAnalizer
}

func Crawl() {
	topAnalizer := TopAnalizer{}
	contentsAnalizer := ContentsAnalizer{}
	service.Crawl(topAnalizer, contentsAnalizer)
}
