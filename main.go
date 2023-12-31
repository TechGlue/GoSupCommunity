package main

import (
	"fmt"
	"os"
)

func main() {
	url := os.Args[1:]

	if len(url) == 0 {
		fmt.Println("Please provide a URL")
		return
	}

	body := fetchHtml(url[0])

	DumpToDiscord(parseHTML(body))
}
