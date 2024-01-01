package main

import (
	"fmt"
	"os"
)

func main() {
	url := os.Args[1:]

	if len(url) == 0 {
		fmt.Println("No URL provided")
		return
	}

	body := fetchHtml(url[0])
	DumpToDiscord(parseHTML(body))
}
