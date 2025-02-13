package game

import "github.com/stevezaluk/mtgjson-models/user"

const (
	BattlefieldZoneId = "zone:battlefield"
	ExileZoneId       = "zone:exile"
	GraveyardZoneId   = "zone:graveyard"
	HandZoneId        = "zone:hand"
	DeckZoneId        = "zone:deck"
	CommanderZoneId   = "zone:commander"
)

/*
Zone - Represents a Zone (or container) where cards can be placed within the game
*/
type Zone struct {
	ZoneId string
	Owner  *user.User

	cards []*CardObject

	IsPublic  bool
	IsShared  bool
	IsOrdered bool
}

/*
NewZone - A constructor provided for creating a new Zone. An owner and isShared can not be declared as once,
as if a Zone is shared then it cannot have an owner
*/
func NewZone(zoneId string, owner *user.User, isPublic bool, isShared bool, isOrdered bool) *Zone {
	if owner != nil && isShared {
		return nil
	}

	return &Zone{
		ZoneId:    zoneId,
		Owner:     owner,
		IsPublic:  isPublic,
		IsShared:  isShared,
		IsOrdered: isOrdered,
	}
}

/*
All - Return all the cards currently placed within the Zone
*/
func (zone *Zone) All() []*CardObject {
	return zone.cards
}

/*
Size - Return the number of cards within the Zone
*/
func (zone *Zone) Size() int {
	return len(zone.cards)
}
