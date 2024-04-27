package game_master

import (
	"testing"

	"github.com/jerberlin/dndgame/internal/action"
)

func TestApproveAction(t *testing.T) {
	gm := GameMaster{GMID: "GM1", Name: "Anakin", Status: Active}
	actionInstances := make(map[string]*action.ActionInstance)
	actionInstance := &action.ActionInstance{
		Action:       action.Action{ActionID: "act1", Name: "Firebolt", BaseXPCost: 10},
		CharacterID:  "char1",
		CustomXPCost: 15,
		Approved:     false,
	}
	actionInstances["act1"] = actionInstance

	// Test successful approval
	err := gm.ApproveAction(actionInstances, "act1")
	if err != nil {
		t.Errorf("ApproveAction failed unexpectedly: %v", err)
	}
	if !actionInstance.Approved {
		t.Error("ApproveAction failed to approve the action")
	}

	// Test failure to find action instance
	err = gm.ApproveAction(actionInstances, "nonexistent")
	if err == nil {
		t.Error("ApproveAction did not fail as expected when action instance is not found")
	}
}
