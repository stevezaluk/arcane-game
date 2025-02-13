package game

import (
	"github.com/stevezaluk/arcane-game/models"
	"github.com/stevezaluk/mtgjson-models/card"
	"github.com/stevezaluk/mtgjson-models/user"
)

/*
CardObject - Represents the card played or generated for the game. This needs to differ
from the protobuf models as there are additional values that need to be tracked like
ownership, its parent zone, and the state of the card
*/
type CardObject struct {
	Metadata   *card.CardSet
	Owner      *user.User
	Controller *user.User
	ParentZone *Zone

	IsTapped          bool
	IsFaceDown        bool
	WasPlayedThisTurn bool
}

/*
Tap - Tap down the card object by setting IsTapped to true
*/
func (card *CardObject) Tap() {
	if !card.IsTapped {
		card.IsTapped = true
	}
}

/*
UnTap - UnTap the requested card if IsTapped is true
*/
func (card *CardObject) UnTap() {
	if card.IsTapped {
		card.IsTapped = false
	}
}

/*
NewCardObject - Create a new pointer to a card object. Its Metadata, Owner, and ParentZone are required
*/
func NewCardObject(metadata *card.CardSet, owner *Player, zone *Zone) *CardObject {
	return &CardObject{
		Metadata:          metadata,
		Owner:             owner,
		Controller:        owner,
		ParentZone:        zone,
		WasPlayedThisTurn: true,
	}
}
