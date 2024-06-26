# GoSupCommunity 
## About
A small GO project to help familiarize myself with Go's syntax, structure, and libraries. 

What this program does is given a drop week URL, this service scrapes the contents of the provided drop and returns item names, images, and the corresponding URLs in a well-organized JSON response.

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
### Misc. 
If a service isn't your thing. Checkout the discord webhooks integration that I did. Same functionality different way of delivering the data [discord-webhook](https://github.com/TechGlue/GoSupCommunity/tree/discord-webhook). 
