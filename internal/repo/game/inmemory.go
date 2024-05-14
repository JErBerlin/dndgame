package game

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	model "github.com/jerberlin/dndgame/internal/model/game"
)

// Check implementation
var _ GameRepository = &InMemoryGameRepository{}

type InMemoryGameRepository struct {
	games map[string]*model.Game
	mutex sync.RWMutex
}

func NewInMemoryGameRepository() *InMemoryGameRepository {
	return &InMemoryGameRepository{
		games: make(map[string]*model.Game),
	}
}

func (r *InMemoryGameRepository) CreateGame(g model.Game) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if g.Id == "" {
		g.Id = uuid.New().String()  // Automatically generate a UUID if not provided
	}
	if _, exists := r.games[g.Id]; exists {
		return errors.New("game already exists with this ID")
	}
	r.games[g.Id] = &g
	return nil
}

func (r *InMemoryGameRepository) UpdateGame(gameId string, g *model.Game) error {
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

func (r *InMemoryGameRepository) GetGameById(gameId string) (*model.Game, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if game, exists := r.games[gameId]; exists {
		return game, nil
	}
	return nil, errors.New("game not found")
}

// ListGames retrieves all games stored in the repository.
func (r *InMemoryGameRepository) ListGames() ([]*model.Game, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allGames := make([]*model.Game, 0, len(r.games))
	for _, game := range r.games {
		allGames = append(allGames, game)
	}
	return allGames, nil
}
