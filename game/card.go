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
