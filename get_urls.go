package main

import (
	"strings"

	"golang.org/x/net/html"
)

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

// Tidy this up, its disgusting
func convertToFullPath(rawPath, baseUrl string) string {
	if !strings.HasPrefix(rawPath, "/") {
		return rawPath
	}
	rawPath = strings.TrimLeft(rawPath, "/")
	return baseUrl + rawPath
}
