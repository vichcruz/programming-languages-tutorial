package main

import (
	"fmt"
)

func SyncCrawl(urls []string) []string {
	fmt.Println("Start synchronous crawling")

	// Build the synchronous version of the crawler

	// We define results array (in Go this is called a slice of strings)
	var results []string

	// 1. loop through each URL in the urls array (use for ... range)
	for _, url := range urls {

		// 2. Use the Fetch function we defined in the Fetcher file to get the title for each URL
		title := Fetch(url)

		// 3. Append the fetched title to the results array using the append(slice, element) function
		results = append(results, title)

		// Print the title for each URL
		if title != "" {
			fmt.Printf("Fetched title: %s\n", title)
		} else {
			fmt.Printf("No title found for URL: %s\n", url)
		}

	}

	return results
}
