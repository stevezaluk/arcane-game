package server

import (
	"github.com/spf13/viper"
	"time"
)

/*
ConnectionOptions - A structure for tracking parameterized options to use relating to client and server connections
*/
type ConnectionOptions struct {
	// MaxConnectionCount - The maximum amount of successful connections the server will accept
	MaxConnectionCount uint32

	// ClientTimeout - The number of seconds the server should wait for a client to complete Key Exchange (assuming it is enabled)
	ClientTimeout time.Duration

	// WaitConnectionsTimeout - The number of seconds the server should wait to accept new clients before closing new connections
	WaitConnectionsTimeout time.Duration

	// EnforceACLs - If set to true, then enforce the lists defined in ConnectionOptions.Whitelist and ConnectionOptions.Blacklist
	EnforceACLs bool

	// Whitelist - Explicitly allow the IP Addresses defined in this slice
	Whitelist []string

	// Blacklist - Explicitly block the IP Addresses defined in this slice
	Blacklist []string
}

/*
Connection - Returns an empty ConnectionOptions struct
*/
func Connection() *ConnectionOptions {
	return &ConnectionOptions{}
}

/*
FromConfig - Fills the ConnectionOptions struct with values pulled from Viper
*/
func FromConfig() *ConnectionOptions {
	return &ConnectionOptions{
		MaxConnectionCount:     viper.GetUint32("server.options.max_connections"),
		ClientTimeout:          viper.GetDuration("server.options.client_timeout"),
		WaitConnectionsTimeout: viper.GetDuration("server.options.wait_connections_timeout"),
		EnforceACLs:            viper.GetBool("server.options.enforce_acls"),
		Whitelist:              viper.GetStringSlice("server.options.whitelist"),
		Blacklist:              viper.GetStringSlice("server.options.blacklist"),
	}
}
