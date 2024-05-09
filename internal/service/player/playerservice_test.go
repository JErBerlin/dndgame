package player

import (
	"os"
	"testing"

	"github.com/jerberlin/dndgame/internal/model/character"
	playermodel "github.com/jerberlin/dndgame/internal/model/player"
	"github.com/jerberlin/dndgame/internal/repo/player"
	playerrepo "github.com/jerberlin/dndgame/internal/repo/player"
)

var repo playerrepo.PlayerRepository
var service PlayerService

func TestMain(m *testing.M) {
	repo = playerrepo.NewInMemoryPlayerRepository()
	service = NewPlayerService(repo)

	os.Exit(m.Run())
}

func setupPlayer(repo player.PlayerRepository, playerID, playerName string) *player.Player {
	newPlayer := &playermodel.Player{
		PlayerID: playerID,
		Name:     playerName,
	}
	_ = repo.CreatePlayer(newPlayer)
	return newPlayer
}

func assertPlayerExistence(t *testing.T, playerID string, shouldExist bool) {
	_, err := repo.GetPlayerByID(playerID)
	if shouldExist && err != nil {
		t.Errorf("Expected player %s to exist, but got error: %v", playerID, err)
	} else if !shouldExist && err == nil {
		t.Errorf("Expected player %s not to exist, but it does", playerID)
	}
}

func TestCreatePlayer(t *testing.T) {
	playerID := "test-player-1"
	playerName := "Test Player"

	err := service.CreatePlayer(playerID, playerName)
	if err != nil {
		t.Errorf("CreatePlayer() error = %v, wantErr nil", err)
	}

	// Check if the player is created
	assertPlayerExistence(t, playerID, true)
}

func TestDeletePlayer(t *testing.T) {
	playerID := "test-player-2"
	playerName := "Test Player 2"
	setupPlayer(repo, playerID, playerName)

	err := service.DeletePlayer(playerID)
	if err != nil {
		t.Errorf("DeletePlayer() error = %v, wantErr nil", err)
	}

	// Check if the player is deleted
	assertPlayerExistence(t, playerID, false)
}

func TestAddCharacterToPlayer(t *testing.T) {
	playerID := "test-player-3"
	playerName := "Test Player 3"
	setupPlayer(repo, playerID, playerName)

	character := character.Character{CharacterID: "char1", Name: "Hero"}
	err := service.AddCharacterToPlayer(playerID, character)
	if err != nil {
		t.Errorf("AddCharacterToPlayer() error = %v, wantErr nil", err)
	}

	// Retrieve player to check character assignment
	p, _ := repo.GetPlayerByID(playerID)
	if len(p.Characters) != 1 || p.Characters[0].CharacterID != "char1" {
		t.Errorf("AddCharacterToPlayer() failed to add character, characters found: %v", p.Characters)
	}
}

func TestRemoveCharacterFromPlayer(t *testing.T) {
	playerID := "test-player-4"
	playerName := "Test Player 4"
	character := character.Character{CharacterID: "char2", Name: "Hero 2"}
	p := setupPlayer(repo, playerID, playerName)
	p.Characters = append(p.Characters, character)
	repo.UpdatePlayer(playerID, p)

	err := service.RemoveCharacterFromPlayer(playerID, "char2")
	if err != nil {
		t.Errorf("RemoveCharacterFromPlayer() error = %v, wantErr nil", err)
	}

	// Retrieve player to check character removal
	p, _ = repo.GetPlayerByID(playerID)
	if len(p.Characters) != 0 {
		t.Errorf("RemoveCharacterFromPlayer() failed to remove character, characters left: %v", p.Characters)
	}
}

func TestPerformActionByCharacter(t *testing.T) {
	playerID := "test-player-5"
	playerName := "Test Player 5"
	p := setupPlayer(repo, playerID, playerName) // Correctly set up the player once

	character := character.Character{CharacterID: "char3", Name: "Adventurer"}
	p.Characters = append(p.Characters, character) // Add character directly to the fetched player object
	repo.UpdatePlayer(playerID, p)                 // Update the player in the repository with the new character

	actionID := "action1"

	// Test performing an action by the character
	err := service.PerformActionByCharacter("char3", actionID)
	if err != nil {
		t.Errorf("PerformActionByCharacter() error = %v, wantErr nil", err)
	}

	// Additional checks can be added here if needed to verify the effects of the action

	// Test performing an action by a non-existent character
	err = service.PerformActionByCharacter("char-nonexistent", actionID)
	if err == nil {
		t.Errorf("PerformActionByCharacter() expected error for non-existent character, got nil")
	}

	// Test performing a non-existent action by an existing character
	err = service.PerformActionByCharacter("char3", "non-existent-action")
	if err == nil {
		t.Errorf("PerformActionByCharacter() expected error for non-existent action, got nil")
	}
}
