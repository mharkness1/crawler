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
		expected  string
	}{
		{
			name:      "",
			inputURL:  "",
			inputBody: "",
			expected:  "",
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

/* 	for i, tc := range tests {
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
*/
