package mint_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	storetypes "github.com/depinnetwork/depin-sdk/store/types"
	"cosmossdk.io/x/mint"
	"cosmossdk.io/x/mint/keeper"
	minttestutil "cosmossdk.io/x/mint/testutil"
	"cosmossdk.io/x/mint/types"

	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	"github.com/depinnetwork/depin-sdk/runtime"
	"github.com/depinnetwork/depin-sdk/testutil"
	sdk "github.com/depinnetwork/depin-sdk/types"
	moduletestutil "github.com/depinnetwork/depin-sdk/types/module/testutil"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

const govModuleNameStr = "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"

type ModuleTestSuite struct {
	suite.Suite

	mintKeeper    *keeper.Keeper
	ctx           sdk.Context
	msgServer     types.MsgServer
	stakingKeeper *minttestutil.MockStakingKeeper
	bankKeeper    *minttestutil.MockBankKeeper

	appmodule mint.AppModule
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (s *ModuleTestSuite) SetupTest() {
	encCfg := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, mint.AppModule{})
	key := storetypes.NewKVStoreKey(types.StoreKey)
	storeService := runtime.NewKVStoreService(key)
	env := runtime.NewEnvironment(storeService, log.NewNopLogger())
	testCtx := testutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	s.ctx = testCtx.Ctx

	// gomock initializations
	ctrl := gomock.NewController(s.T())
	accountKeeper := minttestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := minttestutil.NewMockBankKeeper(ctrl)
	stakingKeeper := minttestutil.NewMockStakingKeeper(ctrl)

	accountKeeper.EXPECT().GetModuleAddress(types.ModuleName).Return(sdk.AccAddress{})

	s.mintKeeper = keeper.NewKeeper(
		encCfg.Codec,
		env,
		accountKeeper,
		bankKeeper,
		authtypes.FeeCollectorName,
		govModuleNameStr,
	)
	err := s.mintKeeper.SetMintFn(keeper.DefaultMintFn(types.DefaultInflationCalculationFn, stakingKeeper, s.mintKeeper))
	s.NoError(err)

	s.stakingKeeper = stakingKeeper
	s.bankKeeper = bankKeeper

	err = s.mintKeeper.Params.Set(s.ctx, types.DefaultParams())
	s.NoError(err)

	s.NoError(s.mintKeeper.Minter.Set(s.ctx, types.DefaultInitialMinter()))
	s.msgServer = keeper.NewMsgServerImpl(s.mintKeeper)

	s.appmodule = mint.NewAppModule(encCfg.Codec, s.mintKeeper, accountKeeper)
}

func (s *ModuleTestSuite) TestEpochHooks() {
	s.stakingKeeper.EXPECT().StakingTokenSupply(s.ctx).Return(math.NewIntFromUint64(100000000000), nil).AnyTimes()
	bondedRatio := math.LegacyNewDecWithPrec(15, 2)
	s.stakingKeeper.EXPECT().BondedRatio(s.ctx).Return(bondedRatio, nil).AnyTimes()
	s.bankKeeper.EXPECT().MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(792)))).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToModule(s.ctx, types.ModuleName, authtypes.FeeCollectorName, gomock.Any()).Return(nil)

	err := s.appmodule.BeforeEpochStart(s.ctx, "block", -1)
	s.NoError(err)

	err = s.appmodule.AfterEpochEnd(s.ctx, "epochIdentifier", 1) // just to get coverage up
	s.NoError(err)
}
