package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error reaching %s, returned: %v", rawURL, err)
	}

	if res.StatusCode > 400 {
		return "", fmt.Errorf("error from server at url %s, with status code: %s", rawURL, res.Status)
	}

	if res.Header.Get("Content-Type") != "text/html" {
		return "", fmt.Errorf("response type is not html")
	}

	resHTML, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading html response: %v", err)
	}
	defer res.Body.Close()

	return string(resHTML), nil
}

/*
Use http.Get to fetch the webpage of the rawURL
Return an error if the HTTP status code is an error-level code (400+)
Return an error if the response content-type header is not text/html
Return any other possible errors
Return the webpage's HTML if successful
*/
