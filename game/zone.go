package game

import (
	"github.com/stevezaluk/mtgjson-models/user"
)

type Zone struct {
	ZoneId string

	// Owner (*Player) here
	// cards here

	IsPublic  bool
	IsShared  bool
	IsOrdered bool
}

func NewZone(zoneId string, owner *user.User, isPublic bool, isShared bool, isOrdered bool) *Zone {
	if owner != nil && isShared {
		return nil
	}

	return &Zone{
		ZoneId:    zoneId,
		IsPublic:  isPublic,
		IsShared:  isShared,
		IsOrdered: isOrdered,
	}
}
