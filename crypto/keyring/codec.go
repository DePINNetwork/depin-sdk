package keyring

import (
	"github.com/depinnetwork/depin-sdk/codec"
	"github.com/depinnetwork/depin-sdk/codec/legacy"
	"github.com/depinnetwork/depin-sdk/crypto/hd"
)

func init() {
	RegisterLegacyAminoCodec(legacy.Cdc)
}

// RegisterLegacyAminoCodec registers concrete types and interfaces on the given codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*LegacyInfo)(nil), nil)
	cdc.RegisterConcrete(hd.BIP44Params{}, "crypto/keys/hd/BIP44Params")
	cdc.RegisterConcrete(legacyLocalInfo{}, "crypto/keys/localInfo")
	cdc.RegisterConcrete(legacyLedgerInfo{}, "crypto/keys/ledgerInfo")
	cdc.RegisterConcrete(legacyOfflineInfo{}, "crypto/keys/offlineInfo")
	cdc.RegisterConcrete(LegacyMultiInfo{}, "crypto/keys/multiInfo")
}
