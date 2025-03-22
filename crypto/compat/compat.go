package compat

import (
	"github.com/depinnetwork/por-consensus/crypto"
)

// Sha256 is a compatibility function that redirects to the appropriate implementation
func Sha256(bz []byte) []byte {
	return crypto.Sha256(bz)
}

// Other crypto functions as needed
func RipemdHash(bz []byte) []byte {
	return crypto.Ripemd160(bz)
}

func Sha512(bz []byte) []byte {
	if len(bz) == 0 {
		return nil
	}
	hasher := crypto.Sha512New()
	hasher.Write(bz)
	return hasher.Sum(nil)
}
