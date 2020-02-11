package main

import (
	"crawler/mainichi"
	"crawler/tokyo"
)

func main() {
	mainichi.Crawl()
	tokyo.Crawl()
	// go mainichi.Crawl()
	// go tokyo.Crawl()
	// time.Sleep(10 * time.Second)
}
