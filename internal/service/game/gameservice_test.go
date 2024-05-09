package game

import (
	"os"
	"testing"

	"github.com/jerberlin/dndgame/internal/model/game"
	"github.com/jerberlin/dndgame/internal/repo/game"
)

var repo game.GameRepository
var service GameService

func TestMain(m *testing.M) {
	repo = game.NewInMemoryGameRepository()
	service = NewGameService(repo, nil)

	os.Exit(m.Run())
}

func assertGameStatus(t *testing.T, gameID string, expectedStatus game.GameStatus) {
	g, err := repo.GetGameByID(gameID)
	if err != nil {
		t.Fatalf("GetGameByID() error = %v, wantErr false", err)
	}
	if g.Status != expectedStatus {
		t.Errorf("Game status got = %v, want %v", g.Status, expectedStatus)
	}
}

func TestGameServiceStartAndEndGame(t *testing.T) {
	gameID := "test-game-1"

	// Test starting the game
	if err := service.StartGame(gameID); err != nil {
		t.Errorf("StartGame() error = %v, wantErr false", err)
	}
	assertGameStatus(t, gameID, game.Active)

	// Test ending the game
	if err := service.EndGame(gameID); err != nil {
		t.Errorf("EndGame() error = %v, wantErr false", err)
	}
	assertGameStatus(t, gameID, game.Inactive)
}

func TestGameServiceSetGameStatus(t *testing.T) {
	gameID := "test-game-status"
	setupGame(repo, gameID, game.Active)

	// Test setting game status to inactive
	if err := service.SetGameStatus(gameID, game.Inactive); err != nil {
		t.Errorf("SetGameStatus() error = %v, wantErr false", err)
	}
	assertGameStatus(t, gameID, game.Inactive)
}

func TestGameServiceAddPlayerToGame(t *testing.T) {
	gameID := "test-game-player"
	setupGame(repo, gameID, game.Active)

	playerID := "player123"
	if err := service.AddPlayerToGame(gameID, playerID); err != nil {
		t.Errorf("AddPlayerToGame() error = %v", err)
	}
}

func TestGameServiceRemovePlayerFromGame(t *testing.T) {
	gameID := "test-game-remove-player"
	setupGame(repo, gameID, game.Active)

	playerID := "player-to-remove"
	if err := service.RemovePlayerFromGame(gameID, playerID); err != nil {
		t.Errorf("RemovePlayerFromGame() error = %v, wantErr false", err)
	}
}

func TestGameServiceSetAdventure(t *testing.T) {
	gameID := "test-game-adventure"
	setupGame(repo, gameID, game.Active)

	adventure := game.Adventure{
		Type: game.Quests,
		Mission: game.Mission{
			Name:        "Retrieve the Lost Artifact",
			Description: "Players must navigate the ancient ruins to retrieve a lost artifact.",
		},
	}

	if err := service.SetAdventure(gameID, adventure); err != nil {
		t.Errorf("SetAdventure() error = %v, wantErr false", err)
	}
}

func TestGameServiceAddMissionToGame(t *testing.T) {
	gameID := "test-game-mission"
	setupGame(repo, gameID, game.Active)

	mission := game.Mission{
		Name:        "Defend the Village",
		Description: "Players must defend the village from a band of marauding goblins.",
	}

	if err := service.AddMissionToGame(gameID, mission); err != nil {
		t.Errorf("AddMissionToGame() error = %v, wantErr false", err)
	}
}

func setupGame(repo game.GameRepository, gameID string, status game.GameStatus) {
	newGame := &game.Game{
		GameID: gameID,
		Status: status,
	}
	repo.CreateGame(newGame)
}
