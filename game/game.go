package game

import (
	"github.com/stevezaluk/mtgjson-models/user"
	"github.com/stevezaluk/mtgjson-sdk-client/api"
)

const (
	/*
		Constants used for identifying the game mode selected for
		the game
	*/

	CommanderGameMode = "gamemode:commander"
	ModernGameMode    = "gamemode:modern"
	StandardGameMode  = "gamemode:standard"
)

/*
Game - A representation of a single MTG Game. Process's game commands sent from the client
*/
type Game struct {
	Name     string
	GameMode string

	Players   map[string]*Player
	GameOwner *Player

	Battlefield *Zone
	Exile       *Zone
	Command     *Zone
	API         *api.MtgjsonApi
}

/*
NewGame - Initialize the zones of a new Game and return a pointer to it
*/
func NewGame(name string, gameMode string) *Game {
	battlefield := NewZone(BattlefieldZoneId, nil, true, true, false)
	exile := NewZone(ExileZoneId, nil, true, true, false)

	var command *Zone
	if gameMode == CommanderGameMode {
		command = NewZone(CommanderZoneId, nil, true, true, false)
	}

	return &Game{
		Name:        name,
		GameMode:    gameMode,
		Battlefield: battlefield,
		Exile:       exile,
		Command:     command,
		API:         api.New(),
	}
}

/*
LookupPlayer - Fetch a Player object for an associating user. Returns nil
if the player could not be found
*/
func (game *Game) LookupPlayer(user *user.User) *Player {
	value, found := game.Players[user.Email]
	if !found {
		return nil
	}

	return value
}
