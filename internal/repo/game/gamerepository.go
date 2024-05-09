// internal/repo/game/gamerepository.go

package game

import "github.com/jerberlin/dndgame/internal/model/game"

// GameRepository defines the interface for game data operations.
type GameRepository interface {
	CreateGame(game *game.Game) error
	UpdateGame(gameID string, game *game.Game) error
	DeleteGame(gameID string) error
	GetGameByID(gameID string) (*game.Game, error)
	ListGames() ([]*game.Game, error)
}
