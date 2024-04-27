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
	p := player.Player{PlayerID: "p1", Name: "Thething"}
	g.AddPlayer(p)
	if len(g.Players) != 1 || g.Players[0].PlayerID != "p1" {
		t.Errorf("AddPlayer failed, expected player with ID 'p1', got '%v'", g.Players[0].PlayerID)
	}
}

func TestRemovePlayer(t *testing.T) {
	g := Game{
		Players: []player.Player{{PlayerID: "p1", Name: "Kenny"}, {PlayerID: "p2", Name: "Doe"}},
	}
	err := g.RemovePlayer("p1")
	if err != nil || len(g.Players) != 1 || g.Players[0].PlayerID != "p2" {
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
	c := character.Character{CharacterID: "c1", Name: "Krillin"}
	g.AddCharacter(c)
	if len(g.Characters) != 1 || g.Characters[0].CharacterID != "c1" {
		t.Errorf("AddCharacter failed, expected character with ID 'c1', got '%v'", g.Characters[0].CharacterID)
	}
}

func TestAddAction(t *testing.T) {
	g := Game{}
	a := action.Action{ActionID: "a1", Name: "Strike"}
	g.AddAction(a)
	if len(g.Actions) != 1 || g.Actions[0].ActionID != "a1" {
		t.Errorf("AddAction failed, expected action with ID 'a1', got '%v'", g.Actions[0].ActionID)
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

func TestSetMission(t *testing.T) {
	g := Game{
		Adventure: Adventure{Type: Quests},
	}
	m := Mission{Name: "Rescue", Description: "Save the villagers"}
	g.SetMission(m)
	if g.Adventure.Mission.Name != "Rescue" {
		t.Errorf("SetMission failed, expected mission 'Rescue', got '%v'", g.Adventure.Mission.Name)
	}
}
