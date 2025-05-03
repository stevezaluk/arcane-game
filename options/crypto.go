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

type CryptoOptions struct {
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
func (crypto *CryptoOptions) FromConfig() *CryptoOptions {
	return &CryptoOptions{
		ValidateKeys:        viper.GetBool("kex.validate_keys"),
		EncryptionAlgorithm: KEXRSA, // this is always set to RSA, as others are not supported at the moment
		KeySize:             viper.GetUint32("kex.key_size"),
	}
}
