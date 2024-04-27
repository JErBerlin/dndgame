// Package game manage the overall game state and interactions within the game
package game

import (
	"fmt"
	"time"

	"github.com/jerberlin/dndgame/internal/model/action"
	"github.com/jerberlin/dndgame/internal/model/character"
	"github.com/jerberlin/dndgame/internal/model/game_master"
	"github.com/jerberlin/dndgame/internal/model/player"
)

// GameStatus defines possible states of a game
type GameStatus int

const (
	Inactive GameStatus = iota // when deactivated
	Active                     // after activation
)

// AdventureType defines different kinds of adventures in the game.
type AdventureType int

const (
	DungeonCrawls AdventureType = iota // Exploring underground complexes filled with monsters, traps, and treasure.
	Quests                             // Embarking on missions to retrieve magical items, rescue characters, or defeat a villain.
	Campaigns                          // Longer adventures that could evolve over multiple gaming sessions.
)

// Mission represents a specific task or challenge within an adventure.
type Mission struct {
	Name        string
	Description string
}

// Adventure represents a specific type of game scenario.
type Adventure struct {
	Type    AdventureType
	Mission Mission // singular mission adventure
}

// Game represents the game entity with its list of possible game actions.
// The Game is directed by a Game Master.
// The actions are chosen by Players for one of their characters.
// Each game has at cretion a defined start time and an end time.
// A game is at given point either the status active or inactive.
type Game struct {
	GameID     string
	Name       string
	StartTime  time.Time
	EndTime    time.Time
	Status     GameStatus
	Players    []player.Player
	Characters []character.Character
	GameMaster game_master.GameMaster
	Actions    []action.Action
	Adventure  Adventure // singular adventure
}

// SetStatus changes the status of the game.
func (g *Game) SetStatus(newStatus GameStatus) {
	g.Status = newStatus
}

// AddPlayer adds a new player to the game.
func (g *Game) AddPlayer(p player.Player) {
	g.Players = append(g.Players, p)
}

// RemovePlayer removes a player from the game by ID.
func (g *Game) RemovePlayer(playerID string) (err error) {
	for i, p := range g.Players {
		if p.PlayerID == playerID {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("player with ID %s not found", playerID)
}

// AddCharacter adds a new character to the game.
func (g *Game) AddCharacter(c character.Character) {
	g.Characters = append(g.Characters, c)
}

// AddAction adds a new action template to the game.
func (g *Game) AddAction(a action.Action) {
	g.Actions = append(g.Actions, a)
}

// SetAdventure sets a new adventure to the game.
func (g *Game) SetAdventure(adventure Adventure) {
	g.Adventure = adventure
}

// SetMission sets a mission for the current adventure.
func (g *Game) SetMission(mission Mission) {
	g.Adventure.Mission = mission
}
