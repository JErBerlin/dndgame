package character

import (
	"testing"
)

func TestNewCharacter(t *testing.T) {
	attrs := Attributes{Strength: 10, Dexterity: 12, Constitution: 14, Intelligence: 18, Wisdom: 13, Charisma: 11}
	xp := 10
	char := NewCharacter("char1", "Zanaphia Starfire", Wizard, Human, "A brilliant scholar", attrs, xp, "game-1", "player-1")

	if char.Name != "Zanaphia Starfire" || char.Class != Wizard || char.Race != Human || char.Status != Active {
		t.Errorf("NewCharacter failed to correctly initialize, got: %+v", char)
	}
}

func TestSetStatus(t *testing.T) {
	char := NewCharacter("char1", "Zanaphia Starfire", Wizard, Human, "A brilliant scholar", Attributes{}, 100, "game-1", "player-1")
	char.SetStatus(Inactive)
	if char.Status != Inactive {
		t.Errorf("SetStatus failed to update character status, expected Inactive, got: %v", char.Status)
	}
}

func TestUpdateAttributes(t *testing.T) {
	char := NewCharacter("char1", "Zanaphia Starfire", Wizard, Human, "A brilliant scholar", Attributes{Strength: 10}, 100, "game-1", "player-1")
	newAttrs := Attributes{Strength: 15}
	char.UpdateAttributes(newAttrs)

	if char.Attributes.Strength != 15 {
		t.Errorf("UpdateAttributes failed to update character attributes, expected 15, got: %d", char.Attributes.Strength)
	}
}
