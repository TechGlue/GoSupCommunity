package main

import (
	"context"
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
)

func DumpToDiscord(items []CatalogItem) {
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
