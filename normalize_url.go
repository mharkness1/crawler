package main

import (
	"net/url"
	"path"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Construct a normalized URL
	normalized := url.URL{
		Scheme: parsedUrl.Scheme,
		Host:   strings.ToLower(parsedUrl.Host),
		Path:   parsedUrl.Path,
	}

	// Clean the path to remove extra slashes and normalize
	normalized.Path = path.Clean(normalized.Path)

	// If the path was originally "/" it will be empty after clean
	if parsedUrl.Path == "/" || parsedUrl.Path == "" {
		normalized.Path = ""
	}

	return normalized.String(), nil
}
