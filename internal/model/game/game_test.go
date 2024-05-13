package game

import (
	"testing"
)

func TestSetStatus(t *testing.T) {
	g := Game{
		Status: Inactive,
	}
	g.SetStatus(Active)
	if g.Status != Active {
		t.Errorf("SetStatus failed, expected %v got %v", Active, g.Status)
	}
}

func TestSetAdventure(t *testing.T) {
	g := Game{}
	ad := Adventure{Id: "123Abc"}
	g.SetAdventure(ad.Id)
	if g.AdventureId != "123Abc" {
		t.Errorf("SetAdventure failed, expected adventure Id '123Abc', got '%v'", g.AdventureId)
	}
}
