package crypto

import (

"crypto/rand"
"github.com/dispatchlabs/disgo_commons/constants"
"encoding/hex"
)

// NewHash
func NewHash() [constants.HashLength] byte {
	// TODO: Is this how we should do this?
	bytes := make([]byte, constants.HashLength)
	rand.Read(bytes)
	return Sum256(bytes)
}

// ToHash
func ToHash(bytes []byte) [constants.HashLength]byte {
	hash := [constants.HashLength]byte{}
	copy(hash[:], bytes)
	return hash
}

// ToHashString
func ToHashString(hash [constants.HashLength]byte) string {
	return hex.EncodeToString(hash[:])
}

