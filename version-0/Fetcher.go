package main

import (
	"io"
	"net/http"
	"strings"
)

func Fetch(url string) string {

	// Simple HTTP GET request to check if the URL is reachable
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	// Read actual response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	bodyString := string(bodyBytes)

	// Schedule closing of body at the end of the function
	// using defer to ensure it runs even in error/panic cases
	defer resp.Body.Close()

	// Get title from response body
	titleStart := strings.Index(bodyString, "<title>")
	titleEnd := strings.Index(bodyString, "</title>")
	if titleStart != -1 && titleEnd != -1 {
		title := bodyString[titleStart+len("<title>") : titleEnd]
		return title
	} else {
		return ""
	}
}
