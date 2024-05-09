package game

import (
	"errors"

	"github.com/jerberlin/dndgame/internal/model/game"
	repogame "github.com/jerberlin/dndgame/internal/repo/game"
	"github.com/jerberlin/dndgame/internal/service/player"
	servplayer "github.com/jerberlin/dndgame/internal/service/player"
)

// GameService defines the interface for game-related operations.
type GameService interface {
	StartGame(gameID string) error
	EndGame(gameID string) error
	SetGameStatus(gameID string, status game.GameStatus) error
	AddPlayerToGame(gameID string, playerID string) error
	RemovePlayerFromGame(gameID string, playerID string) error
	SetAdventure(gameID string, adventure game.Adventure) error
	AddMissionToGame(gameID string, mission game.Mission) error
}

type service struct {
	gameRepo   repogame.GameRepository
	playerServ servplayer.PlayerService
}

// Ensure service implements GameService at compile time.
var _ GameService = &service{}

// NewGameService creates a new instance of GameService.
func NewGameService(gr game.GameRepository, ps player.PlayerService) GameService {
	return &service{
		gameRepo:   gr,
		playerServ: ps,
	}
}

// StartGame starts a new game session.
func (s *service) StartGame(gameID string) error {
	_, err := s.gameRepo.GetGameByID(gameID)
	if err == nil {
		return errors.New("game already exists")
	}
	newGame := &game.Game{
		GameID: gameID,
		Status: game.Active,
	}
	return s.gameRepo.CreateGame(newGame)
}

// EndGame ends a specific game session.
func (s *service) EndGame(gameID string) error {
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return errors.New("game not found")
	}
	g.Status = game.Inactive
	return s.gameRepo.UpdateGame(gameID, g)
}

// SetGameStatus updates the game's status.
func (s *service) SetGameStatus(gameID string, status game.GameStatus) error {
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return errors.New("game not found")
	}
	g.Status = status
	return s.gameRepo.UpdateGame(gameID, g)
}

// AddPlayerToGame adds a player to a game.
func (s *service) AddPlayerToGame(gameID string, playerID string) error {
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return errors.New("game not found")
	}
	p, err := s.playerServ.GetPlayerByID(playerID)
	if err != nil {
		return err // Player does not exist or other errors
	}
	return g.AddPlayer(p)
}

// RemovePlayerFromGame removes a player from a game.
func (s *service) RemovePlayerFromGame(gameID string, playerID string) error {
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return errors.New("game not found")
	}
	return g.RemovePlayer(playerID)
}

// SetAdventure sets the adventure for a specific game.
func (s *service) SetAdventure(gameID string, adventure game.Adventure) error {
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return errors.New("game not found")
	}
	g.SetAdventure(adventure)
	return s.gameRepo.UpdateGame(gameID, g)
}

// AddMissionToGame adds a mission to the current adventure in the game.
func (s *service) AddMissionToGame(gameID string, mission game.Mission) error {
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return errors.New("game not found")
	}
	g.AddMission(mission)
	return s.gameRepo.UpdateGame(gameID, g)
}
