package types

import (
	"time"
)

// Transaction
type Transaction struct {
	Hash  [HashLength] byte
	Type  int
	From  Address
	To    Address
	Value int64
	Time  time.Time
}
