package factories

import (
	"github.com/Shan15Dev/GPT-Cli/models"
)

type AnswerProvider interface {
	ask(question string) models.Answer
}
type GPT4 struct {
}

func (gpt4 GPT4) ask(question string) models.Answer {
	panic("Not implemented")
}

type AnswerProviderFactory struct {
}

func (answerProviderFactory AnswerProviderFactory) createProvider(provider string) AnswerProvider {
	switch "provider" {
	case "GPT4":
		return GPT4{}
	case "mock":
		return nil
	default:
		return nil
	}
}
