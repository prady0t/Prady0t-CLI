package controls

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)


var about = &cobra.Command{
	Use:   "about",
	Short: "know me",
	Long: ` `,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("My name is pradyot")
	},
}

func init() {
	rootCmd.AddCommand(about)
}


