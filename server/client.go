package server

import (
	"github.com/stevezaluk/arcane-game/crypto"
	"net"
)

/*
Client - A representation of a client connecting to the server. Processes Key Exchange between
the client and the server and holds connection information for the client
*/
type Client struct {
	// IPAddress - The remote IP Address of the client connection
	IPAddress string

	// Conn - The connection structure used for sending and receiving messages from the client
	conn net.Conn

	// cryptoHandler - Provides logic for generating encryption keys
	cryptoHandler *crypto.EncryptionHandler
}

/*
NewClient - A constructor for the Client structure. Creates new Client structure
and returns a pointer to it
*/
func NewClient(conn net.Conn) *Client {
	return &Client{
		IPAddress: conn.RemoteAddr().String(),
		conn:      conn,
	}
}
