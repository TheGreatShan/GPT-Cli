package davinci

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "davinci",
	Short: "Use OpenAis Davinci Language Model!",
	Long:  `With this command you can use the Davinci language Model by OpenAi. An Organisation founded by Elon Musk and invested by Microsoft'`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	fmt.Printf("Use OpenAi's Davinci Language Model!")
}
