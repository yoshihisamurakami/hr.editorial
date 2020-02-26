package main

import (
	"crawler/mainichi"
	"crawler/tokyo"
	"time"
)

func main() {
	go mainichi.Crawl()
	go tokyo.Crawl()
	//asahi.Crawl()

	time.Sleep(20 * time.Second)
}
