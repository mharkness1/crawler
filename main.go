package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	initialURL := args[1]
	maxConcurrency, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("error parsing max concurrency, should be integer")
		return
	}
	maxPages, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("error parsing max pages, should be integer")
		return
	}

	cfg, err := configure(initialURL, maxPages, maxConcurrency)
	if err != nil {
		fmt.Printf("Error in configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", initialURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(initialURL)
	cfg.wg.Wait()

	printReport(cfg.pages, initialURL)
}
