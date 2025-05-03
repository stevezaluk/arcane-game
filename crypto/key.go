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
