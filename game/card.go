package game

import "github.com/stevezaluk/mtgjson-models/card"

/*
CardObject - Represents the card played or generated for the game. This needs to differ
from the protobuf models as there are additional values that need to be tracked like
ownership, its parent zone, and the state of the card
*/
type CardObject struct {
	Metadata   *card.CardSet
	Owner      *Player
	Controller *Player
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
