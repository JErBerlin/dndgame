package player

import (
	"errors"

	"github.com/jerberlin/dndgame/internal/model/character"
	"github.com/jerberlin/dndgame/internal/model/player"
	repoplayer "github.com/jerberlin/dndgame/internal/repo/player"
)

// PlayerService defines the operations available to manage players.
type PlayerService interface {
	CreatePlayer(playerID, playerName string) error
	DeletePlayer(playerID string) error
	AddCharacterToPlayer(playerID string, character character.Character) error
	RemoveCharacterFromPlayer(playerID, characterID string) error
	PerformActionByCharacter(characterID, actionID string) error
}

type service struct {
	repo repoplayer.PlayerRepository
}

// Ensure service implements PlayerService at compile time.
var _ PlayerService = &service{}

// NewPlayerService creates a new instance of PlayerService.
func NewPlayerService(repo repoplayer.PlayerRepository) PlayerService {
	return &service{repo: repo}
}

func (s *service) CreatePlayer(playerID, playerName string) error {
	if _, err := s.repo.GetPlayerByID(playerID); err == nil {
		return errors.New("player already exists")
	}
	newPlayer := &player.Player{
		PlayerID: playerID,
		Name:     playerName,
	}
	return s.repo.CreatePlayer(newPlayer)
}

func (s *service) DeletePlayer(playerID string) error {
	return s.repo.DeletePlayer(playerID)
}

func (s *service) AddCharacterToPlayer(playerID string, character character.Character) error {
	p, err := s.repo.GetPlayerByID(playerID)
	if err != nil {
		return err
	}
	for _, ch := range p.Characters {
		if ch.CharacterID == character.CharacterID {
			return errors.New("character already assigned to this player")
		}
	}
	p.Characters = append(p.Characters, character)
	return s.repo.UpdatePlayer(playerID, p)
}

func (s *service) RemoveCharacterFromPlayer(playerID, characterID string) error {
	p, err := s.repo.GetPlayerByID(playerID)
	if err != nil {
		return err
	}
	for i, ch := range p.Characters {
		if ch.CharacterID == characterID {
			p.Characters = append(p.Characters[:i], p.Characters[i+1:]...)
			return s.repo.UpdatePlayer(playerID, p)
		}
	}
	return errors.New("character not found in player's list")
}

func (s *service) PerformActionByCharacter(characterID, actionID string) error {
	// Implementation depends on GameMasterService for action approval.
	return errors.New("method not implemented")
}
