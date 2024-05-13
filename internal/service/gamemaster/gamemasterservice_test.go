package gamemaster

import (
	"os"
	"testing"

	"github.com/jerberlin/dndgame/internal/model/action"
	"github.com/jerberlin/dndgame/internal/model/character"

	repoAction "github.com/jerberlin/dndgame/internal/repo/action"
	repoCharacter "github.com/jerberlin/dndgame/internal/repo/character"
	repoGame "github.com/jerberlin/dndgame/internal/repo/game"
	repoPlayer "github.com/jerberlin/dndgame/internal/repo/player"
	srvGame "github.com/jerberlin/dndgame/internal/service/game"
	srvPlayer "github.com/jerberlin/dndgame/internal/service/player"
)

var actionRepo repoAction.ActionRepository
var characterRepo repoCharacter.CharacterRepository
var gameRepo repoGame.GameRepository
var playerRepo repoPlayer.PlayerRepository
var gameService srvGame.GameService
var playerService srvPlayer.PlayerService
var gmService GameMasterService

func TestMain(m *testing.M) {
	// Create repositories
	actionRepo = repoAction.NewInMemoryActionRepository()
	characterRepo = repoCharacter.NewInMemoryCharacterRepository()
	gameRepo = repoGame.NewInMemoryGameRepository()
	playerRepo = repoPlayer.NewInMemoryPlayerRepository()

	// Create services
	gameService = srvGame.NewGameService(gameRepo)

	// Create GameMaster service
	gmService = NewGameMasterService(actionRepo, characterRepo, gameRepo, playerRepo, gameService)

	os.Exit(m.Run())
}

/*
func TestGameMasterServiceApproveActionInstance(t *testing.T) {
	instanceId := "instance1"
	ai := &action.ActionInstance{
		CharacterId:  "char1",
		CustomXPCost: 10,
		Approved:     false,
	}
	actionRepo.CreateActionInstance(ai)

	if err := gmService.ApproveActionInstance(instanceId, nil); err != nil {
		t.Errorf("ApproveActionInstance() error = %v, wantErr nil", err)
	}

	updatedAI, err := actionRepo.GetActionInstanceByID(instanceId)
	if err != nil {
		t.Fatalf("GetActionInstanceByID() error = %v, wantErr nil", err)
	}
	if !updatedAI.Approved {
		t.Errorf("ApproveActionInstance() failed to approve the action instance")
	}
}
*/

func TestGameMasterServiceListActions(t *testing.T) {
	actions := []*action.Action{
		{ActionId: "a1", Name: "Action 1", BaseXPCost: 5},
		{ActionId: "a2", Name: "Action 2", BaseXPCost: 10},
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
		{Id: "c1", Name: "Character 1"},
		{Id: "c2", Name: "Character 2"},
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
	characterId := "char1"
	newCharacter := &character.Character{
		Id:   characterId,
		Name: "Hero",
	}
	characterRepo.CreateCharacter(newCharacter)

	retrievedCharacter, err := gmService.GetCharacter(characterId)
	if err != nil {
		t.Errorf("GetCharacter() error = %v, wantErr nil", err)
	}
	if retrievedCharacter.Id != characterId {
		t.Errorf("GetCharacter() got = %v, want %v", retrievedCharacter.Id, characterId)
	}
}

func TestGameMasterServiceUpdateCharacter(t *testing.T) {
	characterId := "char2"
	newCharacter := &character.Character{
		Id:   characterId,
		Name: "Hero",
	}
	characterRepo.CreateCharacter(newCharacter)
	newCharacter.Name = "Hero Updated"

	if err := gmService.UpdateCharacter(newCharacter); err != nil {
		t.Errorf("UpdateCharacter() error = %v, wantErr nil", err)
	}

	updatedCharacter, _ := characterRepo.GetCharacterByID(characterId)
	if updatedCharacter.Name != "Hero Updated" {
		t.Errorf("UpdateCharacter() failed to update, got = %v", updatedCharacter.Name)
	}
}

func TestGameMasterServiceUpdateCharacterXP(t *testing.T) {
	characterId := "char3"
	newCharacter := &character.Character{
		Id:   characterId,
		Name: "Hero",
		XP:   100,
	}
	characterRepo.CreateCharacter(newCharacter)

	if err := gmService.UpdateCharacterXP(characterId, 50); err != nil {
		t.Errorf("UpdateCharacterXP() error = %v, wantErr nil", err)
	}

	updatedCharacter, _ := characterRepo.GetCharacterByID(characterId)
	if updatedCharacter.XP != 150 {
		t.Errorf("UpdateCharacterXP() failed to update XP, got = %v", updatedCharacter.XP)
	}
}

func TestGameMasterServiceModifyAction(t *testing.T) {
	actionId := "action1"
	newAction := &action.Action{
		ActionId:   actionId,
		Name:       "Strike",
		BaseXPCost: 10,
	}
	actionRepo.CreateAction(newAction)
	newAction.BaseXPCost = 15

	if err := gmService.ModifyAction(actionId, newAction); err != nil {
		t.Errorf("ModifyAction() error = %v, wantErr nil", err)
	}

	updatedAction, _ := actionRepo.GetActionByID(actionId)
	if updatedAction.BaseXPCost != 15 {
		t.Errorf("ModifyAction() failed to update BaseXPCost, got = %v", updatedAction.BaseXPCost)
	}
}

