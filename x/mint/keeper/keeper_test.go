package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"cosmossdk.io/core/appmodule"
	coretesting "cosmossdk.io/core/testing"
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

type KeeperTestSuite struct {
	suite.Suite

	mintKeeper    *keeper.Keeper
	ctx           sdk.Context
	msgServer     types.MsgServer
	stakingKeeper *minttestutil.MockStakingKeeper
	bankKeeper    *minttestutil.MockBankKeeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	encCfg := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, mint.AppModule{})
	key := storetypes.NewKVStoreKey(types.StoreKey)
	storeService := runtime.NewKVStoreService(key)
	env := runtime.NewEnvironment(storeService, coretesting.NewNopLogger())
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
	s.stakingKeeper = stakingKeeper
	s.bankKeeper = bankKeeper

	err := s.mintKeeper.Params.Set(s.ctx, types.DefaultParams())
	s.NoError(err)

	s.NoError(s.mintKeeper.Minter.Set(s.ctx, types.DefaultInitialMinter()))
	s.msgServer = keeper.NewMsgServerImpl(s.mintKeeper)
}

func (s *KeeperTestSuite) TestDefaultMintFn() {
	s.stakingKeeper.EXPECT().StakingTokenSupply(s.ctx).Return(math.NewIntFromUint64(100000000000), nil).AnyTimes()
	bondedRatio := math.LegacyNewDecWithPrec(15, 2)
	s.stakingKeeper.EXPECT().BondedRatio(s.ctx).Return(bondedRatio, nil).AnyTimes()
	s.bankKeeper.EXPECT().MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(792)))).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToModule(s.ctx, types.ModuleName, authtypes.FeeCollectorName, gomock.Any()).Return(nil)
	err := s.mintKeeper.SetMintFn(keeper.DefaultMintFn(types.DefaultInflationCalculationFn, s.stakingKeeper, s.mintKeeper))
	s.NoError(err)

	minter, err := s.mintKeeper.Minter.Get(s.ctx)
	s.NoError(err)

	err = s.mintKeeper.MintFn(s.ctx, &minter, "block", 0)
	s.NoError(err)

	// set a maxSupply and call again. totalSupply will be bigger than maxSupply.
	params, err := s.mintKeeper.Params.Get(s.ctx)
	s.NoError(err)
	params.MaxSupply = math.NewInt(10000000000)
	err = s.mintKeeper.Params.Set(s.ctx, params)
	s.NoError(err)

	err = s.mintKeeper.MintFn(s.ctx, &minter, "block", 0)
	s.NoError(err)

	// modify max supply to be almost reached
	// modify blocksPerYear to mint 2053 coins per block
	// we tried to mint 2053stake, but we only get to mint 2000stake
	params, err = s.mintKeeper.Params.Get(s.ctx)
	s.NoError(err)
	params.MaxSupply = math.NewInt(100000000000 + 2000)
	params.BlocksPerYear = 2434275
	err = s.mintKeeper.Params.Set(s.ctx, params)
	s.NoError(err)

	s.bankKeeper.EXPECT().MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(2000)))).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToModule(s.ctx, types.ModuleName, authtypes.FeeCollectorName, gomock.Any()).Return(nil)

	err = s.mintKeeper.MintFn(s.ctx, &minter, "block", 0)
	s.NoError(err)
}

func (s *KeeperTestSuite) TestBeginBlocker() {
	s.stakingKeeper.EXPECT().StakingTokenSupply(s.ctx).Return(math.NewIntFromUint64(100000000000), nil).AnyTimes()
	bondedRatio := math.LegacyNewDecWithPrec(15, 2)
	s.stakingKeeper.EXPECT().BondedRatio(s.ctx).Return(bondedRatio, nil).AnyTimes()
	s.bankKeeper.EXPECT().MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(792)))).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToModule(s.ctx, types.ModuleName, authtypes.FeeCollectorName, gomock.Any()).Return(nil)
	err := s.mintKeeper.SetMintFn(keeper.DefaultMintFn(types.DefaultInflationCalculationFn, s.stakingKeeper, s.mintKeeper))
	s.NoError(err)
	// get minter (it should get modified afterwards)
	minter, err := s.mintKeeper.Minter.Get(s.ctx)
	s.NoError(err)

	err = s.mintKeeper.BeginBlocker(s.ctx)
	s.NoError(err)

	// get minter again and compare
	newMinter, err := s.mintKeeper.Minter.Get(s.ctx)
	s.NoError(err)
	s.NotEqual(minter, newMinter)

	// now use a mintfn that doesn't do anything
	err = s.mintKeeper.SetMintFn(func(ctx context.Context, env appmodule.Environment, minter *types.Minter, epochId string, epochNumber int64) error {
		return nil
	})
	s.NoError(err)
	err = s.mintKeeper.BeginBlocker(s.ctx)
	s.NoError(err)

	// get minter again and compare
	unchangedMinter, err := s.mintKeeper.Minter.Get(s.ctx)
	s.NoError(err)
	s.Equal(newMinter, unchangedMinter)
}

func (s *KeeperTestSuite) TestMigrator() {
	m := keeper.NewMigrator(s.mintKeeper)
	s.NoError(m.Migrate1to2(s.ctx)) // just to get the coverage up

	// set max supply to one and migrate (should get it to zero)
	params, err := s.mintKeeper.Params.Get(s.ctx)
	s.NoError(err)
	params.MaxSupply = math.OneInt()
	s.NoError(s.mintKeeper.Params.Set(s.ctx, params))

	s.NoError(m.Migrate2to3(s.ctx))

	newParams, err := s.mintKeeper.Params.Get(s.ctx)
	s.NoError(err)
	s.Equal(math.ZeroInt(), newParams.MaxSupply)
}
