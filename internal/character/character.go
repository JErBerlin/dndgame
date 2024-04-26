// Package character manages player and non-player characters within the game.
package character

import "internal/action"

// GameStatus defines possible states of a game
type CharacterStatus int

const (
	Inactive CharacterStatus = iota // when deactivated
	Active                          // after activation
)

// Character represents both player-controlled and non-player characters in the game.
type Character struct {
	CharacterID      string
	Name             string
	XP               int
	Type             string
	Status           CharacterStatus
	ActionsInstances []action.ActionInstance
}

// PerformAction adds a action instance to the character's list of actions
func (c *Character) PerformAction(a action.ActionInstance) {
	// TODO: Implement action performing logic
}

// SetStatus changes the status of the character.
func (c *Character) SetStatus(newStatus CharacterStatus) {
	c.Status = newStatus
}
