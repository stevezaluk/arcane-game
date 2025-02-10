package game

import (
	"github.com/stevezaluk/arcane-game/crypto"
	"github.com/stevezaluk/arcane-game/net"
	"log/slog"
	stdNet "net"
	"strconv"
)

type GameClient struct {
	Conn          stdNet.Conn
	CryptoHandler *crypto.EncryptionHandler
}

/*
NewClient - Constructor for the GameClient. Initializes the crypto.EncryptionHandler for the
client and then returns a pointer to a new GameClient
*/
func NewClient() (*GameClient, error) {
	handler, err := crypto.NewClientHandler()
	if err != nil {
		return nil, err
	}

	return &GameClient{CryptoHandler: handler}, nil
}

/*
Connect - Connect to a running Arcane server and initiate
*/
func (client *GameClient) Connect(ipAddress string, port int) error {
	uri := ipAddress + ":" + strconv.Itoa(port)

	conn, err := stdNet.Dial("tcp", uri)
	if err != nil {
		return err
	}

	client.Conn = conn

	return nil
}

func (client *GameClient) Welcome() {
	conn := client.Conn

	err := net.BasicWrite(conn, "Hello from Client")
	if err != nil {
		slog.Error("Failed to send welcome message to server", "err", err.Error())
		return
	}
}
