// Package character manages player and non-player characters within the game.
package character

// CharacterClass defines common classes a character may belong to.
type CharacterClass int

const (
	Wizard CharacterClass = iota // Also called magic-user in D&D
	Warrior
	Cleric
	Ranger // Also called muntaner or montaraz
)

// CharacterRace defines common races a character may belong to as integer enums.
type CharacterRace int

const (
	Human CharacterRace = iota
	Elf
	Dwarf
	Orc
	Ghost
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
	Id          string
	Name        string
	Class       CharacterClass
	Race        CharacterRace
	Description string
	Attributes  Attributes
	XP          int
	Status      CharacterStatus
	GameId      string  // Foreign key to Game
	PlayerId    string  // Foreign key to Player
}

// NewCharacter creates a new character with specified attributes and characteristics.
func NewCharacter(id, name string, class CharacterClass, race CharacterRace, desc string, attrs Attributes, xp int, gameId, playerId string) *Character {
	return &Character{
		Id:          id,
		Name:        name,
		Class:       class,
		Race:        race,
		Description: desc,
		Attributes:  attrs,
		XP:          xp,
		Status:      Active, // default, can be changed as needed
		GameId:      gameId,
		PlayerId:    playerId,
	}
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
	classNames := [...]string{"Wizard", "Warrior", "Cleric", "Ranger"}
	return classNames[c]
}

// RaceString returns the string representation of the CharacterRace.
func (r CharacterRace) String() string {
	raceNames := [...]string{"Human", "Elf", "Dwarf", "Orc", "Ghost"}
	return raceNames[r]
}

// AddXP adds experience points to the character.
func (c *Character) AddXP(xp int) {
	c.XP += xp
}

// SubtractXP subtracts experience points from the character if possible.
func (c *Character) SubtractXP(xp int) {
	if c.XP - xp >= 0 {
		c.XP -= xp
	} else {
		c.XP = 0  // Ensure XP does not go negative.
	}
}