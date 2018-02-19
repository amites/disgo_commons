package crypto

import (
	"github.com/ebfe/keccak"
	"crypto/rand"
	"github.com/dispatchlabs/disgo_commons/constants"
)

// Sum256
func Sum256(data []byte) (digest [constants.HashLength]byte) {
	hash := keccak.NewSHA3256()
	hash.Write(data)
	hash.Sum(digest[:0])
	return
}

// CreateHash
func CreateHash() [constants.HashLength] byte {
	// TODO: Is this how we should do this?
	bytes := make([]byte, constants.HashLength)
	rand.Read(bytes)
	return Sum256(bytes)
}
