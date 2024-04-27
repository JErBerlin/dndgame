// Package action manages the action templates and instances withing the game.
package action

// Action represents a template for possible actions in the game.
type Action struct {
	ActionID   string
	Name       string
	BaseXPCost int
}

// ActionInstance represents a specific action taken by a character, customised to them and to a given scenario
// The action will be chosen by the player of the character but has to be approved by the game master.
type ActionInstance struct {
	Action       Action
	CharacterID  string
	CustomXPCost int
	Approved     bool
}

// CreateInstance creates a new action instance customized for a character.
func (a *Action) CreateInstance(characterID string, customXPCost int) ActionInstance {
	return ActionInstance{
		Action:       *a,
		CharacterID:  characterID,
		CustomXPCost: customXPCost,
		Approved:     false,
	}
}
