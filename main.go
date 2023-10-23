package main

import (
        "fmt"
        "io"
        "net/http"
		"context"
        "strings"
        "golang.org/x/net/html"
		"github.com/disgoorg/disgo/discord"
		"github.com/disgoorg/disgo/webhook"
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

func fetchCatalogItems(url string) {
        body := fetchHtml(url)
		catlogItems := parseHTML(body)
		dumpToDiscord(catlogItems)
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

func dumpToDiscord(items []CatalogItem)	{
  fmt.Println("-------WEBHOOK STARTED-------")
  client, err := webhook.NewWithURL("https://discord.com/api/webhooks/1145495004354187294/YdKt5wog8g60-RIDARmTKdcYURfRbShidx9QjOKBmUqAjamUvJCxcuc9oHP0c1ytgrtu")
  if err != nil {
	fmt.Println(err, "trouble connecting to webhook")
	return
  }

  for _, item := range items {
	var embed []discord.Embed = make([]discord.Embed, 1)
  	embed[0] = discord.Embed{
  	  Title:       item.ItemName,
  	  Description: item.ItemUrl,
  	}
  	sendEmbed, err := client.CreateEmbeds(embed) 
  	if err != nil {
  	  fmt.Println("error sending webhook: ", err)	
  	  fmt.Println("Embed", sendEmbed)	
  	}
  }

  defer client.Close(context.TODO())
}

