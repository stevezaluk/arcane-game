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
NewCardObject - Create a new pointer to a card object. Its Metadata, Owner, and ParentZone are required
*/
func NewCardObject(metadata *card.CardSet, owner *user.User, zone *Zone) *CardObject {
	return &CardObject{
		Metadata:          metadata,
		Owner:             owner,
		Controller:        owner,
		ParentZone:        zone,
		WasPlayedThisTurn: true,
	}
}

/*
Protobuf - Convert a CardObject to it's protobuf representation
*/
func (card *CardObject) Protobuf() *models.CardObject {
	return &models.CardObject{
		Name:              card.Metadata.Name,
		Description:       card.Metadata.Text,
		Type:              card.Metadata.Type,
		SubTypes:          card.Metadata.Subtypes,
		ColorIdentity:     card.Metadata.ColorIdentity,
		ConvertedManaCost: card.Metadata.ConvertedManaCost,
		Toughness:         card.Metadata.Toughness,
		Power:             card.Metadata.Power,
		IsTapped:          card.IsTapped,
		IsFaceDown:        card.IsFaceDown,
		WasPlayedThisTurn: card.WasPlayedThisTurn,
		Owner:             card.Owner.Email,
		Controller:        card.Owner.Email,
	}
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
