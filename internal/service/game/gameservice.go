package game

import (
	"errors"
	"time"

	"github.com/jerberlin/dndgame/internal/model/game"
	repoGame "github.com/jerberlin/dndgame/internal/repo/game"
	repoPlayer "github.com/jerberlin/dndgame/internal/repo/player"
)

// GameService defines the interface for game-related operations.
type GameService interface {
	StartGame(gameId string, name string) error
	EndGame(gameId string) error
	SetGameStatus(gameId string, status game.GameStatus) error
	AddPlayerToGame(gameId string, playerId string) error
	RemovePlayerFromGame(gameId string, playerId string) error
	SetAdventure(gameId string, adventure game.Adventure) error
	AddMissionToGame(gameId string, mission game.Mission) error
}

type service struct {
	gameRepo   repoGame.GameRepository
	playerRepo repoPlayer.PlayerRepository
}

// Ensure service implements GameService at compile time.
var _ GameService = &service{}

// NewGameService creates a new instance of GameService.
func NewGameService(gr repoGame.GameRepository) GameService {
	return &service{
		gameRepo: gr,
	}
}

// StartGame starts a new game session.
// The precondition is that no game with the given Id exists, since the game will be created at start.
func (s *service) StartGame(gameId string, name string) error {
	_, err := s.gameRepo.GetGameById(gameId)
	if err == nil {
		return errors.New("game already exists")
	}

	gameOpts := &game.GameOptions{
		StartTime: time.Now(),
		Status:    game.Active,
	}

	g := game.NewGame(gameId, name, gameOpts)

	return s.gameRepo.CreateGame(g)
}

// EndGame ends a specific game.
func (s *service) EndGame(gameId string) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return errors.New("game not found")
	}
	g.EndTime = time.Now()
	g.Status = game.Inactive

	return s.gameRepo.UpdateGame(gameId, g)
}

// SetGameStatus updates the game's status.
func (s *service) SetGameStatus(gameId string, status game.GameStatus) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return errors.New("game not found")
	}
	g.Status = status

	return s.gameRepo.UpdateGame(gameId, g)
}

// AddPlayerToGame adds a player to a game.
func (s *service) AddPlayerToGame(gameId string, playerId string) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return errors.New("game not found")
	}
	p, err := s.playerRepo.GetPlayerById(playerId)
	if err != nil {
		return err // Player does not exist or other errors
	}
	g.AddPlayer(*p)
	return s.gameRepo.UpdateGame(gameId, g)
}

// RemovePlayerFromGame removes a player from a game.
func (s *service) RemovePlayerFromGame(gameId string, playerId string) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return errors.New("game not found")
	}

	err = g.RemovePlayer(playerId)
	if err != nil {
		return err
	}

	return s.gameRepo.UpdateGame(gameId, g)
}

// SetAdventure sets the adventure for a specific game.
func (s *service) SetAdventure(gameId string, adventure game.Adventure) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return errors.New("game not found")
	}
	g.SetAdventure(adventure)

	return s.gameRepo.UpdateGame(gameId, g)
}

// AddMissionToGame adds a mission to the current adventure in the game.
func (s *service) AddMissionToGame(gameId string, mission game.Mission) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return errors.New("game not found")
	}
	g.AddMission(mission)

	return s.gameRepo.UpdateGame(gameId, g)
}
