package game

import (
	"github.com/stevezaluk/mtgjson-models/user"
	"net"
)

/*
Player - Represents a player interacting with the game
*/
type Player struct {
	User *user.User
	Conn net.Conn

	Graveyard *Zone
	Hand      *Zone

	// mana pool

	LifeTotal          int
	CommanderDamage    int
	PoisonCounters     int
	EnergyCounters     int
	ExperienceCounters int

	IsMonarch   bool
	IsGameOwner bool
}
