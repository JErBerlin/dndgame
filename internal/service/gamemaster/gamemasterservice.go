package gamemaster

import (
	"github.com/jerberlin/dndgame/internal/model/action"
	"github.com/jerberlin/dndgame/internal/model/character"
	repoAction "github.com/jerberlin/dndgame/internal/repo/action"
	repoCharacter "github.com/jerberlin/dndgame/internal/repo/character"
	repoGame "github.com/jerberlin/dndgame/internal/repo/game"
	repoPlayer "github.com/jerberlin/dndgame/internal/repo/player"
	srvGame "github.com/jerberlin/dndgame/internal/service/game"
)

type GameMasterService interface {
	ListPendingActionInstances() ([]*action.ActionInstance, error)
	ApproveActionInstance(instanceId string, modifiedInstance *action.ActionInstance) error
	ListActions() ([]*action.Action, error)
	ModifyAction(actionId string, modifiedAction *action.Action) error
	ListCharacters() ([]*character.Character, error)
	GetCharacter(characterId string) (*character.Character, error)
	UpdateCharacter(c *character.Character) error
	UpdateCharacterXP(characterId string, xpChange int) error
}

type service struct {
	actionRepo     repoAction.ActionRepository
	characterRepo  repoCharacter.CharacterRepository
	gameRepo       repoGame.GameRepository
	playerRepo     repoPlayer.PlayerRepository
	gameService    srvGame.GameService
}

func NewGameMasterService(
	actionRepo repoAction.ActionRepository, 
	characterRepo repoCharacter.CharacterRepository, 
	gameRepo repoGame.GameRepository, 
	playerRepo repoPlayer.PlayerRepository, 
	gameService srvGame.GameService ) GameMasterService {
	return &service{
		actionRepo:     actionRepo,
		characterRepo:  characterRepo,
		gameRepo:       gameRepo,
		playerRepo:     playerRepo,
		gameService:    gameService,
	}
}

func (s *service) ListPendingActionInstances() ([]*action.ActionInstance, error) {
	return s.actionRepo.ListActionInstances()
}

func (s *service) ApproveActionInstance(instanceId string, modifiedInstance *action.ActionInstance) error {
	if modifiedInstance != nil {
		return s.actionRepo.UpdateActionInstance(modifiedInstance)
	}
	instance, err := s.actionRepo.GetActionInstanceByID(instanceId)
	if err != nil {
		return err
	}
	instance.Approved = true
	return s.actionRepo.UpdateActionInstance(instance)
}

func (s *service) ListActions() ([]*action.Action, error) {
	return s.actionRepo.ListActions()
}

func (s *service) ModifyAction(actionId string, modifiedAction *action.Action) error {
	return s.actionRepo.UpdateAction(modifiedAction)
}

func (s *service) ListCharacters() ([]*character.Character, error) {
	return s.characterRepo.ListCharacters()
}

func (s *service) GetCharacter(characterId string) (*character.Character, error) {
	return s.characterRepo.GetCharacterByID(characterId)
}

func (s *service) UpdateCharacter(c *character.Character) error {
	return s.characterRepo.UpdateCharacter(c)
}

func (s *service) UpdateCharacterXP(characterId string, xpChange int) error {
	char, err := s.characterRepo.GetCharacterByID(characterId)
	if err != nil {
		return err
	}
	char.XP += xpChange

	return s.characterRepo.UpdateCharacter(char)
}
