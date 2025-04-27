package server

import (
	"github.com/spf13/viper"
	"github.com/stevezaluk/arcane-game/crypto"
	"net"
)

/*
IServer - The interface that the Server structure implements
*/
type IServer interface {
	// Sock - Returns a pointer to the net.Listener structure that the server uses
	Sock() *net.Listener

	// Log - Returns a pointer to the Logger structure that the server is using
	Log() *Log
}

/*
Server - The primary construct used for handling user connections and providing logging
*/
type Server struct {
	// sock - The socket used for establishing connections between the server and its clients
	sock *net.Listener

	// cryptoHandler - Provides logic for handling crypto related operations like generating encryption keys
	cryptoHandler *crypto.EncryptionHandler

	// log - Provides logic for creating and writing log files
	log *Log

	// Port - The network port that the server is listening for connections on
	Port int32

	// ConnectionCount - A 32-bit integer for tracking the number of connections to the server
	ConnectionCount int32

	// IsClosed - Determines if new connections to the server has been closed
	IsClosed bool
}

/*
New - Constructs the server and returns a pointer to it. Log is expected to be not nil,
and fully initialized
*/
func New(port int32, log *Log) *Server {
	return &Server{
		log:             log,
		Port:            port,
		ConnectionCount: 0,
		IsClosed:        false,
	}
}

/*
FromConfig - Constructs a server using config values provided by viper
*/
func FromConfig() *Server {
	return New(
		viper.GetInt32("server.port"),
		NewLoggerFromConfig(),
	)
}

/*
Sock - Returns a pointer to the net.Listener structure that the server uses
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
