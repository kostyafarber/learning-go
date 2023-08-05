package main

import (
	"bytes"
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

type createIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type createIssueResponse struct {
	HtmlUrl string `json:"html_url"`
}

var (
	bearerToken = os.Getenv("GITHUB_GO_TOKEN")
	baseUrl     = "https://api.github.com/repos/kostyafarber/learning-go/%s"
)

func main() {
	c := http.Client{}
	issues := fmt.Sprintf(baseUrl, "issues")
	postBody := createIssue{
		Title: "Creating a test post request, again...",
		Body:  "It worked yay!",
	}

	postBodyJson, err := json.Marshal(postBody)
	if err != nil {
		log.Fatalf("Error making new marshals: %s", err)
	}

	post, err := http.NewRequest("POST", issues, bytes.NewBuffer(postBodyJson))
	if err != nil {
		log.Fatalf("Error making new request: %s", err)
	}

	post.Header.Add("Authorization", "Bearer "+bearerToken)
	post.Header.Add("Accept", "application/vnd.github+json")

	resp, err := c.Do(post)
	if err != nil {
		log.Fatalf("Error with get: %v", err)
	}
	defer resp.Body.Close()

	var issueResp createIssueResponse
	if err := json.NewDecoder(resp.Body).Decode(&issueResp); err != nil {
		log.Fatalf("Error with decoder: %v", err)
	}
	fmt.Printf("Issue created! Url of the issue at: %s\n", issueResp.HtmlUrl)
}
