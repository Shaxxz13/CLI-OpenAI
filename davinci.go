package main

import (
	"context"
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("CHATGPT_TOKEN")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	// Get the prompt from the user
	fmt.Print("Enter a prompt: ")

	in := bufio.NewReader(os.Stdin)
	prompt, err := in.ReadString('\n')

	//Stop the CLI with exit input
	if prompt == "exit" {
		fmt.Println("EXIT")
		os.Exit(1)
	}

	

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{prompt},
		MaxTokens: gpt3.IntPtr(10),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	fmt.Println(resp.Choices[0].Text)
}
