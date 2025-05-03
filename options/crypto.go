package options

import "github.com/spf13/viper"

const (
	// KEXRSA - A constant for defining the RSA Algorithm for Key Exchange
	KEXRSA string = "RSA"

	// KEYSIZE2048 - A constant for defining the kye size as 2048 bits
	KEYSIZE2048 uint32 = 2048

	// KEYSIZE4096 - A constant for defining the key size as 4096 bits
	KEYSIZE4096 uint32 = 4096
)

/*
CryptoOptions - Controls options related to Client/Server Key Exchange
*/
type CryptoOptions struct {
	// SecureConnections - If set to true, it forces clients to perform key exchange between the server and the client
	SecureConnections bool

	// ValidateKeys - Controls if the server and client should validate there keys with checksum's to ensure integrity
	ValidateKeys bool

	// EncryptionAlgorithm - Set the symmetric encryption algorithm that should be used for key exchange. Can only be RSA for now
	EncryptionAlgorithm string

	// KeySize - Defines the size in bits that RSA key's will be generated in. Cannot be lower than 2048
	KeySize uint32
}

/*
Crypto - Returns a pointer to a CryptoOptions struct, filled with the recommended options
*/
func Crypto() *CryptoOptions {
	return &CryptoOptions{
		ValidateKeys:        true,
		EncryptionAlgorithm: KEXRSA,
		KeySize:             KEYSIZE4096,
	}
}

/*
FromConfig - Fills the CryptoOptions struct with values pulled from Viper. Overwrites all pre-existing
values
*/
func (opts *CryptoOptions) FromConfig() *CryptoOptions {
	return &CryptoOptions{
		ValidateKeys:        viper.GetBool("kex.validate_keys"),
		EncryptionAlgorithm: KEXRSA, // this is always set to RSA, as others are not supported at the moment
		KeySize:             viper.GetUint32("kex.key_size"),
	}
}

/*
SetSecureConnections - Allows you to determine if clients should be forced to exchange keys with the clients.
Determines if end-to-end encryption is enforced on the client. If set to true, then the client will be forced
to generate a key and exchange it with the server before ConnectionOptions.ClientTimeout expires. If they fail
to do so, either through errors during the process or if ConnectionOptions.ClientTimeout expires, then there
connection is forcibly evicted from the server. Assuming the client fails key-exchange, they will still be
allowed to re-connect and re-attempt key-exchange
*/
func (opts *CryptoOptions) SetSecureConnections(secured bool) *CryptoOptions {
	opts.SecureConnections = secured

	return opts
}

/*
SetValidateKeys - Allows you to control if the server should validate the keys that it receives. Generally, during Key
Exchange when the server receives a client's key it will calculate the checksum of that key and expect a response from
the client with its calculated value of the key. If the values match, then the key is accepted and the client moves
to the next phase of Key Exchange. If this is set to true, then the server will blindly accept any key it is handed.

Generally, this should stay on as it ensures that the key has not been modified during transmission. However, if the
server.Server is running in an environment where it wants incredibly low latency and minimal chatter between client and
server, then it can be disabled. No sensitive information is exposed between server/client communications aside from the
player's email address
*/
func (opts *CryptoOptions) SetValidateKeys(validate bool) *CryptoOptions {
	opts.ValidateKeys = validate

	return opts
}

/*
SetEncryptionAlgorithm - Allows you to define the symmetric encryption algorithm the server should use when starting
Key Exchange. Currently, the 'algorithm' parameter is ignored and KEXRSA is set here as other algorithms are not
supported.
*/
func (opts *CryptoOptions) SetEncryptionAlgorithm(algorithm string) *CryptoOptions {
	opts.EncryptionAlgorithm = KEXRSA

	return opts
}

/*
SetKeySize - Allows you to define the size of the keys the server/client generates during Key Exchange. If this value
is set any lower than 2048, it is reset back to 2048. Use either KEYSIZE2048 or KEYSIZE4096 to set these values.
*/
func (opts *CryptoOptions) SetKeySize(size uint32) *CryptoOptions {
	opts.KeySize = size

	return opts
}
