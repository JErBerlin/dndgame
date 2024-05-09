// internal/service/gamemaster/gamemasterservice.go

package gamemaster

import (
	"errors"

	"github.com/jerberlin/dndgame/internal/model/action"
	"github.com/jerberlin/dndgame/internal/model/character"
	"github.com/jerberlin/dndgame/internal/repo/action"
	"github.com/jerberlin/dndgame/internal/repo/character"
	"github.com/jerberlin/dndgame/internal/repo/game"
	"github.com/jerberlin/dndgame/internal/repo/gamemaster"
	"github.com/jerberlin/dndgame/internal/repo/player"
	"github.com/jerberlin/dndgame/internal/service/game"
	"github.com/jerberlin/dndgame/internal/service/player"
)

type GameMasterService interface {
	ListPendingActionInstances() ([]action.ActionInstance, error)
	ApproveActionInstance(instanceID string, modifiedInstance *action.ActionInstance) error
	ListActions() ([]action.Action, error)
	ModifyAction(actionID string, modifiedAction *action.Action) error
	ListCharacters() ([]character.Character, error)
	GetCharacter(characterID string) (*character.Character, error)
	UpdateCharacter(c *character.Character) error
	UpdateCharacterXP(characterID string, xpChange int) error
}

type service struct {
	actionRepo     action.ActionRepository
	characterRepo  character.CharacterRepository
	gameRepo       game.GameRepository
	gamemasterRepo gamemaster.GameMasterRepository
	playerRepo     player.PlayerRepository
	gameService    game.GameService
	playerService  player.PlayerService
}

var _ GameMasterService = &service{}

func NewGameMasterService(actionRepo action.ActionRepository, characterRepo character.CharacterRepository, gameRepo game.GameRepository, gamemasterRepo gamemaster.GameMasterRepository, playerRepo player.PlayerRepository, gameService game.GameService, playerService player.PlayerService) GameMasterService {
	return &service{
		actionRepo:     actionRepo,
		characterRepo:  characterRepo,
		gameRepo:       gameRepo,
		gamemasterRepo: gamemasterRepo,
		playerRepo:     playerRepo,
		gameService:    gameService,
		playerService:  playerService,
	}
}

// ListPendingActionInstances retrieves all action instances that haven't been approved yet.
func (s *service) ListPendingActionInstances() ([]action.ActionInstance, error) {
	return s.actionRepo.ListPendingInstances()
}

// ApproveActionInstance approves a specific action instance, with potential modifications.
func (s *service) ApproveActionInstance(instanceID string, modifiedInstance *action.ActionInstance) error {
	if modifiedInstance != nil {
		return s.actionRepo.UpdateActionInstance(modifiedInstance)
	}
	instance, err := s.actionRepo.GetActionInstanceByID(instanceID)
	if err != nil {
		return err
	}
	instance.Approved = true
	return s.actionRepo.UpdateActionInstance(instance)
}

// ListActions lists all actions available in the game.
func (s *service) ListActions() ([]action.Action, error) {
	return s.actionRepo.GetAllActions()
}

// ModifyAction modifies the details of an existing action.
func (s *service) ModifyAction(actionID string, modifiedAction *action.Action) error {
	return s.actionRepo.UpdateAction(modifiedAction)
}

// ListCharacters lists all characters in the game.
func (s *service) ListCharacters() ([]character.Character, error) {
	return s.characterRepo.GetAllCharacters()
}

// GetCharacter retrieves a single character by ID.
func (s *service) GetCharacter(characterID string) (*character.Character, error) {
	return s.characterRepo.GetCharacterByID(characterID)
}

// UpdateCharacter updates the details of a character.
func (s *service) UpdateCharacter(c *character.Character) error {
	return s.characterRepo.UpdateCharacter(c)
}

// UpdateCharacterXP modifies the XP of a character.
func (s *service) UpdateCharacterXP(characterID string, xpChange int) error {
	char, err := s.characterRepo.GetCharacterByID(characterID)
	if err != nil {
		return err
	}
	char.Attributes.XP += xpChange
	return s.characterRepo.UpdateCharacter(char)
}

// ReviewMissionProgress allows the Game Master to review and adjust the progress of missions within an adventure.
func (s *service) ReviewMissionProgress(gameID string, missionID string) error {
	// Fetch the game and its current adventure state.
	game, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return err
	}

	// Find the specific mission within the adventure to review.
	for _, m := range game.Adventure.Missions {
		if m.ID == missionID {
			// Here you could implement logic to review the progress or outcomes of the mission.
			// This might include checking if mission objectives are met or adjusting the mission status.
			return nil // Assuming no modifications are directly made here
		}
	}
	return errors.New("mission not found")
}

// StartAdventure initializes a new adventure within a game, setting up initial conditions and objectives.
func (s *service) StartAdventure(gameID string, adventure game.Adventure) error {
	// Fetch the game to start the adventure in.
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return err
	}

	// Set the adventure and initialize any required state or conditions.
	g.Adventure = adventure
	return s.gameRepo.UpdateGame(g)
}

// EndAdventure concludes an adventure within a game, potentially triggering game-end conditions or rewards.
func (s *service) EndAdventure(gameID string, adventureID string) error {
	// Fetch the game to end the adventure in.
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return err
	}

	// Verify the adventure to end matches the current one and perform any necessary cleanup or state updates.
	if g.Adventure.ID != adventureID {
		return errors.New("adventure mismatch or already ended")
	}

	// Clear or finalize the adventure state.
	g.Adventure = game.Adventure{} // Assuming a way to clear or reset the adventure.
	return s.gameRepo.UpdateGame(g)
}

// ManageNPCs allows for the addition, update, or removal of NPCs within a game, reflecting the GM's control.
func (s *service) ManageNPCs(gameID string, npc character.Character, operation string) error {
	// Fetch the game to manage NPCs in.
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return err
	}

	// Depending on the operation, add, update, or remove NPCs.
	switch operation {
	case "add":
		g.Characters = append(g.Characters, npc)
	case "update":
		for i, char := range g.Characters {
			if char.CharacterID == npc.CharacterID {
				g.Characters[i] = npc
				return s.gameRepo.UpdateGame(g)
			}
		}
		return errors.New("NPC not found for update")
	case "remove":
		for i, char := range g.Characters {
			if char.CharacterID == npc.CharacterID {
				g.Characters = append(g.Characters[:i], g.Characters[i+1:]...)
				return s.gameRepo.UpdateGame(g)
			}
		}
		return errors.New("NPC not found for removal")
	default:
		return errors.New("invalid operation")
	}

	return s.gameRepo.UpdateGame(g)
}

// SetAdventureOutcome allows the GM to define or update the outcome of an ongoing adventure, affecting the game state.
func (s *service) SetAdventureOutcome(gameID string, outcome string) error {
	// Fetch the game to set the adventure outcome.
	g, err := s.gameRepo.GetGameByID(gameID)
	if err != nil {
		return err
	}

	// Set or update the outcome of the current adventure.
	g.Adventure.Outcome = outcome
	return s.gameRepo.UpdateGame(g)
}
