package main

import (
	"fmt"
	"sync"
)

// ConcCrawl performs concurrent crawling of the provided URLs
// It uses the result channel to send back the titles fetched from each URL
func ConcCrawl(urls []string, result chan string) {
	fmt.Println("Concurrent crawling")

	// We use a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// 1. loop through each URL in the urls array (use for ... range)

	// 2. Start a gorutine for each URL to fetch the title concurrently
	// Use wg.Go(func(){}) to make sure we track the goroutines and wait for them to finish before returning from the function

	// 3. Inside the anonymous function
	// 3.1 use the Fetch function to get the title for each URL

	// 3.2 Use <- to send each fetched title to the result channel

	// Wait for all goroutines to complete
	wg.Wait()
}
