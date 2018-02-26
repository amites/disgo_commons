package types

import (
	"github.com/dispatchlabs/disgo_commons/crypto"
	"github.com/dispatchlabs/disgo_commons/utils"
)
type Endpoint struct {
	Host	string
	Port 	int64
}

type Contact struct {
	Address		string
	Endpoint 	*Endpoint
}

func NewContact() *Contact {
	addr, err := crypto.NewWalletAddress()
	addrString := crypto.ToWalletAddressString(addr)
	if(err != nil) {
		panic(err)
	}
	return &Contact{
		Address:  addrString,
		Endpoint: &Endpoint{
			Port: 1975,
			Host: utils.GetLocalIP(),
		},
	}
}