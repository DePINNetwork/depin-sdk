package std

import (
	"cosmossdk.io/core/registry"

	"github.com/depinnetwork/depin-sdk/codec"
	cryptocodec "github.com/depinnetwork/depin-sdk/crypto/codec"
	sdk "github.com/depinnetwork/depin-sdk/types"
	txtypes "github.com/depinnetwork/depin-sdk/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(registrar registry.AminoRegistrar) {
	sdk.RegisterLegacyAminoCodec(registrar)
	cryptocodec.RegisterCrypto(registrar)
	codec.RegisterEvidences(registrar)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry registry.InterfaceRegistrar) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
