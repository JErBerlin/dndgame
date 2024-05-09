package game

import (
	"errors"
	"sync"

	"github.com/jerberlin/dndgame/internal/model/game"
)

type InMemoryGameRepository struct {
	games map[string]*game.Game
	mutex sync.RWMutex
}

func NewInMemoryGameRepository() *InMemoryGameRepository {
	return &InMemoryGameRepository{
		games: make(map[string]*game.Game),
	}
}

func (r *InMemoryGameRepository) CreateGame(g *game.Game) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.games[g.GameID]; exists {
		return errors.New("game already exists")
	}
	r.games[g.GameID] = g
	return nil
}

func (r *InMemoryGameRepository) UpdateGame(gameID string, g *game.Game) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.games[gameID]; !exists {
		return errors.New("game not found")
	}
	r.games[gameID] = g
	return nil
}

func (r *InMemoryGameRepository) DeleteGame(gameID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.games[gameID]; !exists {
		return errors.New("game not found")
	}
	delete(r.games, gameID)
	return nil
}

func (r *InMemoryGameRepository) GetGameByID(gameID string) (*game.Game, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if game, exists := r.games[gameID]; exists {
		return game, nil
	}
	return nil, errors.New("game not found")
}

// ListGames retrieves all games stored in the repository.
func (r *InMemoryGameRepository) ListGames() ([]*game.Game, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allGames := make([]*game.Game, 0, len(r.games))
	for _, game := range r.games {
		allGames = append(allGames, game)
	}
	return allGames, nil
}
