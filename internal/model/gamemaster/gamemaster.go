// Package game_master presents the game master agent responsible for directing the game.
package gamemaster

// GameMasterStatus defines possible states of a game master.
type GameMasterStatus int

const (
	Inactive GameMasterStatus = iota // when deactivated
	Active                           // after activation
)

// GameMaster represents the game master directing the game.
type GameMaster struct {
	Id     string
	Name   string
	Status GameMasterStatus
	GameId string // Foreign key to Game
}

// NewGameMaster creates a new game master with provided details.
func NewGameMaster(id, name string, status GameMasterStatus, gameId string) *GameMaster {
	return &GameMaster{
		Id:     id,
		Name:   name,
		Status: status,
		GameId: gameId,
	}
}

// SetStatus changes the status of the game master.
func (gm *GameMaster) SetStatus(newStatus GameMasterStatus) {
	gm.Status = newStatus
}

// SetStatus changes the game id of the game master.
func (gm *GameMaster) SetGameId(newGameId string) {
	gm.GameId = newGameId
}
