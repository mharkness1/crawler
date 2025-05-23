package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove final slash",
			inputURL: "blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove http",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "capitalize letters",
			inputURL: "https://BLOG.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "keep query parameters",
			inputURL: "https://blog.boot.dev/path?query=value",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "additional test",
			inputURL: "https://www.wagslane.dev/tags/",
			expected: "www.wagslane.dev/tags",
		},
		{
			name:     "additional test 2",
			inputURL: "https://www.wagslane.dev/tags",
			expected: "www.wagslane.dev/tags",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
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
