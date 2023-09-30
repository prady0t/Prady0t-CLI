package controls

import (
	"Prady0t/prady0t/scraper"

	"github.com/spf13/cobra"
)


var github = &cobra.Command{
	Use:   "github",
	Short: "know my github",
	Long: ` `,
	Run: func(cmd *cobra.Command, args []string) {
		scraper.Finalized_github_output()
	},
}

func init() {
	rootCmd.AddCommand(github)
}


