package multisig

import (
	"github.com/depinnetwork/por-consensus/crypto/bls12381"

	"github.com/depinnetwork/depin-sdk/codec"
	"github.com/depinnetwork/depin-sdk/crypto/keys/bls12_381"
	"github.com/depinnetwork/depin-sdk/crypto/keys/ed25519"
	"github.com/depinnetwork/depin-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
)

// TODO: Figure out API for others to either add their own pubkey types, or
// to make verify / marshal accept a AminoCdc.
const (
	// PubKeyAminoRoute defines the amino route for a multisig threshold public key
	PubKeyAminoRoute = "tendermint/PubKeyMultisigThreshold"
)

// AminoCdc is being deprecated in the SDK. But even if you need to
// use Amino for some reason, please use `codec/legacy.Cdc` instead.
var AminoCdc = codec.NewLegacyAmino()

func init() {
	AminoCdc.RegisterInterface((*cryptotypes.PubKey)(nil), nil)
	AminoCdc.RegisterConcrete(ed25519.PubKey{},
		ed25519.PubKeyName)
	AminoCdc.RegisterConcrete(&secp256k1.PubKey{},
		secp256k1.PubKeyName)
	AminoCdc.RegisterConcrete(&bls12_381.PubKey{},
		bls12381.PubKeyName)
	AminoCdc.RegisterConcrete(&LegacyAminoPubKey{},
		PubKeyAminoRoute)
}
