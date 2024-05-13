package player

import (
	"github.com/jerberlin/dndgame/internal/model/player"
)

// PlayerRepository defines the interface for player data operations.
type PlayerRepository interface {
	CreatePlayer(p player.Player) error
	UpdatePlayer(playerId string, p player.Player) error
	DeletePlayer(playerId string) error
	GetPlayerById(playerId string) (*player.Player, error)
	ListPlayers() ([]player.Player, error)
}
