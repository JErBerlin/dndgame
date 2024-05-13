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
	if _, exists := r.games[g.Id]; exists {
		return errors.New("game already exists")
	}
	r.games[g.Id] = g
	return nil
}

func (r *InMemoryGameRepository) UpdateGame(gameId string, g *game.Game) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.games[gameId]; !exists {
		return errors.New("game not found")
	}
	r.games[gameId] = g
	return nil
}

func (r *InMemoryGameRepository) DeleteGame(gameId string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.games[gameId]; !exists {
		return errors.New("game not found")
	}
	delete(r.games, gameId)
	return nil
}

func (r *InMemoryGameRepository) GetGameById(gameId string) (*game.Game, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if game, exists := r.games[gameId]; exists {
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
