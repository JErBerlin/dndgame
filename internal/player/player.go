// Package player manages the player entities withing the game.
package player

import "internal/character"

// PlayerStatus defines possible states of a player using an enumeration.
type PlayerStatus int

const (
	Inactive PlayerStatus = iota // when deactivated
	Active                       // after activation
)

// Player represents a player in the game.
type Player struct {
	PlayerID   string
	Name       string
	Status     string
	Characters []character.Character
}

// AddCharacter adds a new character to the player's list.
func (p *Player) AddCharacter(c character.Character) {
	// TODO: Implement adding a character the player's list
}

// RemoveCharacter removes a character from the player's list.
func (p *Player) RemoveCharacter(characterID string) {
	// TODO: Implement removing a character
}

// SetStatus changes the status of the player.
func (p *Player) SetStatus(newStatus PlayerStatus) {
	p.Status = newStatus
}
