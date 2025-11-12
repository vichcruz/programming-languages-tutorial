package main

import (
	"fmt"
)

func ConcCrawl(urls []string, titlesChan chan string) {
	fmt.Println("Start concurrent crawling")
	for _, url := range urls {
		go func(u string) {
			title := Fetch(u)
			titlesChan <- title
		}(url)
	}
}
