package player

import (
	model "github.com/jerberlin/dndgame/internal/model/player"
)

// PlayerRepository defines the interface for player data operations.
type PlayerRepository interface {
	CreatePlayer(p model.Player) error
	UpdatePlayer(playerId string, p model.Player) error
	DeletePlayer(playerId string) error
	GetPlayerById(playerId string) (*model.Player, error)
	ListPlayers() ([]*model.Player, error)
}
