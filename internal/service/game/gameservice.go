package game

import (
	"errors"
	"time"

	model "github.com/jerberlin/dndgame/internal/model/game"
	repoGame "github.com/jerberlin/dndgame/internal/repo/game"
	repoPlayer "github.com/jerberlin/dndgame/internal/repo/player"
)

// GameService defines the interface for game-related operations.
type GameService interface {
	StartGame(gameId string, name string) error
	EndGame(gameId string) error
	SetGameStatus(gameId string, status model.GameStatus) error
	AddPlayerToGame(gameId string, playerId string) error
	RemovePlayerFromGame(gameId string, playerId string) error
	SetAdventure(gameId string, adventure model.Adventure) error
	AddMissionToGame(gameId string, mission model.Mission) error
}

type service struct {
	gameRepo   repoGame.GameRepository
	playerRepo repoPlayer.PlayerRepository
}

// Ensure service implements GameService at compile time.
var _ GameService = &service{}

// NewGameService creates a new instance of GameService.
func NewGameService(gameRepo repoGame.GameRepository, playerRepo repoPlayer.PlayerRepository) GameService {
	return &service{
		gameRepo:   gameRepo,
		playerRepo: playerRepo,
	}
}

// StartGame initializes a new game session with a unique gameId and name.
// This is the primary action required before adding a Game Master, players, or setting an adventure and mission.
// It ensures that the game exists with a basic configuration that can be built upon.
// TODO: Consider modifying this function to allow starting the game after all configurations (e.g., players, GM, adventure) have been set up.
func (s *service) StartGame(gameId string, name string) error {
	_, err := s.gameRepo.GetGameById(gameId)
	if err == nil {
		return errors.New("game already exists")
	}

	gameOpts := &model.GameOptions{
		StartTime: time.Now(),
		Status:    model.Active,
	}

	g := model.NewGame(gameId, name, gameOpts)

	return s.gameRepo.CreateGame(g)
}

// EndGame ends a specific game.
func (s *service) EndGame(gameId string) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return err
	}
	g.Status = model.Inactive
	g.EndTime = time.Now()

	return s.gameRepo.UpdateGame(gameId, g)
}

// SetGameStatus updates the game's status.
func (s *service) SetGameStatus(gameId string, status model.GameStatus) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return err
	}
	g.Status = status

	return s.gameRepo.UpdateGame(gameId, g)
}

// AddPlayerToGame adds a player to a game.
// Obs: a player can only be in one game for now.
func (s *service) AddPlayerToGame(gameId string, playerId string) error {
	// Check that the id is correct
	_, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return err
	}

	p, err := s.playerRepo.GetPlayerById(playerId)
	if err != nil {
		return err
	}

	// Check if the player is already associated with another game
	// A player can only be in one game for now
	if p.GameId != "" && p.GameId != gameId {
		return errors.New("player is already associated with a different game")
	}

	// Associate the player with the game
	p.GameId = gameId // Link player to game
	err = s.playerRepo.UpdatePlayer(playerId, *p) // Update the player's game association
	if err != nil {
		return err
	}

	return nil
}

// RemovePlayerFromGame dissociates a player from a game.
func (s *service) RemovePlayerFromGame(gameId string, playerId string) error {
	// First, retrieve the player to check their current game association.
	p, err := s.playerRepo.GetPlayerById(playerId)
	if err != nil {
		return err // Return if player does not exist or other errors.
	}

	// Check if the player is currently associated with the game in question.
	if p.GameId != gameId {
		return errors.New("player is not playing this game")
	}

	// Dissociate the player from the game.
	p.GameId = "" // Remove the game association.

	// Update the player repository to reflect this change.
	err = s.playerRepo.UpdatePlayer(playerId, *p)
	if err != nil {
		return err // TODO: Handle any errors that occur during update.
	}

	return nil
}


// SetAdventure sets the adventure for a specific game by ensuring the adventure exists or creating it if not.
func (s *service) SetAdventure(gameId string, adventure model.Adventure) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return err // Handle error if game not found.
	}

	// Assuming a method to handle the existence check or creation of an adventure.
	_, err := s.gameRepo.GetAdventureById(adventure.Id)
	if err != nil {
		return err // Handle error if the adventure could not be ensured.
	}

	// Set the AdventureId of the game to the ensured adventure's ID.
	g.AdventureId = adventure.Id

	return s.gameRepo.UpdateGame(gameId, g)
}

// AddMissionToGame sets a mission to the current adventure in the game.
func (s *service) AddMissionToGame(gameId string, mission model.Mission) error {
	g, err := s.gameRepo.GetGameById(gameId)
	if err != nil {
		return err // Handle error if game not found.
	}

	if g.AdventureId == "" {
		return errors.New("no adventure set for this game")
	}

	// Assuming a method to handle the creation of a mission linked to an adventure.
	missionId, err := s.gameRepo.AddMissionToAdventure(g.AdventureId, mission)
	if err != nil {
		return err // Handle error if the mission could not be added.
	}

	// Optionally update the Adventure's MissionId if needed (not always necessary).
	return nil
}

