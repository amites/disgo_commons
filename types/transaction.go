package types

import (
	"github.com/dispatchlabs/disgo_commons/constants"
	"github.com/dispatchlabs/disgo_commons/crypto"
	"encoding/json"
	"encoding/hex"
	"time"
	"bytes"
	"encoding/binary"
)

// Transaction
type Transaction struct {
	Hash       [constants.HashLength]byte
	Id         int64
	Type       int
	From       [constants.AddressLength]byte
	To         [constants.AddressLength]byte
	Value      int64
	Time       time.Time
	Signature  [constants.SignatureLength] byte
	Validators [][constants.AddressLength]byte
}

// NewTransaction
func NewTransaction(privateKey []byte, from [constants.AddressLength]byte, to [constants.AddressLength]byte, value int64) *Transaction {
	transaction := &Transaction{}
	transaction.To = to
	transaction.From = from
	transaction.Value = value
	transaction.Time = time.Now()
	hashable := &struct {
		Hash  [constants.HashLength]byte
		Id    int64
		Type  int
		From  [constants.AddressLength]byte
		To    [constants.AddressLength]byte
		Value int64
		Time  time.Time
	}{
		Id:    transaction.Id,
		Type:  transaction.Type,
		From:  transaction.From,
		To:    transaction.To,
		Value: transaction.Value,
		Time:  transaction.Time,
	}
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, hashable)
	transaction.Hash = crypto.NewHash(buffer.Bytes())
	copy(transaction.Signature[:], crypto.NewSignature(privateKey, transaction.Hash))
	return transaction
}

// FromString
func (this Transaction) FromString() string {
	return crypto.ToWalletAddressString(this.From)
}

// ToString
func (this Transaction) ToString() string {
	return crypto.ToWalletAddressString(this.To)
}

// String
func (this Transaction) String() string {
	bytes, error := json.Marshal(this)
	if error != nil {
		return error.Error()
	}
	return string(bytes)
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
	if jsonMap["id"] != nil {
		this.Id = int64(jsonMap["id"].(float64))
	}
	if jsonMap["type"] != nil {
		this.Type = int(jsonMap["type"].(float64))
	}
	if jsonMap["from"] != nil {
		from, error := hex.DecodeString(jsonMap["from"].(string))
		if error != nil {
			return error
		}
		copy(this.From[:], from)
	}
	if jsonMap["to"] != nil {
		to, error := hex.DecodeString(jsonMap["to"].(string))
		if error != nil {
			return error
		}
		copy(this.To[:], to)
	}
	if jsonMap["value"] != nil {
		this.Value = int64(jsonMap["value"].(float64))
	}
	if jsonMap["time"] != nil {
		// TODO: How do we do this?
		//this.Time = jsonMap["value"]
	}

	return nil
}

// MarshalJSON
func (this Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Hash  string
		Id    int64     `json:"id,omitempty"`
		Type  int       `json:"type,omitempty"`
		From  string    `json:"from,omitempty"`
		To    string    `json:"to,omitempty"`
		Value int64     `json:"value,omitempty"`
		Time  time.Time `json:"time,omitempty"`
	}{
		Hash:  hex.EncodeToString(this.From[:]),
		Id:    this.Id,
		Type:  this.Type,
		From:  hex.EncodeToString(this.From[:]),
		To:    hex.EncodeToString(this.To[:]),
		Value: this.Value,
		Time:  this.Time,
	})
}
