package crypto

import (
	"context"
	"errors"
	"fmt"
	"github.com/stevezaluk/arcane-game/models"
	arcaneNet "github.com/stevezaluk/arcane-game/net"
	"log/slog"
	"net"
	"time"
)

// ErrKeyMismatch - Gets returned when the server/client fail to validate a key pair
var ErrKeyMismatch = errors.New("key: There was a key mismatch between the server and the client (the negotiated checksum are not the same)")

// ErrInvalidCryptoResponse - Gets returned when either the client or the server sends improperly formatted key during exchange
var ErrInvalidCryptoResponse = errors.New("key: The response the client/server sent is malformed and cannot be parsed")

/*
EncryptionHandler - Contains logic for exchanging keys between the server and client, and
holds logic for sending encrypted messages
*/
type EncryptionHandler struct {
	serverKey *KeyPair
	clientKey *KeyPair
}

/*
NewClientHandler - Creates a new EncryptionHandler for use with the client. A new KeyPair
will get generated and stored while leaving the server key nil. This is to ensure that we can
set this value during EncryptionHandler.ClientKEX
*/
func NewClientHandler() (*EncryptionHandler, error) {
	clientKey, err := NewKeyPair()
	if err != nil {
		return nil, err
	}

	return &EncryptionHandler{serverKey: nil, clientKey: clientKey}, nil
}

/*
NewServerHandler - Creates a new EncryptionHandler for use on the game server. A new KeyPair
will get generated and stored while leaving the clientKey nil. This is to ensure that we can
set this value during EncryptionHandler.ServerKEX
*/
func NewServerHandler() (*EncryptionHandler, error) {
	serverKey, err := NewKeyPair()
	if err != nil {
		return nil, err
	}

	return &EncryptionHandler{serverKey: serverKey, clientKey: nil}, nil
}

/*
ClientKey - Return a pointer to the client key pair
*/
func (handler *EncryptionHandler) ClientKey() *KeyPair {
	return handler.clientKey
}

/*
ServerKey - Return a pointer to the server key pair
*/
func (handler *EncryptionHandler) ServerKey() *KeyPair {
	return handler.serverKey
}

/*
sendKey - Wrapper around the net.WriteArcaneMessage function. Sends a PEM encoded copy of the public key
stored in the key pair to the connection passed in as an argument. This function should not
be called directly, as there are specific handler functions for Server and Client key exchanges
so it is not exported
*/
func (handler *EncryptionHandler) sendKey(keyPair *KeyPair, conn net.Conn, isClient bool) error {
	identifier := "SERVER"
	if isClient {
		identifier = "CLIENT"
	}

	/*
		Represents a message being sent to the client/server containing our key
	*/
	message := &models.ArcaneMessage{
		Namespace:  models.ArcaneNamespace_CRYPTO_NAMESPACE,
		Action:     "ACCEPT",
		Identifier: identifier,
		Values:     []string{keyPair.PublicKeyPEM()},
	}

	err := arcaneNet.WriteArcaneMessage(conn, message)
	if err != nil {
		return err
	}

	return nil
}

/*
receiveKey - Wrapper around the net.BasicRead function. Waits for a PEM encoded key sent to the server.
Its associating key gets returned so that it can be stored with the players object and with the server. This
function should not be called directly, as there are specific handler functions for Server adn Client key
exchanges so it is not exported.
*/
func (handler *EncryptionHandler) receiveKey(conn net.Conn) (*KeyPair, error) {
	message := &models.ArcaneMessage{}

	err := arcaneNet.ReadArcaneMessage(conn, message)
	if err != nil {
		return nil, err
	}

	if len(message.Values) == 0 || len(message.Values) < 1 {
		return nil, ErrInvalidCryptoResponse
	}

	keyPair, err := FromPEMPublicKey(message.Values[0])
	if err != nil {
		return nil, err
	}

	return keyPair, nil
}

/*
receiveKeyValidation - Generates a checksum from the public key stored in the keyPair passed as an argument. If the
key pairs do not match then it returns an ErrKeyMismatch and its calling function (ServerKEX or ClientKEX),
cancels the context and kills the go routine
*/
func (handler *EncryptionHandler) receiveKeyValidation(keyPair *KeyPair, conn net.Conn) error {
	message := &models.ArcaneMessage{}

	err := arcaneNet.ReadArcaneMessage(conn, message)
	if err != nil {
		return err
	}

	if len(message.Values) == 0 || len(message.Values) < 1 {
		return ErrInvalidCryptoResponse
	}

	if message.Values[0] != keyPair.PublicKeyChecksum() {
		return ErrKeyMismatch
	}

	return nil
}

/*
sendKeyValidation - Generate a checksum for the public key stored in the key pair that was passed as an argument
and send it to the connection
*/
func (handler *EncryptionHandler) sendKeyValidation(keyPair *KeyPair, conn net.Conn, isClient bool) error {
	identifier := "SERVER"
	if isClient {
		identifier = "CLIENT"
	}

	message := &models.ArcaneMessage{
		Namespace:  models.ArcaneNamespace_CRYPTO_NAMESPACE,
		Action:     "VALIDATE",
		Identifier: identifier,
		Values:     []string{keyPair.PublicKeyChecksum()},
	}

	err := arcaneNet.WriteArcaneMessage(conn, message)
	if err != nil {
		return err
	}

	time.Sleep(2)

	return nil
}

