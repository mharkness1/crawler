package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Convert host to lowercase
	host := strings.ToLower(parsedUrl.Host)

	// Handle path - ensure it doesn't have a trailing slash
	path := strings.ToLower(parsedUrl.Path)
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	// Handle query parameters
	var query string = ""
	if parsedUrl.RawQuery != "" {
		query = "?" + parsedUrl.RawQuery
	}

	// Connect host and path with a single slash
	normalizedUrl := host
	if path != "" {
		normalizedUrl += path
	}
	normalizedUrl += query
	/*
		parsedURL.Scheme // "http" or "https"
		parsedURL.Path   // The path portion
		parsedURL.Host   // Contains both the hostname and port
		parsedURL.RawQuery // Query parameters
		General URL structure:
		[scheme:][//[userinfo@]host][/]path[?query][#fragment]
	*/

	return normalizedUrl, nil
}
