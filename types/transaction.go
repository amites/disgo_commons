package types

import (
	"github.com/dispatchlabs/disgo_commons/constants"
	"github.com/dispatchlabs/disgo_commons/crypto"
	"encoding/json"
	"encoding/hex"
	"encoding/binary"
	"bytes"
)

// NewTransaction
func NewTransaction(privateKey []byte, payload *TransactionPayload) *Transaction {
	transaction := &Transaction{}
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, transaction)
	transaction.Hash = crypto.NewHash(buffer.Bytes())
	transaction.Payload = payload
	copy(transaction.Signature[:], crypto.Sign(transaction.Hash, privateKey))
	return transaction
}

// Transaction
type Transaction struct {
	Hash      [constants.HashLength]byte
	Payload   *TransactionPayload
	Signature [constants.SignatureLength] byte
}

// HashString
func (this Transaction) HashString() string {
	return crypto.ToHashString(this.Hash)
}

// UnmarshalJSON
func (this *Transaction) UnmarshalJSON(bytes []byte) error {
	var jsonMap map[string]interface{}
	error := json.Unmarshal(bytes, &jsonMap)
	if error != nil {
		return error
	}

	if jsonMap["hash"] != nil {
		hash, error := hex.DecodeString(jsonMap["hash"].(string))
		if error != nil {
			return error
		}
		copy(this.Hash[:], hash)
	}

	return nil
}

// MarshalJSON
func (this Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Hash    string              `json:"id,omitempty"`
		Payload *TransactionPayload `json:"type,omitempty"`
	}{
		Hash:    hex.EncodeToString(this.Hash[:]),
		Payload: this.Payload,
	})
}
