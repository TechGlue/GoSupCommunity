package main
import(
  "fmt"
  "net/http"
  "io"
  "strings"
  "golang.org/x/net/html"
)

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
			fmt.Printf("Token: %v\n", html.UnescapeString(token.String()))
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
