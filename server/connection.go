package server

import (
	"context"
	"net"
)

/*
Connection - A representation of a client connecting to the server. Processes Key Exchange between
the client and the server and holds connection information for the client
*/
type Connection struct {
	// IPAddress - The remote IP Address of the client connection
	IPAddress string

	// Conn - The connection structure used for sending and receiving messages from the client
	conn *net.TCPConn

	// cryptoHandler - Provides logic for generating encryption keys
}

/*
NewConnection - A constructor for the Client structure. Creates new Client structure
and returns a pointer to it
*/
func NewConnection(conn *net.TCPConn) *Connection {
	return &Connection{
		IPAddress: conn.RemoteAddr().String(),
		conn:      conn,
	}
}

/*
Disconnect - Disconnects the client from the server
*/
func (connection *Connection) Disconnect() error {
	err := connection.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

/*
Initialize - Instructs the connection to perform key negotiation with the server
*/
func (connection *Connection) Initialize(ctx context.Context) error {
	return nil
}
