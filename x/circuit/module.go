package circuit

import (
	"context"
	"encoding/json"
	"fmt"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/appmodule"
	appmodulev2 "cosmossdk.io/core/appmodule/v2"
	"cosmossdk.io/core/codec"
	"cosmossdk.io/core/registry"
	"cosmossdk.io/core/transaction"
	"cosmossdk.io/schema"
	"cosmossdk.io/x/circuit/ante"
	"cosmossdk.io/x/circuit/keeper"
	"cosmossdk.io/x/circuit/types"

	"github.com/depinnetwork/depin-sdk/client"
	"github.com/depinnetwork/depin-sdk/types/module"
)

// ConsensusVersion defines the current circuit module consensus version.
const ConsensusVersion = 1

var (
	_ module.HasGRPCGateway = AppModule{}

	_ appmodule.AppModule                        = AppModule{}
	_ appmodule.HasGenesis                       = AppModule{}
	_ appmodule.HasRegisterInterfaces            = AppModule{}
	_ appmodulev2.HasTxValidator[transaction.Tx] = AppModule{}
)

// AppModule implements an application module for the circuit module.
type AppModule struct {
	cdc    codec.Codec
	keeper keeper.Keeper
}

// IsAppModule implements the appmodule.AppModule interface.
func (AppModule) IsAppModule() {}

// Name returns the circuit module's name.
// Deprecated: kept for legacy reasons.
func (AppModule) Name() string { return types.ModuleName }

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the circuit module.
func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// RegisterInterfaces registers interfaces and implementations of the circuit module.
func (AppModule) RegisterInterfaces(registrar registry.InterfaceRegistrar) {
	types.RegisterInterfaces(registrar)
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(registrar grpc.ServiceRegistrar) error {
	types.RegisterMsgServer(registrar, keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(registrar, keeper.NewQueryServer(am.keeper))

	return nil
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		cdc:    cdc,
		keeper: keeper,
	}
}

// ConsensusVersion implements HasConsensusVersion
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// DefaultGenesis returns default genesis state as raw bytes for the circuit module.
func (am AppModule) DefaultGenesis() json.RawMessage {
	data, err := am.cdc.MarshalJSON(types.DefaultGenesisState())
	if err != nil {
		panic(err)
	}
	return data
}

// ValidateGenesis performs genesis state validation for the circuit module.
func (am AppModule) ValidateGenesis(bz json.RawMessage) error {
	var data types.GenesisState
	if err := am.cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return data.Validate()
}

// InitGenesis performs genesis initialization for the circuit module.
func (am AppModule) InitGenesis(ctx context.Context, data json.RawMessage) error {
	var genesisState types.GenesisState
	if err := am.cdc.UnmarshalJSON(data, &genesisState); err != nil {
		return err
	}

	return am.keeper.InitGenesis(ctx, &genesisState)
}

// ExportGenesis returns the exported genesis state as raw bytes for the circuit
// module.
func (am AppModule) ExportGenesis(ctx context.Context) (json.RawMessage, error) {
	gs, err := am.keeper.ExportGenesis(ctx)
	if err != nil {
		return nil, err
	}
	return am.cdc.MarshalJSON(gs)
}

// TxValidator implements appmodule.HasTxValidator.
func (am AppModule) TxValidator(ctx context.Context, tx transaction.Tx) error {
	validator := ante.NewCircuitBreakerDecorator(&am.keeper)
	return validator.ValidateTx(ctx, tx)
}

// ModuleCodec implements schema.HasModuleCodec.
// It allows the indexer to decode the module's KVPairUpdate.
func (am AppModule) ModuleCodec() (schema.ModuleCodec, error) {
	return am.keeper.Schema.ModuleCodec(collections.IndexingOptions{})
}