/*
ServerKEX - Start the server key exchange routine between the client and the server. First the server
sends it PEM encoded public key to the client and then waits for a response from the client to validate
the key it has stored. If errors arise here they are logged, the connection is cancelled, and associating
go-routines are cancelled
*/
func (handler *EncryptionHandler) ServerKEX(ctx context.Context, conn net.Conn) {
	initMessage := &models.ArcaneMessage{}

	slog.Info("Waiting for crypto initialization message from client", "conn", conn.RemoteAddr())
	err := arcaneNet.ReadArcaneMessage(conn, initMessage)
	if err != nil {
		slog.Error("Failed to parse crypto initialization message. Either structure was incorrect or the read failed", "conn", conn.RemoteAddr(), "err", err.Error())
		arcaneNet.CloseConnection(conn)
		return
	}
	fmt.Println(initMessage)

	if initMessage.Namespace != models.ArcaneNamespace_CRYPTO_NAMESPACE {
		slog.Error("Crypto initialization message contained the wrong namespace. Connection was closed", "conn", conn.RemoteAddr())
		arcaneNet.CloseConnection(conn)
		return
	}

	if initMessage.Action != "INIT" {
		slog.Error("Crypto initialization message contained the wrong action. Connection was closed", "conn", conn.RemoteAddr())
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Info("Starting key exchange between client", "conn", conn.RemoteAddr(), "checkSum", handler.ServerKey().PublicKeyChecksum())

	slog.Debug("Sending server side key pair", "conn", conn.RemoteAddr(), "checkSum", handler.ServerKey().PublicKeyChecksum())
	err = handler.sendKey(handler.ServerKey(), conn, false)
	if err != nil {
		slog.Error("Failed to send key to client", "err", err)
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Debug("Waiting for server side key pair validation", "conn", conn.RemoteAddr())
	err = handler.receiveKeyValidation(handler.ServerKey(), conn)
	if err != nil {
		slog.Error("Key validation for server key pair has failed", "err", err)
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Debug("Server side key validated. Waiting for client side key pair", "conn", conn.RemoteAddr())
	clientKeyPair, err := handler.receiveKey(conn)
	if err != nil {
		slog.Error("Failed to receive key from client", "err", err)
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Debug("Received client key pair", "conn", conn.RemoteAddr(), "checkSum", clientKeyPair.PublicKeyChecksum())
	err = handler.sendKeyValidation(clientKeyPair, conn, true)
	if err != nil {
		slog.Error("Key validation for client key pair has failed", "err", err)
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Debug("Successfully validated client key pair", "conn", conn.RemoteAddr(), "checkSum", clientKeyPair.PublicKeyChecksum())
	slog.Info("Server side key exchange completed", "conn", conn.RemoteAddr())
	handler.clientKey = clientKeyPair
}

/*
ClientKEX - Starts the client side key exchange. First the routine waits to receive and validate
the server KeyPair then sends its own generated KeyPair to the server for validation. Once both are
validated the server proceeds with the rest of the client connection routine. If errors arise here
they are logged, the connection is cancelled, and associating go-routines are cancelled
*/
func (handler *EncryptionHandler) ClientKEX(ctx context.Context, conn net.Conn) {
	slog.Info("Starting key exchange between server", "conn", conn.RemoteAddr())

	initMessage := &models.ArcaneMessage{
		Namespace: models.ArcaneNamespace_CRYPTO_NAMESPACE,
		Action:    "INIT",
	}

	err := arcaneNet.WriteArcaneMessage(conn, initMessage)
	if err != nil {
		slog.Error("Failed to send crypto initialization message to the server", "conn", conn.RemoteAddr(), "err", err.Error())
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Debug("Waiting for server side key pair", "conn", conn.RemoteAddr())
	serverKeyPair, err := handler.receiveKey(conn)
	if err != nil {
		slog.Error("Failed to receive key from server", "err", err.Error())
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Debug("Received server key pair", "conn", conn.RemoteAddr(), "checkSum", serverKeyPair.PublicKeyChecksum())
	err = handler.sendKeyValidation(serverKeyPair, conn, false)
	if err != nil {
		slog.Error("Key validation for server key pair has failed", "err", err)
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Debug("Successfully validated server side key pair. Sending client side key pair", "conn", conn.RemoteAddr())
	handler.serverKey = serverKeyPair

	err = handler.sendKey(handler.ClientKey(), conn, true)
	if err != nil {
		slog.Error("Failed to send client key pair to server", "err", err)
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Debug("Waiting to receive validation for client side key pair", "conn", conn.RemoteAddr())
	err = handler.receiveKeyValidation(handler.ClientKey(), conn)
	if err != nil {
		slog.Error("Key validation for client key pair has failed", "err", err)
		arcaneNet.CloseConnection(conn)
		return
	}

	slog.Debug("Successfully validated client key pair", "conn", conn.RemoteAddr(), "checkSum", handler.clientKey.PublicKeyChecksum())
	slog.Info("Client side key exchange completed", "conn", conn.RemoteAddr())
}
