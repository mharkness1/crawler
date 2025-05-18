package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	pages := make(map[string]int)

	initialURL := args[1]

	fmt.Printf("starting crawl of: %s", initialURL)
	crawlPage(initialURL, initialURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	rawCurrentNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing url: %v", err)
		return
	}

	if _, ok := pages[rawCurrentNormalized]; ok {
		pages[rawCurrentNormalized]++
		return
	}

	pages[rawCurrentNormalized] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v", err)
		return
	}

	URLs, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("error getting urls from html: %v", err)
		return
	}

	fmt.Print(URLs)

	for _, url := range URLs {
		crawlPage(rawBaseURL, url, pages)
	}
}
