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

	Library   *DeckObject
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

/*
NewPlayer - Declare a new Player object. Initializes the player controlled zones automatically
and tracks a reference to the Connection object
*/
func NewPlayer(user *user.User, library *DeckObject, conn net.Conn) *Player {
	player := &Player{
		User:      user,
		Conn:      conn,
		Library:   library,
		LifeTotal: 20,
	}

	graveyard := NewZone(GraveyardZoneId, player, true, false, true)
	hand := NewZone(HandZoneId, player, false, false, false)

	player.Graveyard = graveyard
	player.Hand = hand

	return player
}
