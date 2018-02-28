package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/dispatchlabs/disgo_commons/crypto/secp256k1"
	"github.com/dispatchlabs/disgo_commons/constants"
	"github.com/dispatchlabs/disgo_commons/math"
	"github.com/ebfe/keccak"
)

// Sum256
func Sum256(data []byte) (digest [constants.HashLength]byte) {
	hash := keccak.NewSHA3256()
	hash.Write(data)
	hash.Sum(digest[:0])
	return
}

// GenerateKeyPair
func GenerateKeyPair() (publicKey, privateKey []byte) {
	key, error := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if error != nil {
		panic(error)
	}
	return secp256k1.CompressPubkey(key.PublicKey.X, key.PublicKey.Y), math.PaddedBigBytes(key.D, 32)
}
