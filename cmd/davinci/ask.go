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
		fmt.Printf(fmt.Sprintf("You: %s \n", i[0]))
		fmt.Printf(fmt.Sprintf("GPT-3: %s \n", davinci.GetDavinci(i[0])))
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
