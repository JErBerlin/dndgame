package game

import (
	"testing"

	"github.com/jerberlin/dndgame/internal/model/action"
	"github.com/jerberlin/dndgame/internal/model/character"
	"github.com/jerberlin/dndgame/internal/model/player"
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

func TestAddPlayer(t *testing.T) {
	g := Game{}
	p := player.Player{Id: "p1", Name: "Thething"}
	g.AddPlayer(p)
	if len(g.Players) != 1 || g.Players[0].Id != "p1" {
		t.Errorf("AddPlayer failed, expected player with ID 'p1', got '%v'", g.Players[0].Id)
	}
}

func TestRemovePlayer(t *testing.T) {
	g := Game{
		Players: []player.Player{{Id: "p1", Name: "Kenny"}, {Id: "p2", Name: "Doe"}},
	}
	err := g.RemovePlayer("p1")
	if err != nil || len(g.Players) != 1 || g.Players[0].Id != "p2" {
		t.Errorf("RemovePlayer failed, expected player with ID 'p2', got '%v'", g.Players)
	}
	// Test for non-existent player
	err = g.RemovePlayer("p3")
	if err == nil {
		t.Errorf("Expected error when trying to remove non-existent player 'p3', but got none")
	}
}

func TestAddCharacter(t *testing.T) {
	g := Game{}
	c := character.Character{Id: "c1", Name: "Krillin"}
	g.AddCharacter(c)
	if len(g.Characters) != 1 || g.Characters[0].Id != "c1" {
		t.Errorf("AddCharacter failed, expected character with ID 'c1', got '%v'", g.Characters[0].Id)
	}
}

func TestAddAction(t *testing.T) {
	g := Game{}
	a := action.Action{ActionId: "a1", Name: "Strike"}
	g.AddAction(a)
	if len(g.Actions) != 1 || g.Actions[0].ActionId != "a1" {
		t.Errorf("AddAction failed, expected action with ID 'a1', got '%v'", g.Actions[0].ActionId)
	}
}

func TestSetAdventure(t *testing.T) {
	g := Game{}
	ad := Adventure{Type: Quests}
	g.SetAdventure(ad)
	if g.Adventure.Type != Quests {
		t.Errorf("SetAdventure failed, expected adventure type 'Quests', got '%v'", g.Adventure.Type)
	}
}

func TestAddMission(t *testing.T) {
	g := Game{
		Adventure: Adventure{Type: Quests},
	}
	m := Mission{Name: "Rescue", Description: "Save the villagers"}
	g.AddMission(m)
	if g.Adventure.Mission.Name != "Rescue" {
		t.Errorf("AddMission failed, expected mission 'Rescue', got '%v'", g.Adventure.Mission.Name)
	}
}
