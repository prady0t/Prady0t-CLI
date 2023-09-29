package controls

import (
	"fmt"
	"github.com/spf13/cobra"
)

var hello = &cobra.Command{
	Use:   "hello",
	Short: "hello user",
	Long: ` `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("I greet you.")
	},
}

func init() {
	rootCmd.AddCommand(hello)
}
