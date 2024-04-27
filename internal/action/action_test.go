package action

import (
	"testing"
)

func TestActionCreation(t *testing.T) {
	action := Action{
		ActionID:   "a1",
		Name:       "Fireball",
		BaseXPCost: 50,
	}

	if action.ActionID != "a1" || action.Name != "Fireball" || action.BaseXPCost != 50 {
		t.Errorf("Action creation failed, got: %+v", action)
	}
}

func TestCreateInstance(t *testing.T) {
	action := Action{
		ActionID:   "a1",
		Name:       "Fireball",
		BaseXPCost: 50,
	}
	instance := action.CreateInstance("char123", 60)

	if instance.Action != action || instance.CharacterID != "char123" || instance.CustomXPCost != 60 || instance.Approved {
		t.Errorf("CreateInstance did not properly initialize, got: %+v", instance)
	}
}
