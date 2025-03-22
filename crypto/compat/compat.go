package compat

import (
	cmtcrypto "github.com/depinnetwork/por-consensus/crypto"
)

// Sha256 is a compatibility function that redirects to the appropriate implementation
func Sha256(bz []byte) []byte {
	return cmtcrypto.Sha256(bz)
}

// Other crypto functions can be added here as needed
