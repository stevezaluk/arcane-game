package game

import (
	"github.com/stevezaluk/arcane-game-server/models"
	"github.com/stevezaluk/mtgjson-models/user"
	"net"
)

const (
	/*
		Constants used for identifying a players current phase
	*/

	BeginningPhaseId  = "phase:beginning"
	FirstMainPhaseId  = "phase:first-main"
	CombatPhaseId     = "phase:combat"
	SecondMainPhaseId = "phase:second-main"
	EndPhaseId        = "phase:end"

	/*
		Constants used for identifying a players current step
	*/

	UntapStepId             = "step:untap"
	UpkeepStepId            = "step:upkeep"
	DrawStepId              = "step:draw"
	BeginningOfCombatStepId = "step:beginning-of-combat"
	DeclareAttackersStepId  = "step:declare-attackers"
	DeclareBlockersStepId   = "step:declare-blockers"
	FirstStrikeDamageStepId = "step:first-strike-damage"
	DamageStepId            = "step:damage"
	EndOfCombatStepId       = "step:end-of-combat"
	EndStepId               = "step:end"
	CleanupStepId           = "step:cleanup"
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

	CurrentPhase string
	CurrentStep  string
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

	graveyard := NewZone(GraveyardZoneId, player.User, true, false, true)
	hand := NewZone(HandZoneId, player.User, false, false, false)

	player.Graveyard = graveyard
	player.Hand = hand

	return player
}

/*
Protobuf - Convert a Player object to it's protobuf representation
*/
func (player *Player) Protobuf() *models.Player {
	return &models.Player{
		Email:              player.User.Email,
		Deck:               player.Library.Protobuf(),
		Graveyard:          player.Graveyard.Protobuf(),
		Hand:               player.Hand.Protobuf(),
		LifeTotal:          int64(player.LifeTotal),
		CommanderDamage:    int64(player.CommanderDamage),
		PoisonCounters:     int64(player.PoisonCounters),
		EnergyCounters:     int64(player.EnergyCounters),
		ExperienceCounters: int64(player.ExperienceCounters),
		IsMonarch:          player.IsMonarch,
		IsGameOwner:        player.IsGameOwner,
		CurrentPhase:       player.CurrentPhase,
		CurrentStep:        player.CurrentStep,
	}
}
