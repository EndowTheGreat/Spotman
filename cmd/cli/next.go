package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nextCmd)
}

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Skip to the next song",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		app.Player.Next()
	},
}
