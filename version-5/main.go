// This is the main starting file for the code.
// You can run this code using `go run .`

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
	// How it works
	/*
		1. Main goroutine starts the concurrent crawler as a goroutine
		2. Concurrent crawler launches 23 goroutines (one for each URL) - all start fetching simultaneously
		3. Main goroutine enters the for loop and waits at title := <-concResult
		4. First goroutine to finish sends its title to the channel
		5. Main goroutine immediately receives it, prints it, and loops back to wait for the next one
		6. Second goroutine to finish sends its title → main receives and prints
		7. Third goroutine to finish sends → main receives and prints
		... and so on until all 23 titles are received
	*/
	startConc := time.Now()
	titleResult := make(chan string, len(websiteURLs))

	go ConcCrawl(websiteURLs, titleResult)

	for range websiteURLs {

		title := <-titleResult
		if title != "" {
			fmt.Println("Title received:", title)
		} else {
			fmt.Println("Failed to fetch title.")
		}

	}

	durationConc := time.Since(startConc)
	fmt.Println("Concurrent crawling duration:", durationConc)

	fmt.Println("---------------------------------------------------")

	// Synchronous part
	startSync := time.Now()

	SyncCrawl(websiteURLs)

	fmt.Println("All Titles received.")

	elapsedSync := time.Since(startSync)

	fmt.Println("Synchronous crawling duration:", elapsedSync)
}
