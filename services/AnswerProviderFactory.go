package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Shan15Dev/GPT-Cli/models"
	"github.com/google/uuid"
)

type IAnswerProvider interface {
	Ask(question string) models.Answer
}
type GPT4 struct {
}

func (gpt4 *GPT4) Ask(question string) models.Answer {
	aiBody := models.OpenAIBody{
		Model:       "gpt-4",
		Messages:    []models.Chat{{Role: "user", Content: question}},
		Temperature: 0.7,
	}
	body, _ := json.Marshal(aiBody)

	url := "https://api.openai.com/v1/chat/completions"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+GetKey("API_KEY"))
	client := &http.Client{}
	resp, _ := client.Do(req)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	var response models.OpenAIResp

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println(err)
	}
	return models.Answer{Id: uuid.New(), Answer: response.Choices[0].Message.Content, TimeStamp: time.Now()}
}

type AnswerProviderFactory struct {
}

func (answerProviderFactory *AnswerProviderFactory) CreateProvider(provider string) IAnswerProvider {
	switch provider {
	case "GPT4":
		return &GPT4{}
	case "mock":
		return nil
	default:
		return nil
	}
}
