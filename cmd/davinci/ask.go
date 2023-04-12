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
		fmt.Println(fmt.Sprintf("You: %s", i[0]))
		fmt.Println(fmt.Sprintf("GPT-3: %s", davinci.GetDavinci(i[0])))
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
