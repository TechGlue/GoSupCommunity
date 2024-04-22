package main

import (
	"encoding/json"
	"fmt"
	"GoSupCommunity/scraping"

	"github.com/gofiber/fiber/v2"
)

func FetchItemsJson() {
	app := fiber.New()

	app.Post("/fetchsup", func(c *fiber.Ctx) error {
		url := c.FormValue("url")
		body := scraping.FetchHTML(url)
		return c.SendString(ConvertToJson(scraping.ParseHTML(body)))
	})

	fmt.Println("Server started on port 3000")

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error: Failed to start server")
		fmt.Printf("%s", err)
	}
}

func ConvertToJson(catalogItems []scraping.CatalogItem) string {
	json, err := json.MarshalIndent(catalogItems, "", "  ")
	if err != nil {
		fmt.Println("Error: Failed to convert catalogItems to JSON")
		fmt.Printf("%s", err)
	}
	return string(json)
}
