package main

import (
	"fmt"
	"os"
	"testing"
)

//base test is working and fills the que with 15 items

func TestParseHTML_ScrapePricesIfTheyExist(t *testing.T) {
	loadHtml, err := os.ReadFile("sample.html")

	if err != nil {
		fmt.Println("Error opening/reading file")
	}

	output := parseHTML(string(loadHtml))

	if len(output) != 15 {
		t.Errorf("Expected 15 items, got %d", len(output))
	}

	// Testing Testing price parsing
	for _, item := range output {
		if item.ItemId == "" {
			t.Errorf("Expected ItemId to be filled in, got %s", item.ItemId)
		}

		if item.ItemName == "" {
			t.Errorf("Expected ItemName to be filled in, got %s", item.ItemName)
		}

		if item.ItemUrl == "" {
			t.Errorf("Expected ItemURL to be filled in, got %s", item.ItemUrl)
		}

		if item.ItemImg == "" {
			t.Errorf("Expected ItemImg to be filled in, got %s", item.ItemImg)
		}

		if item.ItemPrice == "0.0" {
			t.Errorf("Expected price to be filled in for %s, got %s", item.ItemName, item.ItemPrice)
		}

		if item.ItemUpVotes == "" {
			t.Errorf("Expected ItemUpVotes to be filled in for %s, got %s", item.ItemName, item.ItemUpVotes)
		}
		if item.ItemDownVotes == "" {
			t.Errorf("Expected ItemDownVotes to be filled in for %s, got %s", item.ItemName, item.ItemDownVotes)
		}

		t.Logf("Item: %s, ItemURL: %s, ItemImg: %s,  Price: %s, UpVotes: %s, DownVotes: %s", item.ItemName, item.ItemImg, item.ItemUrl,  item.ItemPrice, item.ItemUpVotes, item.ItemDownVotes)
	}

	fmt.Println("items loaded: ", len(output))
}
