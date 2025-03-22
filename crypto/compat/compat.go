package compat

import (
	crypto "github.com/depinnetwork/por-consensus/crypto"
	"hash"
)

// Sha256 is a compatibility function that redirects to the appropriate implementation
func Sha256(bz []byte) []byte {
	if len(bz) == 0 {
		return nil
	}
	return crypto.Sha256(bz)
}

// Other crypto functions as needed
func RipemdHash(bz []byte) []byte {
	if len(bz) == 0 {
		return nil
	}
	return crypto.Ripemd160(bz)
}

func Sha512(bz []byte) []byte {
	if len(bz) == 0 {
		return nil
	}
	hasher := crypto.Sha512.New()
	hasher.Write(bz)
	return hasher.Sum(nil)
}

// Sha512New returns a new SHA-512 hasher
func Sha512New() hash.Hash {
	return crypto.Sha512.New()
}
