// Package game_master presents the game master agent responsible for directing the game.
package game_master

import (
	"errors"

	"github.com/jerberlin/dndgame/internal/model/action"
)

// GameMasterStatus defines possible states of a game master.
type GameMasterStatus int

const (
	Inactive GameMasterStatus = iota // when deactivated
	Active                           // after activation
)

// GameMaster represents the game master directing the game.
type GameMaster struct {
	GMID   string
	Name   string
	Status GameMasterStatus
}

// ApproveAction finds an action instance by ID and sets its Approved status to true. Returns an error if not found.
// TODO: decouple action instances from approve action function, f.ex. injecting the action manager service or with event-driven aproval.
func (gm *GameMaster) ApproveAction(instances map[string]*action.ActionInstance, instanceID string) error {
	if ai, exists := instances[instanceID]; exists {
		ai.Approved = true
		return nil
	}
	return errors.New("action instance not found")
}

// SetStatus changes the status of the game master.
func (gm *GameMaster) SetStatus(newStatus GameMasterStatus) {
	gm.Status = newStatus
}
