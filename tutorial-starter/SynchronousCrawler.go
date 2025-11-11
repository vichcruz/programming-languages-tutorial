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
	// Make sure to use := for variable assignment inside the loop
	// You can deconstruct the slice in the loop into index and url
	// To omit the index, use underscore _ (like in python)

	// 2. Use the Fetch function we defined in the Fetcher file to get the title for each URL

	// 3. Append the fetched title to the results array using the append(slice, element) function

	// 4. Print the title for each URL if it has one or print a message if no title was found

	return results
}
