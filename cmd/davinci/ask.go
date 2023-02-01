package davinci

import (
	"fmt"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:     "ask",
	Aliases: []string{"a"},
	Short:   "Ask Davinci GPT-3 something",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args
		fmt.Printf(i[0])
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
