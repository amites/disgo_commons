package crypto

import (
	"github.com/dispatchlabs/disgo_commons/constants"
	"math/rand"
	"encoding/hex"
)

// NewWalletAddress
func NewWalletAddress() ([constants.AddressLength]byte, error) {

	// TODO: How do we generate the private key?
	privateKey := make([]byte, constants.HashLength)
	rand.Read(privateKey)

	// Create address.
	hash := Sum256(privateKey)
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

func AddressStringToBytes(address string) ([]byte, error) {
	result, err := hex.DecodeString(address)
	return result, err
}

func ToWalletAddressBytes(address [constants.AddressLength]byte) []byte {
	byteArray := make([]byte, constants.AddressLength)
	copy(byteArray, address[:])
	return byteArray
}