package player

import (
	"os"
	"testing"

	"github.com/jerberlin/dndgame/internal/model/character"
	"github.com/jerberlin/dndgame/internal/model/player"
	repoAction "github.com/jerberlin/dndgame/internal/repo/action"
	repoPlayer "github.com/jerberlin/dndgame/internal/repo/player"
)

var repoTest repoPlayer.PlayerRepository
var actionRepo repoAction.ActionRepository
var serviceTest PlayerService

// TestMain is used as a setup and teardown function that runs before and after all test functions.
// This function is called once for initializing resources needed for all tests in this package.
func TestMain(m *testing.M) {
	repoTest = repoPlayer.NewInMemoryPlayerRepository()
	actionRepo = repoAction.NewInMemoryActionRepository()
	serviceTest = NewPlayerService(repoTest, actionRepo)

	os.Exit(m.Run())
}

func setupPlayer(playerId, playerName string) *player.Player {
	newPlayer := player.Player{
		Id:   playerId,
		Name: playerName,
	}
	_ = repoTest.CreatePlayer(newPlayer)
	return &newPlayer
}

func assertPlayerExistence(t *testing.T, playerId string, shouldExist bool) {
	_, err := repoTest.GetPlayerById(playerId)
	if shouldExist && err != nil {
		t.Errorf("Expected player %s to exist, but got error: %v", playerId, err)
	} else if !shouldExist && err == nil {
		t.Errorf("Expected player %s not to exist, but it does", playerId)
	}
}

func TestCreatePlayer(t *testing.T) {
	playerId := "test-player-1"
	playerName := "Test Player"

	err := serviceTest.CreatePlayer(playerId, playerName)
	if err != nil {
		t.Errorf("CreatePlayer() error = %v, wantErr nil", err)
	}

	// Check if the player is created
	assertPlayerExistence(t, playerId, true)
}

func TestDeletePlayer(t *testing.T) {
	playerId := "test-player-2"
	playerName := "Test Player 2"
	setupPlayer(playerId, playerName)

	err := serviceTest.DeletePlayer(playerId)
	if err != nil {
		t.Errorf("DeletePlayer() error = %v, wantErr nil", err)
	}

	// Check if the player is deleted
	assertPlayerExistence(t, playerId, false)
}

func TestAddCharacterToPlayer(t *testing.T) {
	playerId := "test-player-3"
	playerName := "Test Player 3"
	setupPlayer(playerId, playerName)

	char := character.Character{Id: "char1", Name: "Hero"}
	err := serviceTest.AddCharacterToPlayer(playerId, char)
	if err != nil {
		t.Errorf("AddCharacterToPlayer() error = %v, wantErr nil", err)
	}

	p, _ := repoTest.GetPlayerById(playerId)
	if len(p.Characters) != 1 || p.Characters[0].Id != "char1" {
		t.Errorf("AddCharacterToPlayer() failed to add character, characters found: %v", p.Characters)
	}
}

func TestRemoveCharacterFromPlayer(t *testing.T) {
	playerId := "test-player-4"
	playerName := "Test Player 4"
	setupPlayer(playerId, playerName)

	char := character.Character{Id: "char2", Name: "Hero 2"}
	serviceTest.AddCharacterToPlayer(playerId, char)

	err := serviceTest.RemoveCharacterFromPlayer(playerId, "char2")
	if err != nil {
		t.Errorf("RemoveCharacterFromPlayer() error = %v, wantErr nil", err)
	}

	p, _ := repoTest.GetPlayerById(playerId)
	if len(p.Characters) != 0 {
		t.Errorf("RemoveCharacterFromPlayer() failed to remove character, characters left: %v", p.Characters)
	}
}

/*
func TestPerformActionByCharacter(t *testing.T) {
	playerId := "test-player-5"
	playerName := "Test Player 5"
	setupPlayer(playerId, playerName)

	char := character.Character{Id: "char3", Name: "Adventurer"}
	serviceTest.AddCharacterToPlayer(playerId, char)

	actionId := "action1"
	err := serviceTest.PerformAction(actionId)
	if err != nil {
		t.Errorf("PerformAction() error = %v, wantErr nil", err)
	}

	err = serviceTest.PerformAction(actionId)
	if err == nil {
		t.Errorf("PerformAction() expected error for non-existent character, got nil")
	}

	err = serviceTest.PerformAction("non-existent-action")
	if err == nil {
		t.Errorf("PerformAction() expected error for non-existent action, got nil")
	}
}
*/