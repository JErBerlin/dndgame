package gamemaster

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	model "github.com/jerberlin/dndgame/internal/model/gamemaster"
)

type InMemoryGameMasterRepository struct {
	masters map[string]*model.GameMaster
	mutex   sync.RWMutex
}

func NewInMemoryGameMasterRepository() *InMemoryGameMasterRepository {
	return &InMemoryGameMasterRepository{
		masters: make(map[string]*model.GameMaster),
	}
}

func (r *InMemoryGameMasterRepository) GetGameMaster(id string) (*model.GameMaster, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if gm, exists := r.masters[id]; exists {
		return gm, nil
	}
	return nil, errors.New("game master not found")
}

func (r *InMemoryGameMasterRepository) UpdateGameMaster(gm model.GameMaster) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.masters[gm.Id]; !exists {
		return errors.New("game master not found")
	}
	r.masters[gm.Id] = &gm
	return nil
}

func (r *InMemoryGameMasterRepository) CreateGameMaster(gm model.GameMaster) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if gm.Id == "" {
		gm.Id = uuid.New().String() // Automatically generate a UUID if not provided
	}
	if _, exists := r.masters[gm.Id]; exists {
		return errors.New("game master already exists")
	}
	r.masters[gm.Id] = &gm
	return nil
}

func (r *InMemoryGameMasterRepository) DeleteGameMaster(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.masters[id]; !exists {
		return errors.New("game master not found")
	}
	delete(r.masters, id)
	return nil
}

// ListGameMasters retrieves all game masters stored in the repository.
func (r *InMemoryGameMasterRepository) ListGameMasters() ([]*model.GameMaster, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allGameMasters := make([]*model.GameMaster, 0, len(r.masters))
	for _, gm := range r.masters {
		allGameMasters = append(allGameMasters, gm)
	}
	return allGameMasters, nil
}
