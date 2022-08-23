package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playPauseCmd)
}

var playPauseCmd = &cobra.Command{
	Use:   "playpause",
	Short: "Toggle play/pause on the current song",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		app.Player.PlayPause()
	},
}
