package models

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func DetectEmotion(text string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API key não configurada")
	}

	client := openai.NewClient(apiKey)

	prompt := fmt.Sprintf("Detecte a emoção principal deste texto em português: \"%s\". Responda com uma única palavra.", text)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4o-mini",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	emotion := resp.Choices[0].Message.Content
	return emotion, nil
}
