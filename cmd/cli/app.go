package cli

import "gitlab.com/EndowTheGreat/spotman/internal/player"

type App struct {
	Player *player.Player
}

var app App
