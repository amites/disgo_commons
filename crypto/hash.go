package crypto

import (
	"github.com/ebfe/keccak"
	"github.com/dispatchlabs/disgo_commons/types"
)

// Sum256
func Sum256(data []byte) (digest [types.HashLength]byte) {
	hash := keccak.NewSHA3256()
	hash.Write(data)
	hash.Sum(digest[:0])
	return
}
