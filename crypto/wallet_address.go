package crypto

import (
	"github.com/dispatchlabs/disgo_commons/constants"
	"encoding/hex"
	
	"math/rand"
)

// NewWalletAddress
func NewWalletAddress() ([constants.AddressLength]byte, error) {

	//publicKey, privateKey := secp256k1.GenerateKeyPair()
	// TODO: How do we generate the private key?
	privateKey := make([]byte, constants.HashLength)
	rand.Read(privateKey)

	// Create address.
	publicKey := Sum256(privateKey)
	address := [constants.AddressLength]byte{}
	for i := 0; i < constants.AddressLength; i++ {
		address[i] = publicKey[i+12]
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
