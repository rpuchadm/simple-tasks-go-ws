package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {

	// lee variable del fichero AZURE_OPENAI_API_KEY
	data, err := os.ReadFile("AZURE_OPENAI_API_KEY")
	if err != nil {
		fmt.Println("File AZURE_OPENAI_API_KEY reading error", err)
		return
	}
	AZURE_OPENAI_API_KEY := string(data)
	if AZURE_OPENAI_API_KEY == "" {
		fmt.Println("AZURE_OPENAI_API_KEY file is empty")
		return
	} else {
		fmt.Println("AZURE_OPENAI_API_KEY file is read " + AZURE_OPENAI_API_KEY)
	}
	AZURE_OPENAI_BASE_URL := "https://poligpt.upv.es/gpt/ufasu-ia"

	config := openai.DefaultAzureConfig(AZURE_OPENAI_API_KEY, AZURE_OPENAI_BASE_URL)
	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:  openai.GPT3Dot5TurboInstruct,
			Prompt: "Hello!",
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Text)

}
