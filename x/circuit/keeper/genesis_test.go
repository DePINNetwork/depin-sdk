package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	coretesting "cosmossdk.io/core/testing"
	storetypes "github.com/depinnetwork/depin-sdk/store/types"
	"cosmossdk.io/x/circuit"
	"cosmossdk.io/x/circuit/keeper"
	"cosmossdk.io/x/circuit/types"

	"github.com/depinnetwork/depin-sdk/codec"
	addresscodec "github.com/depinnetwork/depin-sdk/codec/address"
	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	"github.com/depinnetwork/depin-sdk/runtime"
	"github.com/depinnetwork/depin-sdk/testutil"
	moduletestutil "github.com/depinnetwork/depin-sdk/types/module/testutil"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx       context.Context
	keeper    keeper.Keeper
	cdc       *codec.ProtoCodec
	addrBytes []byte
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (s *GenesisTestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(types.ModuleName)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	encCfg := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, circuit.AppModule{})

	sdkCtx := testCtx.Ctx
	s.ctx = sdkCtx
	s.cdc = codec.NewProtoCodec(encCfg.InterfaceRegistry)
	ac := addresscodec.NewBech32Codec("cosmos")

	authority, err := ac.BytesToString(authtypes.NewModuleAddress(types.GovModuleName))
	s.Require().NoError(err)

	bz, err := ac.StringToBytes(authority)
	s.Require().NoError(err)
	s.addrBytes = bz

	s.keeper = keeper.NewKeeper(runtime.NewEnvironment(runtime.NewKVStoreService(key), coretesting.NewNopLogger()), s.cdc, authority, ac)
}

func (s *GenesisTestSuite) TestInitExportGenesis() {
	perms := types.Permissions{
		Level:         3,
		LimitTypeUrls: []string{"test"},
	}
	err := s.keeper.Permissions.Set(s.ctx, s.addrBytes, perms)
	s.Require().NoError(err)

	var accounts []*types.GenesisAccountPermissions
	addr, err := addresscodec.NewBech32Codec("cosmos").BytesToString(s.addrBytes)
	s.Require().NoError(err)
	genAccsPerms := types.GenesisAccountPermissions{
		Address:     addr,
		Permissions: &perms,
	}
	accounts = append(accounts, &genAccsPerms)

	url := "test_url"

	genesisState := &types.GenesisState{
		AccountPermissions: accounts,
		DisabledTypeUrls:   []string{url},
	}

	err = s.keeper.InitGenesis(s.ctx, genesisState)
	s.Require().NoError(err)

	exported, err := s.keeper.ExportGenesis(s.ctx)
	s.Require().NoError(err)
	bz, err := s.cdc.MarshalJSON(exported)
	s.Require().NoError(err)

	var exportedGenesisState types.GenesisState
	err = s.cdc.UnmarshalJSON(bz, &exportedGenesisState)
	s.Require().NoError(err)

	s.Require().Equal(genesisState.AccountPermissions, exportedGenesisState.AccountPermissions)
	s.Require().Equal(genesisState.DisabledTypeUrls, exportedGenesisState.DisabledTypeUrls)
}
