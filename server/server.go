package server

import (
	"github.com/spf13/viper"
	"github.com/stevezaluk/arcane-game/options"
	"log/slog"
	"net"
	"strconv"
)

/*
IServer - The interface that the Server structure implements
*/
type IServer interface {
	// Sock - Returns a pointer to the net.Listener structure that the server uses
	Sock() *net.Listener

	// Log - Returns a pointer to the Logger structure that the server is using
	Log() *Log

	// SetOptions - Sets the connection options for the server
	SetOptions(*options.ConnectionOptions)

	// listen - Creates a new raw TCP socket and instructs the server to start listening on the port specified in server.Port
	listen() error
}

/*
Server - The primary construct used for handling user connections and providing logging
*/
type Server struct {
	// opts - The user-selected options used for new Connections
	opts *options.ConnectionOptions

	// sock - The socket used for establishing connections between the server and its clients
	sock *net.Listener

	// cryptoHandler - Provides logic for handling crypto related operations like generating encryption keys

	// log - Provides logic for creating and writing log files
	log *Log

	// Port - The network port that the server is listening for connections on
	Port int

	// ConnectionCount - A 32-bit integer for tracking the number of connections to the server
	ConnectionCount int32

	// IsClosed - Determines if new connections to the server has been closed
	IsClosed bool
}

/*
New - Constructs the server and returns a pointer to it. Log is expected to be not nil,
and fully initialized with server.NewLogger or server.NewLoggerFromConfig
*/
func New(port int, log *Log, opts *options.ConnectionOptions) *Server {
	return &Server{
		opts:            opts,
		log:             log,
		Port:            port,
		ConnectionCount: 0,
		IsClosed:        false,
	}
}

/*
FromConfig - Constructs a server using config values provided by viper. Automatically creates a server.Log
structure using the log.path value provided by Viper
*/
func FromConfig() *Server {
	return New(
		viper.GetInt("server.port"),
		NewLoggerFromConfig(),
		options.Connection().FromConfig(),
	)
}

/*
Sock - Returns a pointer to the net.Listener structure that the server uses. This will return nil
if it is called before server.Listen has been called
*/
func (server *Server) Sock() *net.Listener {
	return server.sock
}

/*
Log - Returns a pointer to the Logger structure that the server is using
*/
func (server *Server) Log() *Log {
	return server.log
}

/*
SetOptions - Sets the connection options for the server
*/
func (server *Server) SetOptions(opts *options.ConnectionOptions) {
	server.opts = opts
}

/*
listen - Creates a new raw TCP socket and instructs the server to start listening on the port specified in
server.Port. New connections will not be accepted until a subsequent call to acceptConnections is made
*/
func (server *Server) listen() error {
	sock, err := net.Listen("tcp",
		"127.0.0.1:"+strconv.Itoa(server.Port),
	)

	if err != nil {
		return err
	}

	server.sock = &sock

	return nil
}

/*
Start - Primary entrypoint for starting the server. First the lobby is initialized and the server enables listening
on the TCP socket that is created with Server.listen. After listening is enabled (assuming this does not return an
error), the server starts waiting for connections, and processes them. If EnableACLs is set to true here
then the server will evaluate the incoming client IP Address to determine if they are allowed to connect to the server

After the client passes IP Address evaluation (assuming it is enabled) the server initiates the key exchange process
to establish secure, end-to-end encrypted communication. Server/Client keys are not persisted on disk and are unique
to the session. If EnableSecureConnections is enabled here then the no key exchange is performed for any users, and
they are immediately sent to the Lobby for pre-user processing (see the game.Lobby structure for a diagram of this)
*/
func (server *Server) Start() {
	slog.Info("Starting game server", "port", server.Port)
	err := server.listen()
	if err != nil {
		slog.Error("Failed to start listening for connections", "err", err.Error())
		return
	}

	slog.Info("Server now waiting for client connections", "maxConnections", server.opts.MaxConnectionCount)
}
