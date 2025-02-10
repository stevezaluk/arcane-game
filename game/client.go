package game

import (
	"github.com/spf13/viper"
	"github.com/stevezaluk/arcane-game/crypto"
	"github.com/stevezaluk/arcane-game/net"
	"log/slog"
	stdNet "net"
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

func ClientConnect() (*GameClient, error) {
	ip := viper.GetString("client.server_ip") + ":" + viper.GetString("client.server_port")
	conn, err := stdNet.Dial("tcp", ip)
	if err != nil {
		return nil, err
	}

	return &GameClient{
		Conn: conn,
	}, nil
}

func (client *GameClient) Welcome() {
	conn := client.Conn

	err := net.BasicWrite(conn, "Hello from Client")
	if err != nil {
		slog.Error("Failed to send welcome message to server", "err", err.Error())
		return
	}
}
