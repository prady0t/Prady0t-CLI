package controls

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var hello = &cobra.Command{
	Use:   "hello",
	Short: "hello user",
	Long: ` `,
	Run: func(cmd *cobra.Command, args []string) {
		printer()
	},
}

func printer(){
	color.Red("I greet you!")
	color.Cyan("Long before the dawn of mankind, when electrons were free to move and were \nnot bounded by the unbreakable vow of entanglement,\n when electrons were not domesticated\nto produce answers as per a human's will,\nthere lived a Gromflomite serving the glactic federation.\nStruggling through life he constantly had the urge to get recognized.\nBut he died and we don't know his name anymore.")
	color.Blue("I want to get recognized too!")
}

func init() {
	rootCmd.AddCommand(hello)
}
