package game

import (
	"github.com/stevezaluk/arcane-game-server/models"
	"github.com/stevezaluk/mtgjson-models/deck"
	"github.com/stevezaluk/mtgjson-models/user"
)

/*
DeckObject - Represents the deck zone that the player interacts with through the game
*/
type DeckObject struct {
	Metadata   *deck.Deck
	Owner      *user.User
	Controller *user.User
	Zone       *Zone

	IsTopCardRevealed bool
}

/*
NewDeckObject - Creates a pointer to the DeckObject struct marking the Owner and Controller as the
player that was passed in the constructor
*/
func NewDeckObject(deck *deck.Deck, owner *user.User) *DeckObject {
	zone := NewZone(DeckZoneId, owner, false, false, true)

	return &DeckObject{
		Metadata:   deck,
		Owner:      owner,
		Controller: owner,
		Zone:       zone,
	}
}

/*
Protobuf - Convert a DeckObject to it's protobuf representation
*/
func (deck *DeckObject) Protobuf() *models.DeckObject {
	return &models.DeckObject{
		Name:              deck.Metadata.Name,
		Owner:             deck.Owner.Email,
		Controller:        deck.Controller.Email,
		Zone:              deck.Zone.Protobuf(),
		IsTopCardRevealed: deck.IsTopCardRevealed,
	}
}
