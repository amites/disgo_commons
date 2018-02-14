package types

// IConsensus
type IConsensus interface {

	CreateTransaction(transaction Transaction, transactions []Transaction) (*Transaction, error)
}

