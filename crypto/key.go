package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
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

/*
PublicKeyPEM - Convert the public key to a PEM encoded string for transmission
*/
func (key *KeyPair) PublicKeyPEM() string {
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
	}

	return string(pem.EncodeToMemory(block))
}

/*
PublicKeyChecksum - Generate a SHA-256 checksum for a RSA public key
*/
func (key *KeyPair) PublicKeyChecksum() string {
	hash := sha256.Sum256([]byte(key.PublicKeyPEM()))
	return hex.EncodeToString(hash[:])
}
