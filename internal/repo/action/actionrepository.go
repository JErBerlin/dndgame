// internal/repo/action/actionrepository.go

package action

import "github.com/jerberlin/dndgame/internal/model/action"

type ActionRepository interface {
	CreateAction(a *action.Action) error
	UpdateAction(a *action.Action) error
	DeleteAction(actionId string) error
	GetActionByID(actionId string) (*action.Action, error)
	CreateActionInstance(ai *action.ActionInstance) error
	UpdateActionInstance(ai *action.ActionInstance) error
	GetActionInstanceByID(instanceID string) (*action.ActionInstance, error)
	ListActions() ([]*action.Action, error)                 // Retrieve all actions defined in the game.
	ListActionInstances() ([]*action.ActionInstance, error) // Retrieve all action instances.
	ListReadyActionsByCharacter(characterId string) ([]*action.ActionInstance, error)
}