package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type CatalogItem struct {
  ItemName string
  Price int64 
  ItemCategory string
  ItemUrl string
}

func main() {
  url := "https://www.supremecommunity.com/season/fall-winter2023/droplist/2023-10-05/"
  body := fetchHtml(url)
  parseHTML(body)

}

//div catalog-inner is what has the 
func parseHTML(rawHTML string){
  tokenizer := html.NewTokenizer(strings.NewReader(rawHTML))
  for {
	tokenType := tokenizer.Next()
  	token := tokenizer.Token()

	//Error handling
  	if tokenType == html.ErrorToken {
  	 if tokenizer.Err() == io.EOF {
  	  return
  	 }
  	 fmt.Printf("Error: %v", tokenizer.Err())
  	 return
  	}

	switch token.Data{
	  case "div":
		for _, a := range token.Attr {
		  if a.Val == "catalog-inner"{
			count := 0
			//loop through child tokens until we find div with class of catalog-list

			fmt.Println(token.Data)
			for token.Data != "div"{
			  fmt.Println("Number of divs skipped", count)
			  tokenType = tokenizer.Next()

			  //skip the blank line
			  tokenType = tokenizer.Next()

			  token = tokenizer.Token()
			  fmt.Println("Token", token.Data)
			  count++
			}

			fmt.Println("Final Number of divs skipped", count)
			fmt.Printf("End Of Token: %v\n", html.UnescapeString(token.String()))
		  }
		}
	}
 }
}

func fetchHtml(url string) string {
  fmt.Println("Fetching HTML from", url, "\n")
  resp, err := http.Get(url)
  if err != nil || resp.StatusCode != 200{
	fmt.Println("Error: Failed to fetch the HTML from", url)
	fmt.Printf("%s", err)
  }else{
	body, error := io.ReadAll(resp.Body)
	if error != nil {
	  fmt.Printf("%s", error)
	}
	resp.Body.Close()
	return string(body)
  }
  return "ERROR"
}
