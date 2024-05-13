package game

import (
	"os"
	"testing"

	"github.com/jerberlin/dndgame/internal/model/game"
	repoGame "github.com/jerberlin/dndgame/internal/repo/game"
)

var repo repoGame.GameRepository
var srvGame GameService

func TestMain(m *testing.M) {
	repo = repoGame.NewInMemoryGameRepository()
	srvGame = NewGameService(repo)

	os.Exit(m.Run())
}

func assertGameStatus(t *testing.T, gameId string, expectedStatus game.GameStatus) {
	g, err := repo.GetGameById(gameId)
	if err != nil {
		t.Fatalf("GetGameById() error = %v, wantErr false", err)
	}
	if g.Status != expectedStatus {
		t.Errorf("Game status got = %v, want %v", g.Status, expectedStatus)
	}
}

/*
	func TestGameServiceStartAndEndGame(t *testing.T) {
		gameId := "test-game-1"
		name := "the quest for test"

		// Test starting the game
		if err := srvGame.StartGame(gameId, name); err != nil {
			t.Errorf("StartGame() error = %v, wantErr false", err)
		}
		assertGameStatus(t, gameId, game.Active)

		// Test ending the game
		if err := srvGame.EndGame(gameId); err != nil {
			t.Errorf("EndGame() error = %v, wantErr false", err)
		}
		assertGameStatus(t, gameId, game.Inactive)
	}
*/
func TestGameServiceSetGameStatus(t *testing.T) {
	gameId := "test-game-status"
	setupGame(repo, gameId, game.Active)

	// Test setting game status to inactive
	if err := srvGame.SetGameStatus(gameId, game.Inactive); err != nil {
		t.Errorf("SetGameStatus() error = %v, wantErr false", err)
	}
	assertGameStatus(t, gameId, game.Inactive)
}

func TestGameServiceAddPlayerToGame(t *testing.T) {
	gameId := "test-game-player"
	setupGame(repo, gameId, game.Active)

	playerId := "player123"
	if err := srvGame.AddPlayerToGame(gameId, playerId); err != nil {
		t.Errorf("AddPlayerToGame() error = %v", err)
	}
}

func TestGameServiceRemovePlayerFromGame(t *testing.T) {
	gameId := "test-game-remove-player"
	setupGame(repo, gameId, game.Active)

	playerId := "player-to-remove"
	if err := srvGame.RemovePlayerFromGame(gameId, playerId); err != nil {
		t.Errorf("RemovePlayerFromGame() error = %v, wantErr false", err)
	}
}

func TestGameServiceSetAdventure(t *testing.T) {
	gameId := "test-game-adventure"
	setupGame(repo, gameId, game.Active)

	adventure := game.Adventure{
		Type: game.Quests,
		Mission: game.Mission{
			Name:        "Retrieve the Lost Artifact",
			Description: "Players must navigate the ancient ruins to retrieve a lost artifact.",
		},
	}

	if err := srvGame.SetAdventure(gameId, adventure); err != nil {
		t.Errorf("SetAdventure() error = %v, wantErr false", err)
	}
}

func TestGameServiceAddMissionToGame(t *testing.T) {
	gameId := "test-game-mission"
	setupGame(repo, gameId, game.Active)

	mission := game.Mission{
		Name:        "Defend the Village",
		Description: "Players must defend the village from a band of marauding goblins.",
	}

	if err := srvGame.AddMissionToGame(gameId, mission); err != nil {
		t.Errorf("AddMissionToGame() error = %v, wantErr false", err)
	}
}

func setupGame(repo repoGame.GameRepository, gameId string, status game.GameStatus) {
	newGame := &game.Game{
		Id:     gameId,
		Status: status,
	}
	repo.CreateGame(newGame)
}
