package options

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
