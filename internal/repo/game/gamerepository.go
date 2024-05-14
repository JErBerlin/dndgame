// internal/repo/game/gamerepository.go
package game

import model "github.com/jerberlin/dndgame/internal/model/game"

// GameRepository defines the interface for game data operations.
type GameRepository interface {
	CreateGame(game model.Game) error
	UpdateGame(gameId string, game *model.Game) error
	DeleteGame(gameId string) error
	GetGameById(gameId string) (*model.Game, error)
	ListGames() ([]*model.Game, error)
}
