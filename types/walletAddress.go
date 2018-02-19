package types

import (
	"crypto/rand"
	"github.com/dispatchlabs/disgo_commons/crypto"

)

// Address
type WalletAddress [AddressLength] byte

func NewWalletAddress() (*WalletAddress, error) {

	// TODO: How do we generate the private key?
	privateKey := make([]byte, HashLength)
	rand.Read(privateKey)

	// Create address.
	hash := crypto.Sum256(privateKey)
	address := WalletAddress{}
	for i := 0; i < AddressLength; i++ {
		address[i] = hash[i+12]
	}

	return &address, nil
}