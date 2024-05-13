package player

import (
	"errors"
	"sync"

	"github.com/jerberlin/dndgame/internal/model/player"
)

type InMemoryPlayerRepository struct {
	players map[string]*player.Player
	mutex   sync.RWMutex
}

func NewInMemoryPlayerRepository() *InMemoryPlayerRepository {
	return &InMemoryPlayerRepository{
		players: make(map[string]*player.Player),
	}
}

func (r *InMemoryPlayerRepository) CreatePlayer(p player.Player) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.players[p.Id]; exists {
		return errors.New("player already exists")
	}
	r.players[p.Id] = &p
	return nil
}

func (r *InMemoryPlayerRepository) UpdatePlayer(playerId string, p player.Player) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.players[playerId]; !exists {
		return errors.New("player not found")
	}
	r.players[playerId] = &p
	return nil
}

func (r *InMemoryPlayerRepository) DeletePlayer(playerId string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.players[playerId]; !exists {
		return errors.New("player not found")
	}
	delete(r.players, playerId)
	return nil
}

func (r *InMemoryPlayerRepository) GetPlayerById(playerId string) (*player.Player, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if player, exists := r.players[playerId]; exists {
		return player, nil
	}
	return nil, errors.New("player not found")
}

// ListPlayers retrieves all players stored in the repository.
func (r *InMemoryPlayerRepository) ListPlayers() ([]player.Player, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allPlayers := make([]player.Player, 0, len(r.players))
	for _, p := range r.players {
		allPlayers = append(allPlayers, *p)
	}
	return allPlayers, nil
}
