package main

import (
        "fmt"
        "io"
        "net/http"
        "strings"
        "golang.org/x/net/html"
)

type CatalogItem struct {
        ItemId       string
        ItemName     string
        ItemUrl      string
}

func main() {
        fetchCatalogItems("https://www.supremecommunity.com/season/fall-winter2023/droplist/2023-10-05/")
}

func convertCatalogItemsToJSON(items []CatalogItem) string {
        return "JSON"
}

func fetchCatalogItems(url string) []CatalogItem {
        body := fetchHtml(url)
		catlogItems := parseHTML(body)
		return catlogItems
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
        fmt.Println("Fetching HTML from", url, "\n")
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
