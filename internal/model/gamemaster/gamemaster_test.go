package gamemaster

import (
	"testing"
)

func TestSetStatus(t *testing.T) {
	gm := NewGameMaster("gm1", "Game Master One", Inactive, "game123")

	// Change the status of the game master.
	gm.SetStatus(Active)

	if gm.Status != Active {
		t.Errorf("SetStatus failed: expected %v, got %v", Active, gm.Status)
	}
}

func TestSetGameId(t *testing.T) {
	gm := NewGameMaster("gm1", "Game Master One", Active, "game123")

	// Change the game ID of the game master.
	newGameId := "game456"
	gm.SetGameId(newGameId)

	if gm.GameId != newGameId {
		t.Errorf("SetGameId failed: expected %s, got %s", newGameId, gm.GameId)
	}
}
