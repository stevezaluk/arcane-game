package game

import (
	"github.com/stevezaluk/mtgjson-models/user"
)

type Zone struct {
	ZoneId string
	Owner  *user.User

	// cards here

	IsPublic  bool
	IsShared  bool
	IsOrdered bool
}
