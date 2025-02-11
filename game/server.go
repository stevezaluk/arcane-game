package game

import (
	"context"
	"github.com/spf13/viper"
	"github.com/stevezaluk/arcane-game/crypto"
	"log/slog"
	stdNet "net"
	"sync"
)

/*
GameServer - A representation of the game server as a whole. Responsible for connecting clients,
processing game commands, and interacting with the game.Game object
*/
type GameServer struct {
	Listener *stdNet.Listener

	ConnectionCount int
	IsClosed        bool

	Game          *Game
	CryptoHandler *crypto.EncryptionHandler
}

/*
NewServer - Initialize the game server and any crypto related functions
*/
func NewServer(lobbyName string, gameMode string) (*GameServer, error) {
	handler, err := crypto.NewServerHandler()
	if err != nil {
		return nil, err
	}

	gameObject := NewGame(lobbyName, gameMode)

	return &GameServer{
		Game:          gameObject,
		CryptoHandler: handler,
	}, nil
}

/*
listen - Instruct the server to start listening for connections
*/
func (server *GameServer) listen() error {
	listener, err := stdNet.Listen("tcp", "127.0.0.1:"+viper.GetString("server.port"))
	if err != nil {
		return err
	}

	server.Listener = &listener

	return nil
}

/*
handleClient - Provides an entrypoint for the client to start key negotiation with the server
*/
func (server *GameServer) handleClient(wg *sync.WaitGroup, conn stdNet.Conn) {
	server.CryptoHandler.ServerKEX(context.Background(), conn)

	defer wg.Done()
}

/*
waitForConnections - Instructs the server to wait for connections and accept them until the server reaches it's max
connection count
*/
func (server *GameServer) waitForConnections() {
	var wg sync.WaitGroup

	sock := *server.Listener
	for {
		if server.ConnectionCount == viper.GetInt("server.max_connections") {
			server.IsClosed = true
			break
		}

		if !server.IsClosed {
			conn, err := sock.Accept()
			if err != nil {
				continue
			}

			slog.Info("Client connected", "remoteAddr", conn.RemoteAddr())
			server.ConnectionCount++
			wg.Add(1)

			go server.handleClient(&wg, conn)
		}
	}

	wg.Wait()
}

/*
Start - Primary entrypoint for starting the server
*/
func (server *GameServer) Start() {
	slog.Info("Starting game server", "lobbyName", server.Game.Name, "gameMode", server.Game.GameMode)
	err := server.listen()
	if err != nil {
		slog.Error("Failed to start listening for connections", "err", err.Error())
		return
	}

	slog.Info("Server now waiting for new connections")
	server.waitForConnections()
}
