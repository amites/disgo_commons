package types

// WalletAccount
type WalletAccount struct {
	Id           string
	Address      WalletAddress
	Name         string
	Balance      int64
}
