package character

import (
	"errors"
	"sync"

	"github.com/jerberlin/dndgame/internal/model/character"
)

type InMemoryCharacterRepository struct {
	characters map[string]*character.Character
	mutex      sync.RWMutex
}

func NewInMemoryCharacterRepository() *InMemoryCharacterRepository {
	return &InMemoryCharacterRepository{
		characters: make(map[string]*character.Character),
	}
}

func (r *InMemoryCharacterRepository) CreateCharacter(c *character.Character) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.characters[c.Id]; exists {
		return errors.New("character already exists")
	}
	r.characters[c.Id] = c
	return nil
}

func (r *InMemoryCharacterRepository) UpdateCharacter(c *character.Character) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if existingChar, exists := r.characters[c.Id]; !exists {
		return errors.New("character not found")
	} else {
		// Update the character's attributes and XP.
		existingChar.Attributes = c.Attributes
		existingChar.XP = c.XP
		existingChar.Status = c.Status // TODO: full overwrite?
		existingChar.ActionInstances = c.ActionInstances // TODO: full overwrite?
	}
	return nil
}


func (r *InMemoryCharacterRepository) DeleteCharacter(characterId string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, exists := r.characters[characterId]; !exists {
		return errors.New("character not found")
	}
	delete(r.characters, characterId)
	return nil
}

func (r *InMemoryCharacterRepository) GetCharacterByID(characterId string) (*character.Character, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if character, exists := r.characters[characterId]; exists {
		return character, nil
	}
	return nil, errors.New("character not found")
}

// ListCharacters retrieves all characters stored in the repository.
func (r *InMemoryCharacterRepository) ListCharacters() ([]*character.Character, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allCharacters := make([]*character.Character, 0, len(r.characters))
	for _, char := range r.characters {
		allCharacters = append(allCharacters, char)
	}
	return allCharacters, nil
}
