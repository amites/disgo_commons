package types

import (
"time"
"math/big"
)

// Transaction
type Transaction struct {
	Hash  [HashLength] byte
	Type  int
	From  Address
	To    Address
	Value *big.Int
	Time  time.Time
}
