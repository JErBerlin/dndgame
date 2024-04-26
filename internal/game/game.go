// Package game manage the overall game state and interactions within the game
package game

import (
	"internal/action"
	"internal/character"
	"internal/game_master"
	"internal/player"
	"time"
)

// GameStatus defines possible states of a game
type GameStatus int

const (
	Inactive GameStatus = iota // when deactivated
	Active                     // after activation
)

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
	Player     []player.Player
	Characters []character.Character
	GameMaster game_master.GameMaster
	Actions    []action.Action
}

// AddPlayer adds a new player to the game.
func (g *Game) AddPlayer(p player.Player) {
	// TODO: Implement adding a player to the game
}

// RemovePlayer removes a player from the game.
func (g *Game) RemovePlayer(playerID string) {
	// TODO: Implement removing a player
}

// AddCharacter adds a new character to the game.
func (g *Game) AddCharacter(c character.Character) {
	// TODO: Implement adding a character to the game
}

// AddAction adds a new action template to the game.
func (g *Game) AddAction(a action.Action) {
	// TODO: Implement adding an action to the game
}
