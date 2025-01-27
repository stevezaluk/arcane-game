package net

import (
	"github.com/spf13/viper"
	"log/slog"
	stdNet "net"
)

type GameClient struct {
	Conn stdNet.Conn
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

	err := BasicWrite(conn, "Hello from Client")
	if err != nil {
		slog.Error("Failed to send welcome message to server", "err", err.Error())
		return
	}
}
