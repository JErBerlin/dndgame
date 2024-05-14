package player

import (
	"testing"
)

func TestSetStatus(t *testing.T) {
	p := Player{
		Id:     "player1",
		Name:   "John Doe",
		Status: Inactive,
	}
	p.SetStatus(Active)
	if p.Status != Active {
		t.Errorf("Failed to set player status, expected %v, got %v", Active, p.Status)
	}
}
