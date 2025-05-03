package options

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
	return &ConnectionOptions{
		MaxConnectionCount:     4,
		ClientTimeout:          30,
		WaitConnectionsTimeout: 120,
		EnforceACLs:            false,
		Whitelist:              []string{},
		Blacklist:              []string{},
	}
}

/*
FromConfig - Fills the ConnectionOptions struct with values pulled from Viper. Overwrites an pre-existing
values
*/
func (opts *ConnectionOptions) FromConfig() *ConnectionOptions {
	return &ConnectionOptions{
		MaxConnectionCount:     viper.GetUint32("server.max_connections"),
		ClientTimeout:          viper.GetDuration("server.client_timeout"),
		WaitConnectionsTimeout: viper.GetDuration("server.wait_connections_timeout"),
		EnforceACLs:            viper.GetBool("server.enforce_acls"),
		Whitelist:              viper.GetStringSlice("server.whitelist"),
		Blacklist:              viper.GetStringSlice("server.blacklist"),
	}
}

/*
SetMaxConnectionCount - Allows you to define the maximum amount of successful connections the server
should accept. If this is not set, then the default value of 4 is used. This includes the "Game Leader",
colloquially called the "Host". A successful connection is defined as a client that completes key exchange
before the ClientTimeout is fired.

An upper limit for this is not set, so there truly is not limit to what you can set this number to (aside from the
32-bit unsigned integer limit). I recommend setting this to 6 connections (or under) as feasibly there is no way to
play a game of Magic: The Gathering (in a reasonable amount of time) with more than 6 people.
*/
func (opts *ConnectionOptions) SetMaxConnectionCount(max uint32) *ConnectionOptions {
	opts.MaxConnectionCount = max

	return opts
}

/*
SetClientTimeout - Allows you to define the amount of time the client has to complete key exchange before
there connection gets forcibly closed and removed from the server. Reconnections even after this has expired
are still allowed.

Timeout must be defined in seconds
*/
func (opts *ConnectionOptions) SetClientTimeout(timeout uint32) *ConnectionOptions {
	opts.ClientTimeout = time.Second * time.Duration(timeout)

	return opts
}

/*
SetWaitConnectionsTimeout - Allows you to define the amount of time the server should wait to aggregate
new connections before preventing new connections and starting game logic.

Currently, there is no way for the Game Leader (an arbitrarily defined Player in control of the server) to force the
Server to start, so the Server strictly abides by these timers and will not start until they expire. Be sure to properly
configure these if the defaults are not suitable.

If the servers successful connection count is zero after this timer fires, then the server is automatically closed
and stopped as there is no need to process game logic with no active connections

Timeout must be defined in seconds
*/
func (opts *ConnectionOptions) SetWaitConnectionsTimeout(timeout uint32) *ConnectionOptions {
	opts.WaitConnectionsTimeout = time.Second * time.Duration(timeout)

	return opts
}

/*
SetEnforceACLs - Instructs the server to validate client IP Addresses present in ConnectionOptions.Whitelist
and ConnectionOptions.Blacklist when accepting connections. Any client IP Addresses placed in these lists
are ignored unless ConnectionOptions.EnforceACLs is set to true
*/
func (opts *ConnectionOptions) SetEnforceACLs() *ConnectionOptions {
	opts.EnforceACLs = true

	return opts
}

/*
SetWhitelist - Set the server whitelist to the value defined in 'acl'. When a new client
attempts to connect to the server, its reported IP Address is checked to ensure that it
exists in this list. If it does not, then its connection is closed.

This validation happens before key exchange, to ensure that sensitive information (like
session keys) only get sent to trusted clients.

If a client IP Address is improperly formatted (ex: 11.11.11.111111111) then it is ignored
*/
func (opts *ConnectionOptions) SetWhitelist(acl []string) *ConnectionOptions {
	opts.Whitelist = acl

	return opts
}

/*
SetBlacklist - Set the server blacklist to the value defined in 'acl'. When a new client
attempts to connect to the server, its reported IP Address is checked to ensure that it
does not exist in this list. If it does, then its connection is closed.

This validation happens before key exchange, to ensure that sensitive information (like
session keys) only get sent to trusted clients

If a client IP Address is improperly formatted (ex: 11.11.11.111111111) then it is ignored
*/
func (opts *ConnectionOptions) SetBlacklist(acl []string) *ConnectionOptions {
	opts.Blacklist = acl

	return opts
}
