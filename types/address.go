package types

import (
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

// Address
type Address [AddressLength] byte

// NewAddress
func NewAddress() (*Address, error) {

	// TODO: How do we generate the private key?
	privateKey, error := hex.DecodeString("b205a1e03ddf50247d8483435cd91f9c732bad281ad420061ab4310c33166276")
	if error != nil {
		log.WithFields(log.Fields{
			"method": "NewAddress",
		}).Info("unable to create new address")
		return nil, error
	}

	// Create address.
	address := Address{}
	publicKey := sha3.Sum256(privateKey)
	for i := 0; i < AddressLength; i++ {
		address[i] = publicKey[i+12]
	}

	return &address, nil
}
