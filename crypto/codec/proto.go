package codec

import (
	"cosmossdk.io/core/registry"

	"github.com/depinnetwork/depin-sdk/crypto/keys/bls12_381"
	"github.com/depinnetwork/depin-sdk/crypto/keys/ed25519"
	"github.com/depinnetwork/depin-sdk/crypto/keys/multisig"
	"github.com/depinnetwork/depin-sdk/crypto/keys/secp256k1"
	"github.com/depinnetwork/depin-sdk/crypto/keys/secp256r1"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry registry.InterfaceRegistrar) {
	var pk *cryptotypes.PubKey
	registry.RegisterInterface("cosmos.crypto.PubKey", pk)
	registry.RegisterImplementations(pk, &ed25519.PubKey{})
	registry.RegisterImplementations(pk, &secp256k1.PubKey{})
	registry.RegisterImplementations(pk, &bls12_381.PubKey{})
	registry.RegisterImplementations(pk, &multisig.LegacyAminoPubKey{})

	var priv *cryptotypes.PrivKey
	registry.RegisterInterface("cosmos.crypto.PrivKey", priv)
	registry.RegisterImplementations(priv, &secp256k1.PrivKey{})
	registry.RegisterImplementations(priv, &ed25519.PrivKey{})
	registry.RegisterImplementations(priv, &bls12_381.PrivKey{})
	secp256r1.RegisterInterfaces(registry)
}
