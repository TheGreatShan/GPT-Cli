package models

type OpenAIBody struct {
	Model       string  `json:"model"`
	Messages    []Chat  `json:"messages"`
	Temperature float32 `json:"temperature"`
}

type OpenAIResp struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Message      Chat   `json:"message"`
	} `json:"choices"`
	Created int    `json:"created"`
	Id      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Chat struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
