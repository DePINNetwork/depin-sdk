package legacy

import (
	"github.com/depinnetwork/depin-sdk/codec"
	cryptocodec "github.com/depinnetwork/depin-sdk/crypto/codec"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

// Cdc defines a global generic sealed Amino codec to be used throughout sdk. It
// has all CometBFT crypto and evidence types registered.
//
// TODO: Deprecated - remove this global.
var Cdc = codec.NewLegacyAmino()

func init() {
	cryptocodec.RegisterCrypto(Cdc)
	codec.RegisterEvidences(Cdc)
	sdk.RegisterLegacyAminoCodec(Cdc)
}

// PrivKeyFromBytes unmarshals private key bytes and returns a PrivKey
func PrivKeyFromBytes(privKeyBytes []byte) (privKey cryptotypes.PrivKey, err error) {
	err = Cdc.Unmarshal(privKeyBytes, &privKey)
	return
}

// PubKeyFromBytes unmarshals public key bytes and returns a PubKey
func PubKeyFromBytes(pubKeyBytes []byte) (pubKey cryptotypes.PubKey, err error) {
	err = Cdc.Unmarshal(pubKeyBytes, &pubKey)
	return
}
