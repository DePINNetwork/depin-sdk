package keyring

import (
	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"

	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
)

// Keyring is an interface used for signing transactions.
// It aims to be simplistic and easy to use.
type Keyring interface {
	// List returns the names of all keys stored in the keyring.
	List() ([]string, error)

	// LookupAddressByKeyName returns the address of the key with the given name.
	LookupAddressByKeyName(name string) ([]byte, error)

	// GetPubKey returns the public key of the key with the given name.
	GetPubKey(name string) (cryptotypes.PubKey, error)

	// Sign signs the given bytes with the key with the given name.
	Sign(name string, msg []byte, signMode signingv1beta1.SignMode) ([]byte, error)

	// KeyType returns the type of the key.
	KeyType(name string) (uint, error)

	// KeyInfo given a key name or address returns key name, key address and key type.
	KeyInfo(nameOrAddr string) (string, string, uint, error)
}
