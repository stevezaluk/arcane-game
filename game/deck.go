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
