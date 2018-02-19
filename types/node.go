package types

import "github.com/dispatchlabs/disgo_commons/constants"

// Node
type Node struct {
	Hash  [constants.HashLength] byte
	Ip    string
}
