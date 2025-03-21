package keys

import (
	bls "github.com/depinnetwork/por-consensus/crypto/bls12381"

	"github.com/depinnetwork/depin-sdk/crypto/keys/bls12_381"
	"github.com/depinnetwork/depin-sdk/crypto/keys/ed25519"
	"github.com/depinnetwork/depin-sdk/crypto/keys/secp256k1"
	"github.com/depinnetwork/depin-sdk/crypto/types"
)

// JSONPubKey defines a public key that are parse from JSON file.
// convert PubKey to JSONPubKey needs a in between step
type JSONPubkey struct {
	KeyType string `json:"type"`
	Value   []byte `json:"value"`
}

func (pk JSONPubkey) Address() types.Address {
	switch pk.KeyType {
	case ed25519.PubKeyName:
		ed25519 := ed25519.PubKey{
			Key: pk.Value,
		}
		return ed25519.Address()
	case secp256k1.PubKeyName:
		secp256k1 := secp256k1.PubKey{
			Key: pk.Value,
		}
		return secp256k1.Address()
	case bls.PubKeyName:
		bls12_381 := bls12_381.PubKey{
			Key: pk.Value,
		}
		return bls12_381.Address()
	default:
		return nil
	}
}

func (pk JSONPubkey) Bytes() []byte {
	return pk.Value
}
