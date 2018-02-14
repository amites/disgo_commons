package types

// Account
type Account struct {
	Id           string
	Address      Address
	Name         string
	Balance      int
	Transactions []Transaction
}
