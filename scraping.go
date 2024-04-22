package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type CatalogItem struct {
	ItemId        string `json:"item_id"`
	ItemName      string `json:"item_name"`
	ItemUrl       string `json:"item_url"`
	ItemImg       string `json:"item_img"`
	ItemPrice     string `json:"item_price"`
	ItemUpVotes   string `json:"item_upvotes"`
	ItemDownVotes string `json:"item_downvotes"`
}

func parseHTML(rawHTML string) []CatalogItem {
	var items []CatalogItem
	tokenizer := html.NewTokenizer(strings.NewReader(rawHTML))

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()

		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				fmt.Println("EOF")
				return items
			}
			fmt.Printf("Error: %v", tokenizer.Err())
			return items
		}

		switch token.Data {
		case "div":
			for _, a := range token.Attr {
				if a.Val == "catalog-inner" {
					currentImageURL := ""
					var item CatalogItem

					for {
						if token.Data == "img" {
							currentImageURL = token.Attr[0].Val
						}

						if token.Data == "div" {
							for _, a := range token.Attr {
								switch a.Key {
								case "data-usdprice":
									item.ItemPrice = a.Val
								case "data-upvotes":
									item.ItemUpVotes = a.Val
								case "data-downvotes":
									item.ItemDownVotes = a.Val
								}
							}
						}

						if token.Data == "a" {
							for _, a := range token.Attr {
								if a.Val == "Go to home" {
									return items
								}
								switch a.Key {
								case "data-usdprice":
									fmt.Println("Price: ", a.Val)
								case "href":
									item.ItemUrl = craftURL(a.Val)
								case "data-itemid":
									item.ItemId = a.Val
								case "data-itemname":
									item.ItemName = a.Val
								}
							}
						}
						if item.ItemId != "" && item.ItemName != "" && item.ItemUrl != "" && currentImageURL != "" && item.ItemUpVotes != "" && item.ItemDownVotes != "" {
							item.ItemImg = craftURL(currentImageURL)

							// if no item is currently present then marking price as N/A
							if item.ItemPrice == "" {
							  item.ItemPrice = "N/A"
							}

							items = append(items, item)

							// reset temp variables to defaults 
							currentImageURL = ""
							item = CatalogItem{}
						}

						tokenType = tokenizer.Next()
						token = tokenizer.Token()
					}
				}
			}
		}
	}
}

func craftURL(suffix string) string {
	return "https://www.supremecommunity.com" + suffix
}

func fetchHTML(url string) string {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("Error: Failed to fetch the HTML from", url)
		fmt.Printf("%s", err)
	} else {
		body, error := io.ReadAll(resp.Body)
		if error != nil {
			fmt.Printf("%s", error)
		}
		resp.Body.Close()
		return string(body)
	}
	return "ERROR"
}
