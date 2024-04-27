// Package game manage the overall game state and interactions within the game
package game

import (
	"fmt"
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