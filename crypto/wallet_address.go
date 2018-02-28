package crypto

import (
	"github.com/dispatchlabs/disgo_commons/constants"
	"encoding/hex"
)

// NewWalletAddress
func NewWalletAddress() ([constants.AddressLength]byte, error) {

	publicKey, _ := GenerateKeyPair()

	// Create address.
	hash := Sum256(publicKey[1:])
	address := [constants.AddressLength]byte{}
	for i := 0; i < constants.AddressLength; i++ {
		address[i] = hash[i+12]
	}

	return address, nil
}

// ToWalletAddress
func ToWalletAddress(bytes []byte) [constants.AddressLength]byte {
	address := [constants.AddressLength]byte{}
	copy(address[:], bytes)
	return address
}

// ToWalletAddressString
func ToWalletAddressString(address [constants.AddressLength]byte) string {
	return hex.EncodeToString(address[:])
}
