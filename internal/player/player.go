package player

import (
	"errors"
	"fmt"

	"github.com/godbus/dbus/v5"
)

type Player struct {
	Conn    *dbus.Conn
	Current *Current
}

type Current struct {
	Artist string
	Title  string
	Album  string
	Image  string
	Status string
}

func NewPlayer() (*Player, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return &Player{}, errors.New(fmt.Sprintf("Failed to connect to session bus: %v", err))
	}
	return &Player{Conn: conn}, nil
}

func (player *Player) FetchMetadata() error {
	metadata, err := player.Conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2").GetProperty("org.mpris.MediaPlayer2.Player.Metadata")
	if err != nil {
		return errors.New("Spotify is not playing")
	}
	playing, err := player.Conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2").GetProperty("org.mpris.MediaPlayer2.Player.PlaybackStatus")
	if err != nil {
		return errors.New("Spotify is not playing")
	}
	player.Current = &Current{
		Artist: metadata.Value().(map[string]dbus.Variant)["xesam:artist"].Value().([]string)[0],
		Title:  metadata.Value().(map[string]dbus.Variant)["xesam:title"].Value().(string),
		Album:  metadata.Value().(map[string]dbus.Variant)["xesam:album"].Value().(string),
		Image:  metadata.Value().(map[string]dbus.Variant)["mpris:artUrl"].Value().(string),
		Status: playing.Value().(string),
	}
	return nil
}

func (player Player) PlayPause() {
	player.Conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2").Call("org.mpris.MediaPlayer2.Player.PlayPause", 0)
}

func (player Player) Play() {
	player.Conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2").Call("org.mpris.MediaPlayer2.Player.Play", 0)
}

func (player Player) Pause() {
	player.Conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2").Call("org.mpris.MediaPlayer2.Player.Pause", 0)
}

func (player Player) Next() {
	player.Conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2").Call("org.mpris.MediaPlayer2.Player.Next", 0)
}

func (player Player) Previous() {
	player.Conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2").Call("org.mpris.MediaPlayer2.Player.Previous", 0)
}
