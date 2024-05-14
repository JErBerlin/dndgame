// internal/repo/action/actionrepository.go

package action

import model "github.com/jerberlin/dndgame/internal/model/action"

type ActionRepository interface {
	CreateAction(a model.Action) error
	UpdateAction(a model.Action) error
	DeleteAction(actionId string) error
	GetActionByID(actionId string) (*model.Action, error)
	CreateActionInstance(ai model.ActionInstance) error
	UpdateActionInstance(ai model.ActionInstance) error
	GetActionInstanceByID(instanceID string) (*model.ActionInstance, error)
	ListActions() ([]*model.Action, error)                 // Retrieve all actions defined in the game.
	ListActionInstances() ([]*model.ActionInstance, error) // Retrieve all action instances.
	ListReadyActionsByCharacter(characterId string) ([]*model.ActionInstance, error)
}