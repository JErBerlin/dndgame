package player

import (
	"errors"

	"github.com/jerberlin/dndgame/internal/model/action"
	"github.com/jerberlin/dndgame/internal/model/character"
	"github.com/jerberlin/dndgame/internal/model/player"
	repoAction "github.com/jerberlin/dndgame/internal/repo/action"
	// repoCharacter "github.com/jerberlin/dndgame/internal/repo/character"
	repoPlayer "github.com/jerberlin/dndgame/internal/repo/player"
	// servGameMaster "github.com/jerberlin/dndgame/internal/service/gamemaster"
)

// PlayerService defines the operations available to manage players.
type PlayerService interface {
	CreatePlayer(playerId, playerName string) error
	DeletePlayer(playerId string) error
	GetPlayerById(playerId string) (*player.Player, error)
	AddCharacterToPlayer(playerId string, character character.Character) error
	RemoveCharacterFromPlayer(playerId, characterId string) error
	PerformAction(instanceId string) error
}

type service struct {
    repo   repoPlayer.PlayerRepository
    actionRepo   repoAction.ActionRepository
}

// Ensure service implements PlayerService at compile time.
var _ PlayerService = &service{}

// NewPlayerService creates a new instance of PlayerService.
func NewPlayerService(playerRepo repoPlayer.PlayerRepository, actionRepo repoAction.ActionRepository) PlayerService {
	return &service{
		repo:         playerRepo,
		actionRepo:   actionRepo,
	}
}

func (s *service) CreatePlayer(playerId, playerName string) error {
	if _, err := s.repo.GetPlayerById(playerId); err == nil {
		return errors.New("player already exists")
	}
	newPlayer := player.Player{
		Id:   playerId,
		Name: playerName,
	}
	return s.repo.CreatePlayer(newPlayer)
}

func (s *service) DeletePlayer(playerId string) error {
	return s.repo.DeletePlayer(playerId)
}

func (s *service) AddCharacterToPlayer(playerId string, character character.Character) error {
	p, err := s.repo.GetPlayerById(playerId)
	if err != nil {
		return err
	}
	for _, ch := range p.Characters {
		if ch.Id == character.Id {
			return errors.New("character already assigned to this player")
		}
	}
	p.Characters = append(p.Characters, character)
	return s.repo.UpdatePlayer(playerId, *p)
}

func (s *service) RemoveCharacterFromPlayer(playerId, characterId string) error {
	p, err := s.repo.GetPlayerById(playerId)
	if err != nil {
		return err
	}
	for i, ch := range p.Characters {
		if ch.Id == characterId {
			p.Characters = append(p.Characters[:i], p.Characters[i+1:]...)
			return s.repo.UpdatePlayer(playerId, *p)
		}
	}
	return errors.New("character not found in player's list")
}

func (s *service) PerformActionByCharacter(characterId, actionId string) error {
	// Implementation depends on GameMasterService for action approval.
	return errors.New("method not implemented")
}

// GetPlayerById retrieves a player by their Id.
func (s *service) GetPlayerById(playerId string) (*player.Player, error) {
	return s.repo.GetPlayerById(playerId)
}

func (s *service) ChooseAction(characterId, actionId string, customXPCost int) error {
    a, err := s.actionRepo.GetActionByID(actionId)
    if err != nil {
        return err
    }

    ai := &action.ActionInstance{
        Action:       *a,
        CharacterId:  characterId,
        CustomXPCost: customXPCost,
        Approved:     false,
        Performed:    false,
    }
    return s.actionRepo.CreateActionInstance(ai)
}

func (s *service) GetReadyActions(characterId string) ([]*action.ActionInstance, error) {
    return s.actionRepo.ListReadyActionsByCharacter(characterId)
}

func (s *service) PerformAction(instanceId string) error {
    ai, err := s.actionRepo.GetActionInstanceByID(instanceId)
    if err != nil {
        return err
    }
    if !ai.Approved || ai.Performed {
        return errors.New("action instance not approved or already performed")
    }

    // TODO: Implement the logic to apply the action's effects here.

    ai.Performed = true  // Mark as performed..
    return s.actionRepo.UpdateActionInstance(ai) // ..or delete if that's the desired behavior.
}
