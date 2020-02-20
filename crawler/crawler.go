package main

import (
	"crawler/mainichi"
	"crawler/tokyo"
	"time"
)

func main() {
	go mainichi.Crawl()
	go tokyo.Crawl()
	// go mainichi.Crawl()
	// go tokyo.Crawl()
	time.Sleep(20 * time.Second)
}
