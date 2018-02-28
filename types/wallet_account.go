package types

import (
	"github.com/dispatchlabs/disgo_commons/constants"
	"encoding/hex"
	"encoding/json"
	"github.com/dispatchlabs/disgo_commons/crypto"
)

// WalletAccount
type WalletAccount struct {
	Id      string
	Address [constants.AddressLength]byte
	Name    string
	Balance int64
}

// NewWalletAccount
func NewWalletAccount() *WalletAccount {
	walletAccount := &WalletAccount{}
	walletAccount.Id = "ID"
	walletAccount.Address, _ = crypto.NewWalletAddress()
	walletAccount.Balance = 0
	return walletAccount
}

// ToAddressString
func (this WalletAccount) ToAddressString() string {
	return crypto.ToWalletAddressString(this.Address)
}

// String
func (this WalletAccount) String() string {
	bytes, error := json.Marshal(this)
	if error != nil {
		return error.Error()
	}
	return string(bytes)
}


// UnmarshalJSON
func (this *WalletAccount) UnmarshalJSON(bytes []byte) error {
	var jsonMap map[string]interface{}
	error := json.Unmarshal(bytes, &jsonMap)
	if error != nil {
		return error
	}

	if jsonMap["id"] != nil {
		this.Id = jsonMap["id"].(string)
	}
	if jsonMap["address"] != nil {
		address, error := hex.DecodeString(jsonMap["address"].(string))
		if error != nil {
			return error
		}
		copy(this.Address[:], address)
	}
	if jsonMap["name"] != nil {
		this.Name = jsonMap["name"].(string)
	}
	if jsonMap["balance"] != nil {
		this.Balance = int64(jsonMap["balance"].(float64))
	}

	return nil
}

// MarshalJSON
func (this WalletAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id      string `json:"id"`
		Address string `json:"address"`
		Name    string `json:"name,omitempty"`
		Balance int64  `json:"balance"`
	}{
		Id:      this.Id,
		Address: hex.EncodeToString(this.Address[:]),
		Name:    this.Name,
		Balance: this.Balance,
	})
}
