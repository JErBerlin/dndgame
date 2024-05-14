package action

import (
	"testing"
)

func TestActionCreation(t *testing.T) {
	action := NewAction("a1", "Fireball", 50, "game123")

	if action.ActionId != "a1" || action.Name != "Fireball" || action.BaseXPCost != 50 || action.GameId != "game123" {
		t.Errorf("Action creation failed, got: %+v", action)
	}
}

func TestNewActionInstance(t *testing.T) {
	instance := NewActionInstance("inst1", "a1", "char123", 60)

	if instance.Id != "inst1" || instance.ActionId != "a1" || instance.CharacterId != "char123" || instance.CustomXPCost != 60 || instance.Approved || instance.Performed {
		t.Errorf("NewActionInstance did not properly initialize, got: %+v", instance)
	}
}
