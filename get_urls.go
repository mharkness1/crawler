package main

import "strings"

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	return nil, nil
}

func convertToFullPath(rawPath, baseUrl string) string {
	if !strings.HasPrefix(rawPath, "/") {
		return ""
	}
	return baseUrl + rawPath
}
