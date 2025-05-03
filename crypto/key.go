package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
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
func FromPublicKey(public []byte) (*KeyPair, error) {
	publicKey, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return nil, err
	}

	return &KeyPair{publicKey: *publicKey.(*rsa.PublicKey)}, nil
}

/*
GenerateKey - Generates a Key Pair that with the key size that was defined. If size is less than
2048, then the size is changed to 2048
*/
func GenerateKey(size int) (*KeyPair, error) {
	if size < 2048 {
		size = 2048
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		privateKey: privateKey,
		publicKey:  privateKey.PublicKey,
	}, nil
}

/*
MarshalKey - Marshal's a public key to DER encoded bytes so that it can be transferred efficiently
*/
func (key *KeyPair) MarshalKey() ([]byte, error) {
	if key.publicKey.N.BitLen() == 0 {
		return nil, errors.New("no public key specified for use with PKIX marshalling")
	}

	content, err := x509.MarshalPKIXPublicKey(&key.publicKey)
	if err != nil {
		return nil, err
	}

	return content, nil
}

/*
Checksum - Returns a byte array representing the SHA-256 checksum of the public key. Most commonly used for validating
that keys have not been modified during transfer
*/
func (key *KeyPair) Checksum() ([32]byte, error) {
	if key.publicKey.N.BitLen() == 0 {
		return [32]byte{}, errors.New("no public key specified for use with SHA-256 checksum")
	}

	content, err := key.MarshalKey()
	if err != nil {
		return [32]byte{}, err
	}

	return sha256.Sum256(content), nil
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
Decrypt - Decrypts a byte slice using the private key stored within the key pair. If no private
key is specified here, then nil is returned and an un-named error is returned outlining this
problem.

This function expects the 'message' parameter to be a base64 (URL) encoded byte array, usually received
from the client.
*/
func (key *KeyPair) Decrypt(message []byte) ([]byte, error) {
	if key.privateKey == nil {
		return nil, errors.New("no private key specified for use with decryption")
	}

	cipherBytes := make([]byte, base64.URLEncoding.DecodedLen(len(message)))

	n, err := base64.URLEncoding.Decode(cipherBytes, message)
	if err != nil {
		return nil, err
	}

	data, err := key.privateKey.Decrypt(rand.Reader, cipherBytes[:n], &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return nil, err
	}

	return data, err
}
