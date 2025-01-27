package game

import "github.com/stevezaluk/mtgjson-models/deck"

/*
DeckObject - Represents the deck zone that the player interacts with through the game
*/
type DeckObject struct {
	Metadata   *deck.Deck
	Owner      *Player
	Controller *Player
	Zone       *Zone

	IsTopCardRevealed bool
}

/*
NewDeck - Creates a pointer to the DeckObject struct marking the Owner and Controller as the
player that was passed in the constructor
*/
func NewDeck(deck *deck.Deck, owner *Player) *DeckObject {
	zone := NewZone(DeckZoneId, owner, false, false, true)

	// convert object to zone here

	return &DeckObject{
		Metadata:   deck,
		Owner:      owner,
		Controller: owner,
		Zone:       zone,
	}
}
