package davinci

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
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

func getKey() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	return os.Getenv("API_KEY")
}

func getDavinci() {
	aiBody := OpenAIBody{
		Model:             "text-davinci-003",
		Prompt:            "",
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

	return resp
}
