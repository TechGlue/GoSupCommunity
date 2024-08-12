package main

import (
	"fmt"
	"os"
	"GoSupCommunity/scraping"
)

func main() {
	url := os.Args[1:]

	if len(url) == 0 {
		fmt.Println("No URL provided - Ensure a URL is provided as an argument")
		return
	}

	body := scraping.FetchHTML(url[0])
	DumpToDiscord(scraping.ParseHTML(body))
}
