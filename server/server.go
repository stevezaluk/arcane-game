package server

import (
	"context"
	"github.com/spf13/viper"
	"github.com/stevezaluk/arcane-game-server/options"
	"log/slog"
	"net"
	"strconv"
	"sync"
)

/*
IServer - The interface that the Server structure implements
*/
type IServer interface {
	// Sock - Returns a pointer to the net.Listener structure that the server uses
	Sock() *net.TCPListener

	// Log - Returns a pointer to the Logger structure that the server is using
	Log() *Log

	// SetOptions - Sets the connection options for the server
	SetOptions(*options.ConnectionOptions)

	// listen - Creates a new raw TCP socket and instructs the server to start listening on the port specified in server.Port
	listen() error

	// acceptConnection - Helper function for adding the successful connection to go-routine safe map stored in the Server struct
	acceptConnection(*Connection)

	// acceptConnections - Instructs the server to loop infinitely and await new client connections
	acceptConnections()

	// Start - Primary entrypoint for starting the server
	Start()
}

/*
Server - The primary construct used for handling user connections and providing logging
*/
type Server struct {
	// opts - The user-selected options used for new Connections
	opts *options.ConnectionOptions

	// log - Provides logic for creating and writing log files
	log *Log

	// sock - The socket used for establishing connections between the server and its clients
	sock *net.TCPListener

	// connections - A thread-safe map used for storing successfully connected clients.
	connections *sync.Map

	// cryptoHandler - Provides logic for handling crypto related operations like generating encryption keys

	// Port - The network port that the server is listening for connections on
	Port int

	// ConnectionCount - A 32-bit unsigned integer representing the number of securely connected clients
	ConnectionCount uint32

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
func (server *Server) Sock() *net.TCPListener {
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

	/*
		Were casting this here as due to the network flag that was specified
		in net.Listen guarantee's that we will get a net.TCPListener object
		returned to us.

		This ensures that we can store a pointer for this struct (to avoid copying),
		as net.Listen returns net.Listener which is an interface
	*/
	server.sock = sock.(*net.TCPListener)

	return nil
}

/*
acceptConnection - Performs post-connection actions. Primarily increments the Server.ConnectionCount value and
adds the client object to a go-routine safe map found at Server.clients. This differs from a standard map as this
uses an internal mutex to ensure go-routine safety.
*/
func (server *Server) acceptConnection(client *Connection) {
	slog.Info("Client has connected successfully", "addr", client)

	server.ConnectionCount++
	server.connections.Store(client.IPAddress, client)
}

/*
acceptConnections - Intended to be called after Server.listen. Instructs the server to accept new connections either until
the max connection count has been reached or until the servers acceptConnectionsTimer has expired. Which ever comes
first

if Server.sock is nil (indicating that Server.listen) was not called before this call, then the function returns
and does not accept new connections

A sync.WaitGroup is defined here and is incremented to ensure that all go-routines stay in sync. If the client's
clientConnectionTimer expires, then the client is forcibly removed from the server. This is to ensure that newly
connected clients do not hold the server from starting. This value is customizable using either the function:
Server.SetConnectionTimers or by calling FromConfig. If FromConfig is called then these values will be propagated
with the values defined in viper (or the defaults if they are not set; 120 seconds and 30 seconds respectfully)

To be clear here, these timers determine the time it takes for the to both complete Server/Client key exchange and
for the server to fetch its user data from the MTGJSON API.
*/
func (server *Server) acceptConnections() {
	/*
		We return here abruptly, as this indicates that Server.listen was not called before
		this function was.
	*/
	if server.Sock() == nil {
		return
	}

	var wg sync.WaitGroup

	for {
		if server.ConnectionCount == server.opts.MaxConnectionCount {
			slog.Debug("Servers max connection count has been reached. No new connections will be accepted")
			server.IsClosed = true
			break
		}

		if server.IsClosed {
			break
		}

		conn, err := server.Sock().Accept()
		if err != nil {
			slog.Error("Failed to accept connection for client", "err", err)
			continue
		}

		/*
			server.Sock().Accept() blocks until a new connection is received. When this is
			accepted, we place the client on a new go-routine to ensure that we can finish
			server/client key exchange and not hold up any new connections.

			This is being wrapped with a closure as I don't want the client struct having
			awareness of concurrency primitives (design choice). Only the Server struct should
			ever be aware of this.
		*/
		go func() {
			slog.Info("Attempting connection from client", "addr", conn.RemoteAddr())
			/*
				Were casting our connection again here, for the same reason as in Server.listen()
			*/
			connection := NewConnection(conn.(*net.TCPConn))

			/*
				This timeout context ensures that we can provide a deadline for the go-routine to
				finish and is how Server.opts.ClientTimeout is enforced
			*/

			ctx, cancel := context.WithTimeout(
				context.Background(),
				server.opts.ClientTimeout,
			)

			kexErr := connection.Initialize(ctx)
			if kexErr != nil {
				slog.Error("Client has failed to perform key negotiation. Client connection will be closed", "err", kexErr)

			}

			defer wg.Done()
			defer cancel()
		}()
	}

	/*
		Calling wg.Wait() here as we want all clients to finish initializing (completing server/client
		key-exchange) before we start game logic / pre-processing the player(s)
	*/
	wg.Wait()
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
	server.acceptConnections()
}
