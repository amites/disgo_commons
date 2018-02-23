package crypto

import (
"github.com/ebfe/keccak"
"crypto/rand"
"github.com/dispatchlabs/disgo_commons/constants"
"encoding/hex"
)

// Sum256
func Sum256(data []byte) (digest [constants.HashLength]byte) {
	hash := keccak.NewSHA3256()
	hash.Write(data)
	hash.Sum(digest[:0])
	return
}

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

