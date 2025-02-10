package crypto

/*
EncryptionHandler - Contains logic for exchanging keys between the server and client, and
holds logic for sending encrypted messages
*/
type EncryptionHandler struct {
	serverKey *KeyPair
	clientKey *KeyPair
}

/*
HandlerFromServerKey - Creates a new encryption handler from an existing server key. This generates
a fresh key pair that can be used with the client (specifically within the context of a player)
*/
func HandlerFromServerKey(serverKey *KeyPair) (*EncryptionHandler, error) {
	clientKey, err := NewKeyPair()
	if err != nil {
		return nil, err
	}

	return &EncryptionHandler{serverKey: serverKey, clientKey: clientKey}, nil
}
