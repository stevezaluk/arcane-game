package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
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
NewKeyPair - A constructor for the *KeyPair object. Generates and validates an RSA key and
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

/*
EncryptMessage - Encrypt and base64 encode a plain text message using the public key and return its cipher
text
*/
func (key *KeyPair) EncryptMessage(message string) (string, error) {
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &key.PublicKey, []byte(message), nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.WithPadding(base64.StdPadding).EncodeToString(cipherText), nil
}

/*
DecryptMessage - Decrypt a base64 encoded encrypted message using the key pairs private key
*/
func (key *KeyPair) DecryptMessage(cipher string) (string, error) {
	cipherText, err := base64.StdEncoding.WithPadding(base64.StdPadding).DecodeString(cipher)
	if err != nil {
		return "", err
	}

	plainText, err := key.privateKey.Decrypt(nil, cipherText, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
