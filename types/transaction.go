package types

import (
	"time"
	"encoding/json"
	"encoding/hex"

	"github.com/dispatchlabs/disgo_commons/crypto"
	"github.com/dispatchlabs/disgo_commons/constants"
)

// Transaction
type Transaction struct {
	Id                int64
	Hash              [constants.HashLength]byte
	Type              int
	From              [constants.AddressLength]byte
	To                [constants.AddressLength]byte
	Value             int64
	Time              time.Time
	CurrentValidators [][constants.AddressLength]byte
}

// NewTransaction
func NewTransaction() (*Transaction) {
	transaction := Transaction{}
	transaction.Hash = crypto.NewHash()
	transaction.Time = time.Now()
	return &transaction
}

// UnmarshalJSON
func (this *Transaction) UnmarshalJSON(bytes []byte) error {
	var jsonMap map[string]interface{}
	error := json.Unmarshal(bytes, &jsonMap)
	if error != nil {
		return error
	}

	if jsonMap["id"] != nil {
		this.Id = int64(jsonMap["id"].(float64))
	}
	if jsonMap["hash"] != nil {
		hash, error := hex.DecodeString(jsonMap["hash"].(string))
		if error != nil {
			return error
		}
		copy(this.Hash[:], hash)
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
		this.Value= int64(jsonMap["value"].(float64))
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
		Id    int64     `json:"id,omitempty"`
		Hash  string    `json:"hash,omitempty"`
		Type  int       `json:"type,omitempty"`
		From  string    `json:"from,omitempty"`
		To    string    `json:"to,omitempty"`
		Value int64     `json:"value,omitempty"`
		Time  time.Time `json:"time,omitempty"`
	}{
		Id:    this.Id,
		Hash:  hex.EncodeToString(this.Hash[:]),
		Type:  this.Type,
		From:  hex.EncodeToString(this.From[:]),
		To:    hex.EncodeToString(this.To[:]),
		Value: this.Value,
		Time:  this.Time,
	})
}

// HashString
func (this Transaction) HashString() string {
	return crypto.ToHashString(this.Hash)
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
