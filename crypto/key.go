package crypto

import (
	"crypto/rand"
	"crypto/rsa"
)

/*
KeyPair - Represents a private/public 4096-bit RSA key pair
*/
type KeyPair struct {
	privateKey *rsa.PrivateKey
	PublicKey  rsa.PublicKey
}

/*
NewKeyPair - A constructor for the *KeyPair object. Generates and validates a RSA key and
stores it within the KeyPair structure
*/
func NewKeyPair() (*KeyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		privateKey: privateKey,
		PublicKey:  privateKey.PublicKey,
	}, nil
}
