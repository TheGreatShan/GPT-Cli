package davinci

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type OpenAIBody struct {
	Model             string  `json:"model"`
	Prompt            string  `json:"prompt"`
	Temperature       float32 `json:"temperature"`
	Max_tokens        int     `json:"max_tokens"`
	Top_p             float32 `json:"top_p"`
	Frequency_penalty float32 `json:"frequency_penalty"`
	Presence_penalty  float32 `json:"presence_penalty"`
}

type OpenAIResp struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func getKey() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	return os.Getenv("API_KEY")
}

func GetDavinci(inputText string) []byte {
	aiBody := OpenAIBody{
		Model:             "text-davinci-003",
		Prompt:            inputText,
		Temperature:       1,
		Max_tokens:        256,
		Top_p:             1,
		Frequency_penalty: 0,
		Presence_penalty:  0,
	}
	body, _ := json.Marshal(aiBody)
	url := "https://api.openai.com/v1/completions"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getKey())

	client := &http.Client{}
	resp, _ := client.Do(req)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	return respBody
}
