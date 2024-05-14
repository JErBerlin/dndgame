// Package game (model) manage the overall game state and interactions.
package game

import (
	"time"
)

// GameStatus defines possible states of a game
type GameStatus int

const (
	Inactive GameStatus = iota // when deactivated
	Active                     // after activation
)

// Adventure represents a specific type of game scenario.
type Adventure struct {
	Id        string
	Type      AdventureType
	MissionId string // Foreign key linking to Mission. Single mission adventures for now.
}

// AdventureType defines different kinds of adventures in the game.
type AdventureType int

const (
	DungeonCrawls AdventureType = iota // Exploring underground complexes filled with monsters, traps, and treasure.
	Quests                             // Embarking on missions to retrieve magical items, rescue characters, or defeat a villain.
	Campaigns                          // Longer adventures that could evolve over multiple gaming sessions.
)

func (a Adventure) TypeString() string {
	adventureNames := [...]string{"DungeonCrawls", "Quests", "Campaigns"}
	return adventureNames[a.Type]
}

// Mission represents a specific task or challenge within an adventure.
type Mission struct {
	Id          string
	Name        string
	Description string
}

// Game represents the game entity with its list of possible game actions.
// The Game is directed by a Game Master.
// The actions are chosen by Players for one of their characters.
// A game has at a given point either the status active or inactive.
type Game struct {
	Id           string
	Name         string
	StartTime    time.Time
	EndTime      time.Time
	Status       GameStatus
	GameMasterId string // Foreign key linking to GameMaster
	AdventureId  string // Foreign key linking to Adventure
}

type GameOptions struct {
	StartTime time.Time
	EndTime   time.Time
	Status    GameStatus
}

// NewGame creates a new game with provided id, name, and optional settings.
// GameOptions is optational but when given, only one GameOptions is supported.
func NewGame(id string, name string, opts ...*GameOptions) *Game {
	g := &Game{
		Id:        id,
		Name:      name,
		StartTime: time.Now(), // Default start time (can be overridden)
		Status:    Inactive,   // Default status
	}

	if len(opts) > 0 && opts[0] != nil {
		if !opts[0].StartTime.IsZero() {
			g.StartTime = opts[0].StartTime
		}
		if !opts[0].EndTime.IsZero() {
			g.EndTime = opts[0].EndTime
		}
		g.Status = opts[0].Status
	}

	return g
}

// SetStatus changes the status of the game.
func (g *Game) SetStatus(newStatus GameStatus) {
	g.Status = newStatus
}

// SetAdventure sets a new adventure to the game.
func (g *Game) SetAdventure(adventureId string) {
	g.AdventureId = adventureId
}
