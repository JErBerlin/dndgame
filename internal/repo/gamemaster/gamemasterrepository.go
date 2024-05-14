// internal/repo/gamemaster/gamemasterrepository.go

package gamemaster

import model "github.com/jerberlin/dndgame/internal/model/gamemaster"

type GameMasterRepository interface {
	GetGameMaster(id string) (*model.GameMaster, error)
	UpdateGameMaster(gm model.GameMaster) error
	CreateGameMaster(gm model.GameMaster) error
	DeleteGameMaster(id string) error
	ListGameMasters() ([]*model.GameMaster, error)
}
