//future features
//being able to integrate with discord bots aka making a bot command
//Add a feature to grab images and display to webhook
//Add a feature to grab the current likes and dislikes of an item

package main

import (
  		"strings"
        "golang.org/x/net/html"
		"fmt"
		"io"
		"net/http"
)

type CatalogItem struct {
        ItemId       string `json:"item_id"`
        ItemName     string `json:"item_name"`
        ItemUrl      string `json:"item_url"`
}


func main() {
  FetchItemsJson()
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
                                        for {
                                                if token.Data == "a" {
                                                        var item CatalogItem
                                                        for _, a := range token.Attr {
                                                                if a.Val == "Go to home" {
																		//Found home tag indicating end of catalog items
                                                                        return items
                                                                }
                                                                switch a.Key {
                                                                case "href":
                                                                        item.ItemUrl = craftURL(a.Val)
                                                                case "data-itemid":
                                                                        item.ItemId = a.Val
                                                                case "data-itemname":
                                                                        item.ItemName = a.Val
                                                                }
                                                      }
													  if item.ItemId != "" && item.ItemName != "" && item.ItemUrl != "" {
														items = append(items, item)
													  }
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

func fetchHtml(url string) string {
        fmt.Println("Fetching HTML from", url)
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

