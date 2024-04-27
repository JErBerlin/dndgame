package player

import (
	"testing"

	"github.com/jerberlin/dndgame/internal/character"
)

func TestAddCharacter(t *testing.T) {
	p := Player{
		PlayerID:   "player1",
		Name:       "John Doe",
		Status:     Active,
		Characters: []character.Character{},
	}
	char := character.Character{
		CharacterID: "char1",
		Name:        "Hero",
	}
	p.AddCharacter(char)

	if len(p.Characters) != 1 || p.Characters[0].CharacterID != "char1" {
		t.Errorf("Failed to add character to player, expected 1 character, got %d", len(p.Characters))
	}
}

func TestRemoveCharacter(t *testing.T) {
	char1 := character.Character{CharacterID: "char1", Name: "Hero1"}
	char2 := character.Character{CharacterID: "char2", Name: "Hero2"}
	p := Player{
		PlayerID:   "player1",
		Name:       "John Doe",
		Status:     Active,
		Characters: []character.Character{char1, char2},
	}

	p.RemoveCharacter("char1")

	if len(p.Characters) != 1 || p.Characters[0].CharacterID != "char2" {
		t.Errorf("Failed to remove character from player, expected 1 character left, got %d", len(p.Characters))
	}

	// Test removing a non-existent character
	err := p.RemoveCharacter("char3")
	if err == nil {
		t.Errorf("Expected an error when trying to remove a non-existent character, but got none")
	}
}

func TestSetStatus(t *testing.T) {
	p := Player{
		PlayerID: "player1",
		Name:     "John Doe",
		Status:   Inactive,
	}
	p.SetStatus(Active)
	if p.Status != Active {
		t.Errorf("Failed to set player status, expected %v, got %v", Active, p.Status)
	}
}
