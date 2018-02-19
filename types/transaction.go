package types

import (
	"time"
	"encoding/json"
	"encoding/hex"

	"github.com/dispatchlabs/disgo_commons/crypto"
)

// Transaction
type Transaction struct {
	TxId  int64
	Hash  [HashLength]byte
	Type  int
	From  WalletAddress
	To    WalletAddress
	Value int64
	Time  time.Time
	CurrentValidators []WalletAddress
}

// NewTransactionFromJson
func NewTransactionFromJson(bytes []byte) (*Transaction, error) {

	var jsonMap map[string]interface{}
	error := json.Unmarshal(bytes, &jsonMap);
	if error != nil {
		return nil, error
	}

	transaction := Transaction{}
	transaction.Hash = crypto.CreateHash()
	from, error := hex.DecodeString(jsonMap["from"].(string))
	if error != nil {
		return nil, error
	}
	for i := 0; i < AddressLength; i++ {
		transaction.From[i] = from[i]
	}
	to, error := hex.DecodeString(jsonMap["to"].(string))
	if error != nil {
		return nil, error
	}
	for i := 0; i < AddressLength; i++ {
		transaction.To[i] = to[i]
	}
	//transaction.Value = jsonMap["value"].(int64)

	return &transaction, error
}

// ToJson
func (transaction *Transaction) ToJson() ([]byte, error) {
	return json.Marshal(transaction)
}
