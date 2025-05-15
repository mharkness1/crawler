package main

import (
	"reflect"
	"testing"
)

func TestGetUrls(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "standard functionality <a> tag",
			inputURL: "https://blog.boot.dev",
			inputBody: `
						<html>
							<body>
								<a href="/path/one">
									<span>Boot.dev</span>
								</a>
								<a href="https://other.com/path/one">
									<span>Boot.dev</span>
								</a>
							</body>
						</html>
						`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Test %v - '%s' FAIL: expected output: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestConvertToFullPath(t *testing.T) {
	tests := []struct {
		name      string
		pathInput string
		urlInput  string
		expected  string
	}{
		{
			name:      "standard test",
			pathInput: "/path",
			urlInput:  "www.boot.dev",
			expected:  "www.boot.dev/path",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := convertToFullPath(tc.pathInput, tc.urlInput)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
