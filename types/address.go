package types

import (
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"github.com/dispatchlabs/disgo_commons/crypto"
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
	hash := crypto.Sum256(privateKey)
	address := Address{}
	for i := 0; i < HashLength; i++ {
		address[i] = hash[i]
	}

	return &address, nil
}
