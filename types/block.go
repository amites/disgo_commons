package types

type Block struct {
	Hash          [HashLength]byte
	PreviousBlock *Block
	NextBlock     *Block
	Transaction   []Transaction
}
