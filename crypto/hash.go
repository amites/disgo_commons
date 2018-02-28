package crypto

import (
	"github.com/dispatchlabs/disgo_commons/constants"
	"encoding/hex"
)

// NewHash
func NewHash(bytes []byte) [constants.HashLength]byte {
	return Sum256(bytes)
}

// ToHash
func ToHash(bytes []byte) [constants.HashLength]byte {
	hash := [constants.HashLength]byte{}
	copy(hash[:], bytes)
	return hash
}

func ToBytes(hashValue [constants.HashLength]byte) []byte {
	byteArray := make([]byte, constants.AddressLength)
	copy(byteArray, hashValue[:])
	return byteArray

}

// ToHashString
func ToHashString(hash [constants.HashLength]byte) string {
	return hex.EncodeToString(hash[:])
}
