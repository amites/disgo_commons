package types

import (
	log "github.com/sirupsen/logrus"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/google/uuid"
)

// Account
type Account struct {
	Id           string
	Address      Address
	Name         string
	Balance      int
	Transactions []Transaction
}

// NewAccount
func NewAccount() (*Account, error) {

	// TODO: How do we generate private key?
	privateKey, error := hex.DecodeString("b205a1e03ddf50247d8483435cd91f9c732bad281ad420061ab4310c33166276")
	if error != nil {
		log.WithFields(log.Fields{
			"method": "NewAccount",
		}).Info("unable to create new account")
		return nil, error
	}

	// Create account.
	account := Account{}
	account.Id = uuid.New().String()
	publicKey := sha3.Sum256(privateKey)
	for i := 0; i < AddressLength; i++ {
		account.Address[i] = publicKey[i+12]
	}

	return &account, nil
}
