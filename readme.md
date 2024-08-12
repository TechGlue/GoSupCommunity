# About
Basic web scraper to grab new catalog items from Supreme Community, given a drop week URL with discord webhooks integration. 

Note: GoSupCommunity is meant for personal use. Please be mindful before sending requests to any server.

# Usage

Modify the ***credentials.json** file with your discord webhook URL. Look [here](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks) for more info on creating a webhook. 
```json
{
    "webhookurl": "insert webhook url here"
}
```

Ensure all dependencies are installed and run. 
```bash
  go mod download  
  go build 
  ./GoSupCommunity https://www.supremecommunity.com/season/fall-winter2023/droplist/2023-10-26/
```
