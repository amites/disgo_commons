package types

import (
	"crypto/rand"
	"github.com/dispatchlabs/disgo_commons/crypto"

	"github.com/dispatchlabs/disgo_commons/constants"
)

// Address
type WalletAddress [constants.AddressLength] byte

// NewWalletAddress
func NewWalletAddress() (*WalletAddress, error) {

	// TODO: How do we generate the private key?
	privateKey := make([]byte, constants.HashLength)
	rand.Read(privateKey)

	// Create address.
	hash := crypto.Sum256(privateKey)
	address := WalletAddress{}
	for i := 0; i < constants.AddressLength; i++ {
		address[i] = hash[i+12]
	}

	return &address, nil
}

func GetAddressFromBytes(bytes []byte) (*WalletAddress) {
	addr := WalletAddress{}
	for i, val := range bytes {
		addr[i] = val
	}
	return &addr
}

func GetBytesFromAddress(wa *WalletAddress) (*[]byte) {
	byteArray := make([]byte, constants.AddressLength)
	for i, val := range wa {
		byteArray[i] = val
	}
	return &byteArray
}