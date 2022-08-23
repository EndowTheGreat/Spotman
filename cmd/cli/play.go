package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play the current song",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		app.Player.Play()
	},
}
