// Package action manages the action templates and instances withing the game.
package action

// Action represents a template for possible actions in the game.
type Action struct {
	ActionId   string
	Name       string
	BaseXPCost int
	GameId     string // Foreign key to Game
}

// ActionInstance represents a specific action taken by a character, customized to them and to a given scenario.
// The action will be chosen by the player of the character but has to be approved by the game master.
type ActionInstance struct {
	Id           string
	ActionId     string // Foreign key to Action
	CharacterId  string // Foreign key to Character
	CustomXPCost int
	Approved     bool
	Performed    bool
}

// NewAction creates a new action with specified details and associated game.
func NewAction(actionId, name string, baseXPCost int, gameId string) *Action {
	return &Action{
		ActionId:   actionId,
		Name:       name,
		BaseXPCost: baseXPCost,
		GameId:     gameId,
	}
}

// NewActionInstance creates a new action instance customized for a character.
func NewActionInstance(id, actionId, characterId string, customXPCost int) *ActionInstance {
	return &ActionInstance{
		Id:           id,
		ActionId:     actionId,
		CharacterId:  characterId,
		CustomXPCost: customXPCost,
		Approved:     false, // by default not approved
		Performed:    false, // and not performed yet
	}
}
