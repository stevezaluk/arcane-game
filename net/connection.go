package net

import (
	"github.com/golang/protobuf/proto"
	"github.com/stevezaluk/arcane-game/models"
	"log/slog"
	"net"
)

const bufferSize = 32000

/*
basicRead - Read protobuf messages from the server or from a client. A reference to a pre-created
protobuf message should be passed into the protoBuf parameter
*/
func basicRead(conn net.Conn, protoBuf proto.Message) error {
	buffer := make([]byte, bufferSize)

	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(buffer[:n], protoBuf)
	if err != nil {
		return err
	}

	return nil

}

/*
basicWrite - Use a connection object to either write messages to the server or the client
*/
func basicWrite(conn net.Conn, protoBuf proto.Message) error {
	buffer, err := proto.Marshal(protoBuf)
	if err != nil {
		return err
	}

	_, err = conn.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}

/*
ReadArcaneMessage - Reads a models.ArcaneMessage protobuf from the client or the server. If either send a malformed
request then the function returns an error describing the Unmarshalling error. Acts as a wrapper around basicRead
*/
func ReadArcaneMessage(conn net.Conn, message *models.ArcaneMessage) error {
	return basicRead(conn, message)
}

/*
WriteArcaneMessage - Writes a models.ArcaneMessage protobuf to the client or the server. Acts as a wrapper
around basicWrite
*/
func WriteArcaneMessage(conn net.Conn, message *models.ArcaneMessage) error {
	return basicWrite(conn, message)
}

/*
ReadGameState - Read game state either from the server. If the message is malformed then the function
returns an error describing the Unmarshaling error. Acts as a wrapper around basicRead
*/
func ReadGameState(conn net.Conn, state *models.GameState) error {
	return basicRead(conn, state)
}

/*
WriteGameState - Write game state either from the server. Acts as a wrapper around basicRead
*/
func WriteGameState(conn net.Conn, state *models.GameState) error {
	return basicWrite(conn, state)
}

/*
CloseConnection - Close a client connection and log any errors that occur
*/
func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		slog.Error("Failed to close client connection", "conn", conn.RemoteAddr(), "err", err.Error())
	}

	slog.Info("Closed client connection", "conn", conn.RemoteAddr())
}
