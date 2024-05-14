// internal/repo/character/characterrepository.go
package character

import model "github.com/jerberlin/dndgame/internal/model/character"

type CharacterRepository interface {
	CreateCharacter(c model.Character) error
	UpdateCharacter(c model.Character) error
	DeleteCharacter(characterId string) error
	GetCharacterByID(characterId string) (*model.Character, error)
	ListCharacters() ([]*model.Character, error)
}
