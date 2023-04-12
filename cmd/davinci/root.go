package davinci

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "davinci",
	Short: "Use OpenAis Davinci Language Model!",
	Long:  `With this command you can use the Davinci language Model by OpenAi. An Organisation founded by Elon Musk and invested by Microsoft`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
