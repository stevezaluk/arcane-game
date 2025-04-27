package server

import "time"

/*
ConnectionOptions - A structure for tracking parameterized options to use relating to client and server connections
*/
type ConnectionOptions struct {
	// MaxConnectionCount - The maximum amount of successful connections the server will accept
	MaxConnectionCount uint32

	// ClientTimeout - The number of seconds the server should wait for a client to complete Key Exchange (assuming it is enabled)
	ClientTimeout *time.Duration

	// WaitConnectionsTimeout - The number of seconds the server should wait to accept new clients before closing new connections
	WaitConnectionsTimeout *time.Duration
}
