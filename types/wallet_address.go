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
