package action

import (
	"errors"
	"sync"
	"github.com/google/uuid"
	model "github.com/jerberlin/dndgame/internal/model/action"
)

// Check implementation
var _ ActionRepository = &InMemoryActionRepository{}

type InMemoryActionRepository struct {
	actions         map[string]*model.Action
	actionInstances map[string]*model.ActionInstance
	mutex           sync.RWMutex
}

func NewInMemoryActionRepository() *InMemoryActionRepository {
	return &InMemoryActionRepository{
		actions:         make(map[string]*model.Action),
		actionInstances: make(map[string]*model.ActionInstance),
	}
}

func (r *InMemoryActionRepository) CreateAction(a model.Action) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if a.ActionId == "" {
		a.ActionId = uuid.New().String() // Automatically generate an ID if not provided
	}
	if _, exists := r.actions[a.ActionId]; exists {
		return errors.New("action already exists")
	}
	r.actions[a.ActionId] = &a
	return nil
}

func (r *InMemoryActionRepository) UpdateAction(a model.Action) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.actions[a.ActionId]; !exists {
		return errors.New("action not found")
	}
	r.actions[a.ActionId] = &a
	return nil
}

func (r *InMemoryActionRepository) DeleteAction(actionId string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.actions[actionId]; !exists {
		return errors.New("action not found")
	}
	delete(r.actions, actionId)
	return nil
}

func (r *InMemoryActionRepository) GetActionByID(actionId string) (*model.Action, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if a, exists := r.actions[actionId]; exists {
		return a, nil
	}
	return nil, errors.New("action not found")
}

func (r *InMemoryActionRepository) CreateActionInstance(ai model.ActionInstance) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if ai.Id == "" {
		ai.Id = uuid.New().String() // Automatically generate an ID if not provided
	}
	if _, exists := r.actionInstances[ai.Id]; exists {
		return errors.New("action instance already exists")
	}
	r.actionInstances[ai.Id] = &ai
	return nil
}

func (r *InMemoryActionRepository) UpdateActionInstance(ai model.ActionInstance) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.actionInstances[ai.Id]; !exists {
		return errors.New("action instance not found")
	}
	r.actionInstances[ai.Id] = &ai
	return nil
}

func (r *InMemoryActionRepository) GetActionInstanceByID(instanceID string) (*model.ActionInstance, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if ai, exists := r.actionInstances[instanceID]; exists {
		return ai, nil
	}
	return nil, errors.New("action instance not found")
}

func (r *InMemoryActionRepository) ListActions() ([]*model.Action, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allActions := make([]*model.Action, 0, len(r.actions))
	for _, a := range r.actions {
		allActions = append(allActions, a)
	}
	return allActions, nil
}

func (r *InMemoryActionRepository) ListActionInstances() ([]*model.ActionInstance, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allInstances := make([]*model.ActionInstance, 0, len(r.actionInstances))
	for _, instance := range r.actionInstances {
		allInstances = append(allInstances, instance)
	}
	return allInstances, nil
}

func (r *InMemoryActionRepository) ListReadyActionsByCharacter(characterId string) ([]*model.ActionInstance, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	readyActions := make([]*model.ActionInstance, 0)
	for _, instance := range r.actionInstances {
		if instance.CharacterId == characterId && instance.Approved && !instance.Performed {
			readyActions = append(readyActions, instance)
		}
	}
	return readyActions, nil
}
