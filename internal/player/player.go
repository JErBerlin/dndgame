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
	p.Characters = append(p.Characters, c)
}

// RemoveCharacter removes a character from the player's list by ID.
func (p *Player) RemoveCharacter(characterID string) {
	for i, char := range p.Characters {
		if char.CharacterID == characterID {
			p.Characters = append(p.Characters[:i], p.Characters[i+1:]...)
			return
		}
	}
	// TODO: Handle error if character not found
}

// SetStatus changes the status of the player.
func (p *Player) SetStatus(newStatus PlayerStatus) {
	p.Status = newStatus
}
