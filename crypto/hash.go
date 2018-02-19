package crypto

import (
	"github.com/ebfe/keccak"
	"crypto/rand"
)

// Sum256
func Sum256(data []byte) (digest [32]byte) {
	hash := keccak.NewSHA3256()
	hash.Write(data)
	hash.Sum(digest[:0])
	return
}

// CreateHash
func CreateHash() [32] byte {
	// TODO: Is this how we should do this?
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return Sum256(bytes)
}
