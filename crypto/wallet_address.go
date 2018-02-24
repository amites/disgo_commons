package crypto

import (
	"github.com/dispatchlabs/disgo_commons/constants"
	"encoding/hex"
	"github.com/haltingstate/secp256k1-go/dep"

	log "github.com/sirupsen/logrus"
)

// NewWalletAddress
func NewWalletAddress() ([constants.AddressLength]byte, error) {

	publicKey, privateKey := secp256k1.GenerateKeyPair()


	/*
	// TODO: How do we generate the private key?
	privateKey := make([]byte, constants.HashLength)
	rand.Read(privateKey)
	*/

	log.Info(hex.EncodeToString(privateKey))
	log.Info(hex.EncodeToString(publicKey))

	// Create address.
	/*
	publicKey := Sum256(privateKey)

	address := [constants.AddressLength]byte{}
	for i := 0; i < constants.AddressLength; i++ {
		address[i] = publicKey[i+12]
	}

	foo:= "3ebCa6556E28314925433E6D785cd429dbe363A3"

	foo1 := hex.EncodeToString(address[:])

	log.Info(foo1)

	if foo == foo1 {
			log.Info("YES")
	}
	*/

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
