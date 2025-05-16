package main

import (
	"strings"

	"golang.org/x/net/html"
)

// TO DO: ADD MORE TESTS
func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var foundURLs []string = []string{}
	r := strings.NewReader(htmlBody)
	parsedHTML, err := html.Parse(r)
	if err != nil {
		return []string{}, err
	}

	htmlIter := parsedHTML.Descendants()
	for i := range htmlIter {
		if i.Type == html.ElementNode && i.Data == "a" {
			for _, j := range i.Attr {
				if j.Key == "href" {
					foundURLs = append(foundURLs, j.Val)
				}
			}
		}
	}

	for i, path := range foundURLs {
		foundURLs[i] = convertToFullPath(path, rawBaseURL)
		continue
	}

	return foundURLs, nil
}

func convertToFullPath(rawPath, baseUrl string) string {
	if !strings.HasPrefix(rawPath, "/") {
		return rawPath
	}
	return baseUrl + rawPath
}
