# About
Basic web scraper to grab new catalog items from Supreme Community, given a drop week URL with discord webhooks integration. 

Note: GoSupCommunity is meant for personal use. Please be mindful before sending requests to any server.

# Usage

Modify the credentials.json file with you discord webhook URL. 

Ensure all dependencies are installed and run. 
```bash
  go mod download  
  go build 
  ./GoSupCommunity https://www.supremecommunity.com/season/fall-winter2023/droplist/2023-10-26/
```
