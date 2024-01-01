package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

func DumpToDiscord(items []CatalogItem) {
	client, err := webhook.NewWithURL(fetchCredentials())
	if err != nil {
		fmt.Println(err, "trouble connecting to webhook")
		return
	}

	for _, item := range items {
		var embed []discord.Embed = make([]discord.Embed, 1)
		embed[0] = discord.Embed{
			Title:       item.ItemName,
			Description: item.ItemUrl,
			Image:       &discord.EmbedResource{URL: item.ItemImg},
		}
		sendEmbed, err := client.CreateEmbeds(embed)
		if err != nil {
			fmt.Println("error sending webhook: ", err)
			fmt.Println("Embed", sendEmbed)
		}
	}

	defer client.Close(context.TODO())
}
