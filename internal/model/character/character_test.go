package character

import (
	"testing"

	"github.com/jerberlin/dndgame/internal/model/action"
)

func TestNewCharacter(t *testing.T) {
	attrs := Attributes{Strength: 10, Dexterity: 12, Constitution: 14, Intelligence: 18, Wisdom: 13, Charisma: 11}
	char := NewCharacter("char1", "Zanaphia Starfire", Wizard, Human, "A brilliant scholar", attrs)

	if char.Name != "Zanaphia Starfire" || char.Class != Wizard || char.Race != Human || char.Status != Active {
		t.Errorf("NewCharacter failed to correctly initialize, got: %+v", char)
	}
}

func TestChooseAction(t *testing.T) {
	char := NewCharacter("char1", "Zanaphia Starfire", Wizard, Human, "A brilliant scholar", Attributes{})
	act := action.Action{ActionID: "act1", Name: "Firebolt", BaseXPCost: 5}
	char.ChooseAction(act, 10)

	if len(char.ActionInstances) != 1 || char.ActionInstances[0].CustomXPCost != 10 {
		t.Errorf("ChooseAction did not correctly add an action instance, got: %+v", char.ActionInstances)
	}
}

func TestSetStatus(t *testing.T) {
	char := NewCharacter("char1", "Zanaphia Starfire", Wizard, Human, "A brilliant scholar", Attributes{})
	char.SetStatus(Inactive)
	if char.Status != Inactive {
		t.Errorf("SetStatus failed to update character status, expected Inactive, got: %v", char.Status)
	}
}

func TestUpdateAttributes(t *testing.T) {
	char := NewCharacter("char1", "Zanaphia Starfire", Wizard, Human, "A brilliant scholar", Attributes{Strength: 10})
	newAttrs := Attributes{Strength: 15}
	char.UpdateAttributes(newAttrs)

	if char.Attributes.Strength != 15 {
		t.Errorf("UpdateAttributes failed to update character attributes, expected 15, got: %d", char.Attributes.Strength)
	}
}
