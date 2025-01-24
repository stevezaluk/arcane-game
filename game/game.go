package game

import (
	"github.com/stevezaluk/mtgjson-sdk-client/api"
)

/*
Game - A representation of a single MTG Game. Process's game commands sent from the client
*/
type Game struct {
	Name     string
	GameMode string

	Players   []*Player
	GameOwner *Player

	Battlefield *Zone
	Exile       *Zone
	Command     *Zone
	API         *api.MtgjsonApi
}
