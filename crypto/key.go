package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

/*
KeyPair - An abstraction of an RSA Key Pair that is used in Arcane's Key Exchange
between the client and server
*/
type KeyPair struct {
	// publicKey - The public key that matches the private key in privateKey
	publicKey rsa.PublicKey

	// privateKey - The private key generated that will be used for decryption
	privateKey *rsa.PrivateKey
}

/*
FromPublicKey - Takes an RSA public key and returns a Key Pair for it. Used when accepting the clients
public key
*/
func FromPublicKey(public rsa.PublicKey) *KeyPair {
	return &KeyPair{publicKey: public}
}

/*
FromPrivateKey - Takes a pointer to an RSA private key and constructs a KeyPair structure from it.
If key validation fails, then a nil pointer along with the error will be returned
*/
func FromPrivateKey(private *rsa.PrivateKey) (*KeyPair, error) {
	err := private.Validate()
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		privateKey: private,
		publicKey:  private.PublicKey,
	}, nil
}

/*
GenerateKey - Generates a Key Pair that with the key size that was defined. If size == 0,
then a 2048-bit key is used.
*/
func GenerateKey(size int) (*KeyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, err
	}

	return FromPrivateKey(privateKey)
}

/*
Encrypt - Encrypt a byte slice using the public key stored within the key pair. If no public key
is stored in the key pair, then nil is returned and an error outlining this problem is returned

Since server.Server sends data using protobuf's in wire format (encoded to bytes), the message is not
converted to a string or any other structure at any time during the encryption process. This ensures
that the server does not need to waste time arbitrarily converting between different types

Once encryption has completed, the cipher text is base64 encoded to get the size of the data as small
as possible.
*/
func (key *KeyPair) Encrypt(message []byte) ([]byte, error) {
	if key.publicKey.N.BitLen() == 0 {
		return nil, errors.New("no public key specified for use with encryption")
	}

	cipherBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &key.publicKey, message, nil)
	if err != nil {
		return nil, err
	}

	ret := make([]byte, base64.URLEncoding.EncodedLen(len(cipherBytes)))

	base64.URLEncoding.Encode(
		ret,
		cipherBytes,
	)

	return ret, nil
}

/*
Decrypt - Decypts a byte slice using the private key stored within the key pair. If no private
key is specified here, then nil is returned and an un-named error is returned outlining this
problem.

This function expects the 'message' parameter to be a base64 (URL) encoded byte array, usually received
from the client.
*/
func (key *KeyPair) Decrypt(message []byte) ([]byte, error) {
	if key.privateKey == nil {
		return nil, errors.New("no private key specified for use with decryption")
	}

	cipherBytes, err := base64.URLEncoding.DecodeString(string(message))
	if err != nil {
		return nil, err
	}

	data, err := key.privateKey.Decrypt(rand.Reader, cipherBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return nil, err
	}

	return data, err
}
