package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"golang.org/x/oauth2"
)

const (
	baseURL = "https://api.openai.com/v1/"
)

func main() {
	// Set up the HTTP client with the access token
	token := os.Getenv("CHATGPT_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "CHATGPT_TOKEN environment variable must be set")
		os.Exit(1)
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)

	// Set up the request to the ChatGPT API
	prompt := "What's your favorite color?"
	model := "text-davinci-002"
	maxTokens := 50
	temperature := 0.5
	reqBody := map[string]interface{}{
		"prompt":       prompt,
		"model":        model,
		"max_tokens":   maxTokens,
		"temperature":  temperature,
		"presence":     0.5,
		"stop":         "",
		"stream":       false,
		"max_tokens":   50,
		"temperature":  0.5,
		"top_p":        1,
		"frequency_penalty": 0,
		"presence_penalty": 0,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %s", err)
		os.Exit(1)
	}
	resp, err := tc.Post(baseURL+"chat", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making request: %s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %s", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
