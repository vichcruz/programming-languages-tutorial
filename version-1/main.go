// This is the main starting file for the code.
// You can run this code using `go run main.go`

package main

import (
	"fmt"
	"time"
)

func main() {
	// List of website URLs to crawl
	websiteURLs := []string{
		"https://crawler-test.com/content/custom_text",
		"https://crawler-test.com/content/no_h1",
		"https://crawler-test.com/content/error_page",
		"https://www.wikipedia.org/",
		"https://www.bbc.com/",
		"https://www.nytimes.com/",
		"https://edition.cnn.com/",
		"https://www.reddit.com/",
		"https://www.stackoverflow.com/",
		"https://news.ycombinator.com/",
		"https://go.dev/",
		"https://developer.mozilla.org/",
		"https://www.w3.org/",
		"https://httpstat.us/200",
		"https://httpstat.us/404",
		"https://httpstat.us/500",
		"https://www.ecma-international.org/",
		"https://httpbin.org/html",
		"https://reqres.in/",
		"https://www.iana.org/domains/reserved",
	}

	fmt.Println("Number of websites to crawl:", len(websiteURLs))

	// Concurrent part
	go ConcCrawl(websiteURLs)

	// Synchronous part
	startSync := time.Now()

	titles := SyncCrawl(websiteURLs)

	fmt.Println("All Titles received: ", len(titles))

	elapsedSync := time.Since(startSync)

	fmt.Println("Synchronous crawling duration:", elapsedSync)
}
