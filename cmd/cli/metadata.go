package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var attribute string

func init() {
	rootCmd.AddCommand(metadataCmd)
	metadataCmd.Flags().StringVar(&attribute, "attribute", "", "Metadata attribute to fetch")
}

var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Fetch the metadata for the current song",
	Run: func(cmd *cobra.Command, args []string) {
		if attribute == "" {
			fmt.Printf("Artist: %v\nTitle: %v\nAlbum: %v\nImage: %v\nStatus: %v\n", app.Player.Current.Artist, app.Player.Current.Title, app.Player.Current.Album, app.Player.Current.Image, app.Player.Current.Status)
			return
		}
		switch attribute {
		case "artist":
			fmt.Println(app.Player.Current.Artist)
		case "title":
			fmt.Println(app.Player.Current.Title)
		case "album":
			fmt.Println(app.Player.Current.Album)
		case "image":
			fmt.Println(app.Player.Current.Image)
		case "status":
			fmt.Println(app.Player.Current.Status)
		default:
			fmt.Printf("Invalid attribute: %v\n", attribute)
		}
	},
}
