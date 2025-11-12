package main

import (
	"fmt"
)

func ConcCrawl(urls []string) {
	fmt.Println("Start concurrent crawling")
	var length = len(urls)
	for i, url := range urls {
		go func(u string) {
			title := Fetch(u)
			fmt.Printf("[CONCURRENT %d/%d] Fetched title: %s\n", i+1, length, title)
		}(url)
	}
}
