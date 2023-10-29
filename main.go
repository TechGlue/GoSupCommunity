//future features
//being able to integrate with discord bots aka making a bot command
//Add a feature to grab images and display to webhook
//Add a feature to grab the current likes and dislikes of an item

package main

import (
        "fmt"
        "io"
        "strings"
        "golang.org/x/net/html"
		"net/http"
)

type CatalogItem struct {
        ItemId       string
        ItemName     string
        ItemUrl      string
}

func main() {
  //start listening
  fmt.Println("Starting server on port 8080")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	  fmt.Fprintf(w, "Path: %s!", r.URL.Path[1:])
  })

  http.ListenAndServe(":8080", nil)
  //fetchCatalogItems("https://www.supremecommunity.com/season/fall-winter2023/droplist/2023-10-26/")
}

func fetchCatalogItems(url string) {
        body := fetchHtml(url)
		catlogItems := parseHTML(body)
		DumpToDiscord(catlogItems)
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
                                                    items = append(items, item)
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

