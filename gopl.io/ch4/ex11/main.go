package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Issues []struct {
	NodeId string `json:"node_id"`
	Title  string
}

var (
	bearerToken = os.Getenv("GITHUB_GO_TOKEN")
)

func main() {
	url := "https://api.github.com/repos/kostyafarber/learning-go/issues"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error with get: %v", err)
	}
	defer resp.Body.Close()

	var issues Issues
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		log.Fatalf("Error with decoder: %v", err)
	}
	fmt.Printf("Title of the error is: %s\n", issues[0].Title)
}
