package main

import (
	"crawler/asahi"
	"crawler/mainichi"
	"crawler/tokyo"
	"time"
)

func main() {
	go mainichi.Crawl()
	go tokyo.Crawl()
	go asahi.Crawl()

	time.Sleep(20 * time.Second)
}
