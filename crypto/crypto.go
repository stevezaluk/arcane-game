package crypto

import (
	"context"
	"errors"
	arcaneNet "github.com/stevezaluk/arcane-game/net"
	"log/slog"
	"net"
)

// ErrKeyMismatch - Gets returned when the server/client fail to validate a key pair
var ErrKeyMismatch = errors.New("key: There was a key mismatch between the server and the client (the negotiated checksum are not the same)")

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
sendKey - Wrapper around the BasicWrite function. Sends a PEM encoded copy of the public key
stored in the key pair to the connection passed in as an argument. This function should not
be called directly, as there are specific handler functions for Server and Client key exchanges
so it is not exported
*/
func (handler *EncryptionHandler) sendKey(keyPair *KeyPair, conn net.Conn) error {
	err := arcaneNet.BasicWrite(conn, keyPair.PublicKeyPEM())
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
	buffer, err := arcaneNet.BasicRead(conn)
	if err != nil {
		return nil, err
	}

	keyPair, err := FromPEMPublicKey(buffer)
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
	buffer, err := arcaneNet.BasicRead(conn)
	if err != nil {
		return err
	}

	if buffer != keyPair.PublicKeyChecksum() {
		return ErrKeyMismatch
	}

	return nil
}

/*
sendKeyValidation - Generate a checksum for the public key stored in the key pair that was passed as an argument
and send it to the connection
*/
func (handler *EncryptionHandler) sendKeyValidation(keyPair *KeyPair, conn net.Conn) error {
	err := arcaneNet.BasicWrite(conn, keyPair.PublicKeyChecksum())
	if err != nil {
		return err
	}

	return nil
}

/*
ServerKEX - Start the server key exchange routine between the client and the server. First the server
sends it PEM encoded public key to the client and then waits for a response from the client to validate
the key it has stored. If errors arise here they are logged, the connection is cancelled, and associating
go-routines are cancelled
*/
func (handler *EncryptionHandler) ServerKEX(ctx context.Context, conn net.Conn) {
	slog.Info("Starting key exchange between client", "conn", conn.RemoteAddr())

	err := handler.sendKey(handler.ServerKey(), conn)
	if err != nil {
		slog.Error("Failed to send key to client", "err", err)
		return
	}

	err = handler.receiveKeyValidation(handler.ServerKey(), conn)
	if err != nil {
		slog.Error("Key validation for server key pair has failed", "err", err)
		return
	}

	clientKeyPair, err := handler.receiveKey(conn)
	if err != nil {
		slog.Error("Failed to receive key from client", "err", err)
		return
	}

	err = handler.sendKeyValidation(clientKeyPair, conn)
	if err != nil {
		slog.Error("Key validation for client key pair has failed", "err", err)
	}

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

	serverKeyPair, err := handler.receiveKey(conn)
	if err != nil {
		slog.Error("Failed to receive key from server", "err", err.Error())
		return
	}

	err = handler.sendKeyValidation(serverKeyPair, conn)
	if err != nil {
		slog.Error("Key validation for server key pair has failed", "err", err)
		return
	}

	handler.serverKey = serverKeyPair

	err = handler.sendKey(handler.ClientKey(), conn)
	if err != nil {
		slog.Error("Failed to send client key pair to server", "err", err)
		return
	}

	err = handler.receiveKeyValidation(handler.ClientKey(), conn)
	if err != nil {
		slog.Error("Key validation for client key pair has failed", "err", err)
		return
	}
}
