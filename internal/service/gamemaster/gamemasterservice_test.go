package gamemaster

import (
	"os"
	"testing"

	"github.com/jerberlin/dndgame/internal/model/action"
	"github.com/jerberlin/dndgame/internal/model/character"
	repoaction "github.com/jerberlin/dndgame/internal/repo/action"
	repocharacter "github.com/jerberlin/dndgame/internal/repo/character"
	servgame "github.com/jerberlin/dndgame/internal/service/game"
)

var actionRepo repoaction.ActionRepository
var characterRepo repocharacter.CharacterRepository
var gameService servgame.GameService
var gmService GameMasterService

func TestMain(m *testing.M) {
	actionRepo = repoaction.NewInMemoryActionRepository()
	characterRepo = repocharacter.NewInMemoryCharacterRepository()
	gameService = servgame.NewGameService(nil) // Assuming dependencies are properly injected or mocked
	gmService = servgamemaster.NewGameMasterService(actionRepo, characterRepo, gameService)

	os.Exit(m.Run())
}

func TestGameMasterServiceApproveActionInstance(t *testing.T) {
	instanceID := "instance1"
	ai := &action.ActionInstance{
		CharacterID:  "char1",
		CustomXPCost: 10,
		Approved:     false,
	}
	actionRepo.CreateActionInstance(ai)

	if err := gmService.ApproveActionInstance(instanceID, nil); err != nil {
		t.Errorf("ApproveActionInstance() error = %v, wantErr nil", err)
	}

	updatedAI, err := actionRepo.GetActionInstanceByID(instanceID)
	if err != nil {
		t.Fatalf("GetActionInstanceByID() error = %v, wantErr nil", err)
	}
	if !updatedAI.Approved {
		t.Errorf("ApproveActionInstance() failed to approve the action instance")
	}
}

func TestGameMasterServiceListActions(t *testing.T) {
	actions := []*action.Action{
		{ActionID: "a1", Name: "Action 1", BaseXPCost: 5},
		{ActionID: "a2", Name: "Action 2", BaseXPCost: 10},
	}
	for _, act := range actions {
		actionRepo.CreateAction(act)
	}

	resultActions, err := gmService.ListActions()
	if err != nil {
		t.Fatalf("ListActions() error = %v, wantErr nil", err)
	}
	if len(resultActions) != len(actions) {
		t.Errorf("ListActions() got %v actions, want %v", len(resultActions), len(actions))
	}
}

func TestGameMasterServiceListCharacters(t *testing.T) {
	characters := []*character.Character{
		{CharacterID: "c1", Name: "Character 1"},
		{CharacterID: "c2", Name: "Character 2"},
	}
	for _, char := range characters {
		characterRepo.CreateCharacter(char)
	}

	resultCharacters, err := gmService.ListCharacters()
	if err != nil {
		t.Fatalf("ListCharacters() error = %v, wantErr nil", err)
	}
	if len(resultCharacters) != len(characters) {
		t.Errorf("ListCharacters() got %v characters, want %v", len(resultCharacters), len(characters))
	}
}

func TestGameMasterServiceGetCharacter(t *testing.T) {
	characterID := "char1"
	newCharacter := &character.Character{
		CharacterID: characterID,
		Name:        "Hero",
	}
	characterRepo.CreateCharacter(newCharacter)

	retrievedCharacter, err := gmService.GetCharacter(characterID)
	if err != nil {
		t.Errorf("GetCharacter() error = %v, wantErr nil", err)
	}
	if retrievedCharacter.CharacterID != characterID {
		t.Errorf("GetCharacter() got = %v, want %v", retrievedCharacter.CharacterID, characterID)
	}
}

func TestGameMasterServiceUpdateCharacter(t *testing.T) {
	characterID := "char2"
	newCharacter := &character.Character{
		CharacterID: characterID,
		Name:        "Hero",
	}
	characterRepo.CreateCharacter(newCharacter)
	newCharacter.Name = "Hero Updated"

	if err := gmService.UpdateCharacter(newCharacter); err != nil {
		t.Errorf("UpdateCharacter() error = %v, wantErr nil", err)
	}

	updatedCharacter, _ := characterRepo.GetCharacterByID(characterID)
	if updatedCharacter.Name != "Hero Updated" {
		t.Errorf("UpdateCharacter() failed to update, got = %v", updatedCharacter.Name)
	}
}

func TestGameMasterServiceUpdateCharacterXP(t *testing.T) {
	characterID := "char3"
	newCharacter := &character.Character{
		CharacterID: characterID,
		Name:        "Hero",
		Attributes: character.Attributes{
			XP: 100,
		},
	}
	characterRepo.CreateCharacter(newCharacter)

	if err := gmService.UpdateCharacterXP(characterID, 50); err != nil {
		t.Errorf("UpdateCharacterXP() error = %v, wantErr nil", err)
	}

	updatedCharacter, _ := characterRepo.GetCharacterByID(characterID)
	if updatedCharacter.Attributes.XP != 150 {
		t.Errorf("UpdateCharacterXP() failed to update XP, got = %v", updatedCharacter.Attributes.XP)
	}
}

func TestGameMasterServiceModifyAction(t *testing.T) {
	actionID := "action1"
	newAction := &action.Action{
		ActionID:   actionID,
		Name:       "Strike",
		BaseXPCost: 10,
	}
	actionRepo.CreateAction(newAction)
	newAction.BaseXPCost = 15

	if err := gmService.ModifyAction(actionID, newAction); err != nil {
		t.Errorf("ModifyAction() error = %v, wantErr nil", err)
	}

	updatedAction, _ := actionRepo.GetActionByID(actionID)
	if updatedAction.BaseXPCost != 15 {
		t.Errorf("ModifyAction() failed to update BaseXPCost, got = %v", updatedAction.BaseXPCost)
	}
}
