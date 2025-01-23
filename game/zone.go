package game

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
	Owner  *Player

	// cards here

	IsPublic  bool
	IsShared  bool
	IsOrdered bool
}

/*
NewZone - A constructor provided for creating a new Zone. An owner and isShared can not be declared as once,
as if a Zone is shared then it cannot have an owner
*/
func NewZone(zoneId string, owner *Player, isPublic bool, isShared bool, isOrdered bool) *Zone {
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
