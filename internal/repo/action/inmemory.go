// internal/repo/action/inmemoryactionrepository.go

package action

import (
	"errors"
	"sync"

	"github.com/jerberlin/dndgame/internal/model/action"
)

type InMemoryActionRepository struct {
	actions         map[string]*action.Action
	actionInstances map[string]*action.ActionInstance
	mutex           sync.RWMutex
}

func NewInMemoryActionRepository() *InMemoryActionRepository {
	return &InMemoryActionRepository{
		actions:         make(map[string]*action.Action),
		actionInstances: make(map[string]*action.ActionInstance),
	}
}

func (r *InMemoryActionRepository) CreateAction(a *action.Action) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.actions[a.ActionID]; exists {
		return errors.New("action already exists")
	}
	r.actions[a.ActionID] = a
	return nil
}

func (r *InMemoryActionRepository) UpdateAction(a *action.Action) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.actions[a.ActionID]; !exists {
		return errors.New("action not found")
	}
	r.actions[a.ActionID] = a
	return nil
}

func (r *InMemoryActionRepository) DeleteAction(actionID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.actions[actionID]; !exists {
		return errors.New("action not found")
	}
	delete(r.actions, actionID)
	return nil
}

func (r *InMemoryActionRepository) GetActionByID(actionID string) (*action.Action, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if a, exists := r.actions[actionID]; exists {
		return a, nil
	}
	return nil, errors.New("action not found")
}

func (r *InMemoryActionRepository) CreateActionInstance(ai *action.ActionInstance) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.actionInstances[ai.CharacterID]; exists {
		return errors.New("action instance already exists")
	}
	r.actionInstances[ai.CharacterID] = ai
	return nil
}

func (r *InMemoryActionRepository) UpdateActionInstance(ai *action.ActionInstance) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.actionInstances[ai.CharacterID]; !exists {
		return errors.New("action instance not found")
	}
	r.actionInstances[ai.CharacterID] = ai
	return nil
}

func (r *InMemoryActionRepository) GetActionInstanceByID(instanceID string) (*action.ActionInstance, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if ai, exists := r.actionInstances[instanceID]; exists {
		return ai, nil
	}
	return nil, errors.New("action instance not found")
}

// ListActions returns all actions stored in the repository.
func (r *InMemoryActionRepository) ListActions() ([]*action.Action, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allActions := make([]*action.Action, 0, len(r.actions))
	for _, action := range r.actions {
		allActions = append(allActions, action)
	}
	return allActions, nil
}

// ListActionInstances returns all action instances stored in the repository.
func (r *InMemoryActionRepository) ListActionInstances() ([]*action.ActionInstance, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allInstances := make([]*action.ActionInstance, 0, len(r.actionInstances))
	for _, instance := range r.actionInstances {
		allInstances = append(allInstances, instance)
	}
	return allInstances, nil
}
