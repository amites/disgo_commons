package types

// Account
type WalletAccount struct {
	Id           string
	Address      WalletAddress
	Name         string
	Balance      int64
}
