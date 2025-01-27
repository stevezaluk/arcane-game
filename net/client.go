package net

import (
	"fmt"
	"github.com/spf13/viper"
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
	fmt.Println(err)
}
