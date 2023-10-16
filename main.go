package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"	
	"io"
	"net/http"
	"strings"
	"regexp"
	"golang.org/x/net/html"
)

type CatalogItem struct {
  ItemName string
  Price int64 
  ItemCategory string
  ItemUrl string
}

func main() {
  PrepareSupItemsTable()
  //url := "https://www.supremecommunity.com/season/fall-winter2023/droplist/2023-10-05/"
  //body := fetchHtml(url)
  //parseHTML(body)
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
			//loop through child tokens until we find div with class of catalog-list
			for {
			  if token.Data == "div"{
				fmt.Println("Found div")
			  }
			  if token.Data == "a"{
				fmt.Println("Extracted string", token.String())
				extracHref(html.UnescapeString(token.String()))
				break
			  }
			  tokenType = tokenizer.Next()
			  token = tokenizer.Token()
			}
		  }
		}
	}
 }
}

func extracHref(input string) string {
  hrefRegex := regexp.MustCompile(`<a\s+[^>]*href="([^"]+)"[^>]*>`)
  matches := hrefRegex.FindAllStringSubmatch(input, -1)

  // Extract and print href values
  for _, match := range matches {
  	if len(match) >= 2 {
  		href := match[1]
  		fmt.Println("Href:", href)
		return href
  	}
  }
  return "ERROR" 
} 

func craftURL(suffix string) string{
  return "https://www.supremecommunity.com" + suffix
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

func PrepareSupItemsTable(){
  db, _ := sql.Open("sqlite3", "./supItems.db")
  createTable, _ := db.Prepare(`
  CREATE TABLE IF NOT EXISTS items(
	id INTEGER NOT NULL PRIMARY KEY,
	ItemName TEXT NOT NULL,
	Price FLOAT NOT NULL,
	ItemCategory TEXT NOT NULL,
	ItemURL TEXT NOT NULL
  );
  `)
  _, err := createTable.Exec()
  if err != nil{
	fmt.Println("Error creating items table:", err)
  }
  db.Close()
}

func InsertIntoItemsTable(item CatalogItem){
  fmt.Println("Inserting into items table")
  db, _ := sql.Open("sqlite3", "./items.db")
  createItem, _ := db.Prepare("INSERT INTO items(itemName, price, itemCategory, itemURL) VALUES(?, ?, ?, ?)")
  _ ,err := createItem.Exec(item.ItemName, item.Price, item.ItemCategory, item.ItemUrl)
  if err != nil{
	fmt.Println("Error inserting into items table:", err)
  }
  db.Close()
}

func FetchItemsFromTable(){
  fmt.Println("Fetching items from table")
	db, _ := sql.Open("sqlite3", "./supItems.db")
	rows, _ := db.Query("SELECT * FROM items")
	var id int 
	var itemName string
	var price float64
	var itemCategory string
	var itemURL string
	for rows.Next(){
	  rows.Scan(&id, &itemName, &price, &itemCategory, &itemURL)
	  fmt.Println(&id, &itemName, &price, &itemCategory, &itemURL)
	}
	db.Close()
}

func CleanUpTable(){
  db, _ := sql.Open("sqlite3", "./supItems.db")
	cleanTable, _ := db.Prepare("DELETE FROM items")
	_, err := cleanTable.Exec()
	if err != nil{
	fmt.Println("Error cleaning items table:", err)
	}
	db.Close()
}
