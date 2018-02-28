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

// ToPublicKey
func ToPublicKey(hash [constants.HashLength]byte, signature []byte) []byte {
	publicKey, error := secp256k1.RecoverPubkey(hash[:], signature)
	if error != nil {
		panic(error)
	}
	return publicKey
}

// Sign
func Sign(hash [constants.HashLength]byte, privateKey []byte) []byte {
	signature, error := secp256k1.Sign(hash[:], privateKey[:])
	if error != nil {
		panic(error)
	}
	return signature
}

// VerifySignature
func VerifySignature(hash [constants.HashLength]byte, signature []byte) bool {
	return secp256k1.VerifySignature(ToPublicKey(hash, signature), hash[:], signature)
}

/*
func (fs FrontierSigner) SignatureValues(tx *Transaction, sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != 65 {
		panic(fmt.Sprintf("wrong size for signature: got %d, want 65", len(sig)))
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v, nil
}
 */
