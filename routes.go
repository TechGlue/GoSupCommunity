//routes for the endpoints
package main 

import(
		"github.com/gofiber/fiber/v2"
		"fmt"
		"encoding/json"
)

func FetchItemsJson(){
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("Hello, World!, enpoints are /fetchSup to retrieve all Items for the season provided in the URL")
  })

  app.Post("/fetchsup", func(c *fiber.Ctx) error {
	url := c.FormValue("url")
    body := fetchHtml(url)
	return c.SendString(ConvertToJson(parseHTML(body)))
  })

  err := app.Listen(":3000")

  if err != nil {
	fmt.Println("Error: Failed to start server")
	fmt.Printf("%s", err)
  }else{
	fmt.Println("Server started on port 3000")
  }
}

func ConvertToJson( catalogItems []CatalogItem) string{
  json, err := json.MarshalIndent(catalogItems, "", "  ")
  if err != nil {
	fmt.Println("Error: Failed to convert catalogItems to JSON")
	fmt.Printf("%s", err)
  }
  return string(json)
} 
