package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	var isMaxPage bool = false
	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		isMaxPage = true
	}
	cfg.mu.Unlock()
	if isMaxPage == true {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseUrl.Hostname() {
		return
	}

	rawCurrentNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing url: %v", err)
		return
	}

	isFirst := cfg.addPageVisit(rawCurrentNormalized)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v", err)
		return
	}

	URLs, err := getURLsFromHTML(htmlBody, cfg.baseUrl)
	if err != nil {
		fmt.Printf("error getting urls from html: %v", err)
		return
	}

	for _, url := range URLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
