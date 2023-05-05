package davinci

import (
	"fmt"

	"github.com/Shan15Dev/GPT-Cli/pkg/davinci"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:     "ask",
	Aliases: []string{""},
	Short:   "Ask Davinci GPT-3 something",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args
		fmt.Print("You: " + i[0] + "\n")
		fmt.Print("GPT-3: " + davinci.GetDavinci(i[0]) + "\n")
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
