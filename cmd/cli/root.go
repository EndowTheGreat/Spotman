package cli

import (
	"github.com/spf13/cobra"
	"gitlab.com/EndowTheGreat/spotman/internal/player"
)

var rootCmd = &cobra.Command{
	Use: "spotman",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Setup() error {
	var err error
	app.Player, err = player.NewPlayer()
	if err != nil {
		return err
	}
	defer app.Player.Conn.Close()
	if err := app.Player.FetchMetadata(); err != nil {
		return err
	}
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
