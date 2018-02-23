package types

import (
	"github.com/dispatchlabs/disgo_commons/constants"
)

// WalletAccount
type WalletAccount struct {
	Id      string
	Address [constants.AddressLength]byte
	Name    string
	Balance int64
}