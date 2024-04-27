// Package character manages player and non-player characters within the game.
package character

import "internal/action"

// CharacterClass defines common classes a character may belong to.
type CharacterClass string

const (
	Wizard  CharacterClass = "Wizard" // Also called magic-user in dnd.
	Warrior CharacterClass = "Warrior"
	Cleric  CharacterClass = "Cleric"
	Ranger  CharacterClass = "Ranger" // Also called muntaner or montaraz.
)

// CharacterRace defines common races a character may belong to.
type CharacterRace string

const (
	Human CharacterRace = "Human"
	Elf   CharacterRace = "Elf"
	Dwarf CharacterRace = "Dwarf"
	Orc   CharacterRace = "Orc"
	Ghost CharacterRace = "Ghost"
)

// Attributes represents the key characteristics of a character.
// Typically ranges from 3-18 (set by rolling 3d6 or by rolling 4d6 and dropping the lowest).
type Attributes struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

// GameStatus defines possible states of a game
type CharacterStatus int

const (
	Inactive CharacterStatus = iota // when deactivated
	Active                          // after activation
)

// Character represents both player-controlled and non-player characters in the game.
type Character struct {
	CharacterID     string
	Name            string
	Class           CharacterClass
	Race            CharacterRace
	Description     string
	Attributes      Attributes
	Status          CharacterStatus
	ActionInstances []action.ActionInstance
}

// NewCharacter creates a new character with specified attributes and characteristics.
func NewCharacter(id, name string, class CharacterClass, race CharacterRace, desc string, attrs Attributes) *Character {
	return &Character{
		CharacterID:     id,
		Name:            name,
		Class:           class,
		Race:            race,
		Description:     desc,
		Attributes:      attrs,
		Status:          Active, // default, can be changed as needed
		ActionInstances: []action.ActionInstance{},
	}
}

// ChooseAction makes the choice to perform an action by a character. The action needs approval to be effectively executed.
func (c *Character) ChooseAction(act action.Action, customXPCost int) {
	actionInstance := act.CreateInstance(c.CharacterID, customXPCost)
	c.ActionInstances = append(c.ActionInstances, actionInstance)
	// TODO Optional: Notify game master for approval
}

// SetStatus changes the status of the character.
func (c *Character) SetStatus(newStatus CharacterStatus) {
	c.Status = newStatus
}

// UpdateAttributes updates the attributes of a character.
func (c *Character) UpdateAttributes(attrs Attributes) {
	c.Attributes = attrs
}

// ClassString returns the string representation of the CharacterClass.
func (c CharacterClass) String() string {
	return [...]string{"Fighter", "Wizard", "Cleric", "Rogue"}[c]
}

// RaceString returns the string representation of the CharacterRace.
func (r CharacterRace) String() string {
	return [...]string{"Human", "Dwarf", "Elf", "Halfling"}[r]
}
