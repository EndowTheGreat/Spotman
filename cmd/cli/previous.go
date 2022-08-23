package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(previousCmd)
}

var previousCmd = &cobra.Command{
	Use:   "previous",
	Short: "Skip to the previous song",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		app.Player.Previous()
	},
}
