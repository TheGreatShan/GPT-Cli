package davinci

import (
	"fmt"

	"github.com/Shan15Dev/GPT-Cli/services"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:     "ask",
	Aliases: []string{""},
	Short:   "Ask Davinci GPT-4 something",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args
		answerProviderFactory := &services.AnswerProviderFactory{}

		answerProvider := answerProviderFactory.CreateProvider("GPT4")
		if answerProvider == nil {
			panic("Could not create answer provider")
		}
		fmt.Print("You: " + i[0] + "\n")
		fmt.Print("GPT-4: " + answerProvider.Ask(i[0]).Answer + "\n")
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
