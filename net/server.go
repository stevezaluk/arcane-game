package net

import (
	"github.com/stevezaluk/arcane-game/game"
	"net"
)

/*
GameServer - A representation of the game server as a whole. Responsible for connecting clients,
processing game commands, and interacting with the game.Game object
*/
type GameServer struct {
	Listener *net.Listener

	ConnectionCount int
	IsClosed        bool

	Game *game.Game
	// crypto here
}

/*
NewServer - Initialize the game server and any crypto related functions
*/
func NewServer(lobbyName string, gameMode string) (*GameServer, error) {
	gameObject := game.NewGame(lobbyName, gameMode)

	return &GameServer{
		Game: gameObject,
	}, nil
}
