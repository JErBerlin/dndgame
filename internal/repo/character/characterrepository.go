// internal/repo/character/characterrepository.go

package character

import "github.com/jerberlin/dndgame/internal/model/character"

type CharacterRepository interface {
	CreateCharacter(c *character.Character) error
	UpdateCharacter(c *character.Character) error
	DeleteCharacter(characterID string) error
	GetCharacterByID(characterID string) (*character.Character, error)
	ListCharacters() ([]*character.Character, error)
}
