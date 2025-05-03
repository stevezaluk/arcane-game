package crypto

import (
	"crypto/rand"
	"crypto/rsa"
)

/*
KeyPair - An abstraction of an RSA Key Pair that is used in Arcane's Key Exchange
between the client and server
*/
type KeyPair struct {
	// publicKey - The public key that matches the private key in privateKey
	publicKey rsa.PublicKey

	// privateKey - The private key generated that will be used for encryption
	privateKey *rsa.PrivateKey
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

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		privateKey: privateKey,
		publicKey:  privateKey.PublicKey,
	}, nil
}
