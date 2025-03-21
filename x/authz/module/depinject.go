package module

import (
	modulev1 "cosmossdk.io/api/cosmos/authz/module/v1"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"
	"cosmossdk.io/x/authz/keeper"

	"github.com/depinnetwork/depin-sdk/codec"
	cdctypes "github.com/depinnetwork/depin-sdk/codec/types"
)

var _ depinject.OnePerModuleType = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

func init() {
	appconfig.RegisterModule(
		&modulev1.Module{},
		appconfig.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	AddressCodec address.Codec
	Registry     cdctypes.InterfaceRegistry
	Environment  appmodule.Environment
}

type ModuleOutputs struct {
	depinject.Out

	AuthzKeeper keeper.Keeper
	Module      appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	k := keeper.NewKeeper(in.Environment, in.Cdc, in.AddressCodec)
	m := NewAppModule(in.Cdc, k, in.Registry)
	return ModuleOutputs{AuthzKeeper: k, Module: m}
}
