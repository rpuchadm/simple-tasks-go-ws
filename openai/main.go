package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

const AZURE_OPENAI_API_KEY = "your key"
const AZURE_OPENAI_BASE_URL = "their endpoint"

func main() {
	config := openai.DefaultAzureConfig(AZURE_OPENAI_API_KEY, AZURE_OPENAI_BASE_URL)
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
