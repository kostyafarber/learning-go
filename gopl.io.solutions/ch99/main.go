package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	currComic = "https://xkcd.com/info.0.json"
	baseUrl   = "https://xkcd.com/%s/info.0.json"
)

type xkcdClient struct {
	baseUrl string
	client  *http.Client
}

type xkcdResponse struct {
	Num   int    `json:"num"`
	Title string `json:"title"`
}

func createxkcdClient(url string) *xkcdClient {
	return &xkcdClient{
		baseUrl: url,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func getNumberOfComics(xkcd *xkcdClient, url string) int {
	resp, err := xkcd.client.Get(url)
	if err != nil {
		log.Fatalf("error with get: %v", err)
	}
	defer resp.Body.Close()

	xkcdJson := xkcd.parseResponse(resp)
	return xkcdJson.Num
}

func (c *xkcdClient) parseResponse(resp *http.Response) xkcdResponse {
	var xkcdJson xkcdResponse
	if err := json.NewDecoder(resp.Body).Decode(&xkcdJson); err != nil {
		log.Fatalf("error with decode: %v", err)
	}

	return xkcdJson
}

func main() {
	client := createxkcdClient(baseUrl)
	numComics := getNumberOfComics(client, currComic)
	fmt.Println(numComics)
}
