// Package game_master presents the game master agent responsible for directing the game.
package game_master

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

// ApproveAction approves an action instance chosen by a player to be executed by the players character.
func (gm *GameMaster) ApproveAction(instanceID string) {
	ai.Approved = true // For now, it simply approves the action.
}

// SetStatus changes the status of the game master.
func (gm *GameMaster) SetStatus(newStatus GameMasterStatus) {
	gm.Status = newStatus
}
