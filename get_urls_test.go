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
		}, {
			name:     "embedded link in header",
			inputURL: "www.base.com",
			inputBody: `
			<html>
				<header>
					<h1> Bluff <a href="www.home.com">Link</a></h1>
				</header>
				<body>
					<a href="/is_this_a_path/">Path 1</a>
					<a href="/PATH2">Path 2</a>
				</body>
			</html>
			`,
			expected: []string{"www.home.com", "www.base.com/is_this_a_path/", "www.base.com/PATH2"},
		},
		{
			name:     "Listed links",
			inputURL: "www.home.com",
			inputBody: `
			<html>
				<header>
					<h1>No link</h1>
					<nav>
					<ul>
					<li><a href="www.home.com">Home</a>></li>
					<li><a href="www.home.com/about">About</a></li>
					<li><a href="/portfolio">Portfoli</a></li>
					</ul>
					</nav>
				</header>
				<body>
					<h2>Header 2</h2>
				</body>
			</html>			
			`,
			expected: []string{"www.home.com", "www.home.com/about", "www.home.com/portfolio"},
		},
		{
			name:     "untagged plain links",
			inputURL: "www.home.com",
			inputBody: `
			<html>
				<header>
					<h1>Title</h1>
				</header>
				<body>
					<p>www.home.com</p>
				</body>
			</html>
			`,
			expected: []string{},
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
		{
			name:      "failure",
			pathInput: "not_a_path",
			urlInput:  "www.home.com",
			expected:  "not_a_path",
		},
		{
			name:      "multiple slashes",
			pathInput: "/this/is/a/path/",
			urlInput:  "www.home.com",
			expected:  "www.home.com/this/is/a/path/",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := convertToFullPath(tc.pathInput, tc.urlInput)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
