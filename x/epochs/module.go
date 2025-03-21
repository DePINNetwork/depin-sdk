package epochs

import (
	"context"
	"encoding/json"
	"fmt"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/codec"
	"cosmossdk.io/core/registry"
	"cosmossdk.io/schema"
	"cosmossdk.io/x/epochs/keeper"
	"cosmossdk.io/x/epochs/simulation"
	"cosmossdk.io/x/epochs/types"

	"github.com/depinnetwork/depin-sdk/client"
	"github.com/depinnetwork/depin-sdk/types/module"
	simtypes "github.com/depinnetwork/depin-sdk/types/simulation"
)

var (
	_ module.HasAminoCodec       = AppModule{}
	_ module.HasGRPCGateway      = AppModule{}
	_ module.AppModuleSimulation = AppModule{}
	_ module.HasGenesis          = AppModule{}

	_ appmodule.AppModule       = AppModule{}
	_ appmodule.HasBeginBlocker = AppModule{}
)

const ConsensusVersion = 1

// AppModule implements the AppModule interface for the epochs module.
type AppModule struct {
	cdc    codec.Codec
	keeper *keeper.Keeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(cdc codec.Codec, keeper *keeper.Keeper) AppModule {
	return AppModule{
		cdc:    cdc,
		keeper: keeper,
	}
}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// Name returns the epochs module's name.
// Deprecated: kept for legacy reasons.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the epochs module's types for the given codec.
func (AppModule) RegisterLegacyAminoCodec(registrar registry.AminoRegistrar) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the epochs module.
func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(registrar grpc.ServiceRegistrar) error {
	types.RegisterQueryServer(registrar, keeper.NewQuerier(*am.keeper))
	return nil
}

// DefaultGenesis returns the epochs module's default genesis state.
func (am AppModule) DefaultGenesis() json.RawMessage {
	data, err := am.cdc.MarshalJSON(types.DefaultGenesis())
	if err != nil {
		panic(err)
	}
	return data
}

// ValidateGenesis performs genesis state validation for the epochs module.
func (am AppModule) ValidateGenesis(bz json.RawMessage) error {
	var gs types.GenesisState
	if err := am.cdc.UnmarshalJSON(bz, &gs); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return gs.Validate()
}

// InitGenesis performs the epochs module's genesis initialization
func (am AppModule) InitGenesis(ctx context.Context, bz json.RawMessage) error {
	var gs types.GenesisState
	err := am.cdc.UnmarshalJSON(bz, &gs)
	if err != nil {
		return (fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err))
	}

	return am.keeper.InitGenesis(ctx, gs)
}

// ExportGenesis returns the epochs module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx context.Context) (json.RawMessage, error) {
	gs, err := am.keeper.ExportGenesis(ctx)
	if err != nil {
		return nil, err
	}
	return am.cdc.MarshalJSON(gs)
}

// ConsensusVersion implements HasConsensusVersion
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// BeginBlock executes all ABCI BeginBlock logic respective to the epochs module.
func (am AppModule) BeginBlock(ctx context.Context) error {
	return am.keeper.BeginBlocker(ctx)
}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the epochs module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// RegisterStoreDecoder registers a decoder for epochs module's types
func (am AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
	sdr[types.StoreKey] = simtypes.NewStoreDecoderFuncFromCollectionsSchema(am.keeper.Schema)
}

// ModuleCodec implements schema.HasModuleCodec.
// It allows the indexer to decode the module's KVPairUpdate.
func (am AppModule) ModuleCodec() (schema.ModuleCodec, error) {
	return am.keeper.Schema.ModuleCodec(collections.IndexingOptions{})
}
