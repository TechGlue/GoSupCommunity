package main
import(
  "fmt"
  "net/http"
  "io"
)

func main() {
  url := "https://www.supremecommunity.com/season/fall-winter2023/droplist/2023-10-05/"
  fetchHtml(url)

}

func fetchHtml(url string) {
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
	fmt.Printf("%s\n", body)
  }
}
