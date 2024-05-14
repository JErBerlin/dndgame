// Package player (model) manages the player entities.
package player

// PlayerStatus defines possible states of a player using an enumeration.
type PlayerStatus int

const (
	Inactive PlayerStatus = iota // when deactivated
	Active                       // after activation
)

// Player represents a player in the game, linked directly to a specific game via a foreign key.
type Player struct {
	Id     string       `json:"id"`
	Name   string       `json:"name"`
	Status PlayerStatus `json:"status"`
	GameId string       `json:"game_id"` // Foreign key to Game
}

// NewPlayer creates a new player with the given details.
func NewPlayer(id, name, gameId string) *Player {
	return &Player{
		Id:     id,
		Name:   name,
		Status: Active, // Default to active when created
		GameId: gameId,
	}
}

// SetStatus changes the status of the player.
func (p *Player) SetStatus(newStatus PlayerStatus) {
	p.Status = newStatus
}
