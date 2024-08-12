package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"GoSupCommunity/scraping"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
)

type webhookcredentials struct {
	Webhookurl string `json:"webhookurl"`
}

func fetchCredentials() string {
	jsonContent, err := os.ReadFile("credentials.json")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		log.Fatal(err)
	}

	var config webhookcredentials

	err = json.Unmarshal(jsonContent, &config)

	if err != nil {
		fmt.Println("Error unmarshalling JSON data:", err)
		log.Fatal(err)
	}

	return config.Webhookurl
}

func DumpToDiscord(items []scraping.CatalogItem) {
	client, err := webhook.NewWithURL(fetchCredentials())
	if err != nil {
		fmt.Println(err, "trouble connecting to webhook")
		return
	}

	for _, item := range items {
		var embed []discord.Embed = make([]discord.Embed, 1)

		embed[0] = discord.Embed{
			Title:       item.ItemName,
			Image:       &discord.EmbedResource{URL: item.ItemImg, ProxyURL: item.ItemUrl},
			URL:         item.ItemUrl,
			Description: fmt.Sprintf("Price (USD): %s\nUpvotes :arrow_up: : %s\nDownvotes :arrow_down: : %s", item.ItemPrice, item.ItemUpVotes, item.ItemDownVotes),
			Color:       0xad6f49,
		}
		sendEmbed, err := client.CreateEmbeds(embed)
		if err != nil {
			fmt.Println("error sending webhook: ", err)
			fmt.Println("Embed", sendEmbed)
		}
	}

	defer client.Close(context.TODO())
}
