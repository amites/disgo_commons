package crypto

import (
	"github.com/ebfe/keccak"
)

// Sum256
func Sum256(data []byte) (digest [32]byte) {
	hash := keccak.NewSHA3256()
	hash.Write(data)
	hash.Sum(digest[:0])
	return
}
