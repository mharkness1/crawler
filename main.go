package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	initialURL := args[1]
	fmt.Printf("starting crawl of: %s", initialURL)
	htmlBody, err := getHTML(initialURL)
	if err != nil {
		// handle error
	}
	fmt.Printf("%s", htmlBody)
}
