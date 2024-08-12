# GoSupCommunity 
## About
Go project to help familiarize myself with Go's syntax, structure, and libraries. 

Given a drop week URL, this service scrapes the contents of the provided drop and returns item names, images, and the corresponding URLs in a well-organized JSON response.

Note: GoSupCommunity is meant for personal use. Please be mindful before sending requests to any server.

# Building and running
1. Install dependencies:
    ```bash
    go mod download
    ```
2. Build the service:
    ```bash
    go build
    ```
## Usage
1. Start the service:
    ```bash
    ./GoSupCommunity
    ```

2. Access the API endpoint:
    - Trigger scraping and fetch listings: `POST /fetchSup`

### Sample cURL Requests
```
curl --location 'http://localhost:3000/fetchSup' \
--form 'url="https://www.supremecommunity.com/season/fall-winter2023/droplist/2023-12-21/"'
```
### Sample JSON Output
```json
[
  {
    "item_id": "10648",
    "item_name": "Supreme/Corteiz Rules The World Tee",
    "item_url": "https://www.supremecommunity.com/season/itemdetails/10648/supreme-corteiz-rule-the-world-tee/",
    "item_img": "https://www.supremecommunity.com/u/season/add/20231220/94cd60c03a4e42a8946331a4e1b9e0f4_sqr.jpg",
    "item_price": "44",
    "item_upvotes": "1240",
    "item_downvotes": "125"
  },
  {
    "item_id": "10475",
    "item_name": "Supreme®/RIDE Snowboards®",
    "item_url": "https://www.supremecommunity.com/season/itemdetails/10475/supreme-r-ride-r-snowboard/",
    "item_img": "https://www.supremecommunity.com/u/season/fall-winter2023/accessories/fall-winter2023-supreme-ride-snowboard-0-front_sqr.jpg",
    "item_price": "648",
    "item_upvotes": "900",
    "item_downvotes": "174"
  }
]
```

### Misc. 
If a service isn't your thing. Checkout the discord webhooks integration that I did. Same functionality different way of delivering the data [discord-webhook](https://github.com/TechGlue/GoSupCommunity/tree/discord-webhook). 
