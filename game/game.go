package game

import (
	"github.com/stevezaluk/mtgjson-sdk-client/api"
)

const (
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
UnTapStep - Un-tap all the permanents a player controls in the battlefield
*/
func (game *Game) UnTapStep(player *Player) {
	for _, card := range game.Battlefield.All() {
		card.UnTap()
	}
}
