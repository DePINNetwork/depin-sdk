package keeper_test

import (
	gocontext "context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"cosmossdk.io/log"
	storetypes "github.com/depinnetwork/depin-sdk/store/types"
	"cosmossdk.io/x/mint"
	"cosmossdk.io/x/mint/keeper"
	minttestutil "cosmossdk.io/x/mint/testutil"
	"cosmossdk.io/x/mint/types"

	"github.com/depinnetwork/depin-sdk/baseapp"
	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	"github.com/depinnetwork/depin-sdk/runtime"
	"github.com/depinnetwork/depin-sdk/testutil"
	sdk "github.com/depinnetwork/depin-sdk/types"
	moduletestutil "github.com/depinnetwork/depin-sdk/types/module/testutil"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

type MintTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	queryClient types.QueryClient
	mintKeeper  *keeper.Keeper
}

func (suite *MintTestSuite) SetupTest() {
	encCfg := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, mint.AppModule{})
	key := storetypes.NewKVStoreKey(types.StoreKey)
	env := runtime.NewEnvironment(runtime.NewKVStoreService(key), log.NewNopLogger())
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	suite.ctx = testCtx.Ctx

	// gomock initializations
	ctrl := gomock.NewController(suite.T())
	accountKeeper := minttestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := minttestutil.NewMockBankKeeper(ctrl)

	accountKeeper.EXPECT().GetModuleAddress("mint").Return(sdk.AccAddress{})

	suite.mintKeeper = keeper.NewKeeper(
		encCfg.Codec,
		env,
		accountKeeper,
		bankKeeper,
		authtypes.FeeCollectorName,
		govModuleNameStr,
	)

	err := suite.mintKeeper.Params.Set(suite.ctx, types.DefaultParams())
	suite.Require().NoError(err)
	suite.Require().NoError(suite.mintKeeper.Minter.Set(suite.ctx, types.DefaultInitialMinter()))
	queryHelper := baseapp.NewQueryServerTestHelper(testCtx.Ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, keeper.NewQueryServerImpl(suite.mintKeeper))

	suite.queryClient = types.NewQueryClient(queryHelper)
}

func (suite *MintTestSuite) TestGRPCParams() {
	params, err := suite.queryClient.Params(gocontext.Background(), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	kparams, err := suite.mintKeeper.Params.Get(suite.ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(params.Params, kparams)

	inflation, err := suite.queryClient.Inflation(gocontext.Background(), &types.QueryInflationRequest{})
	suite.Require().NoError(err)
	minter, err := suite.mintKeeper.Minter.Get(suite.ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(inflation.Inflation, minter.Inflation)

	annualProvisions, err := suite.queryClient.AnnualProvisions(gocontext.Background(), &types.QueryAnnualProvisionsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(annualProvisions.AnnualProvisions, minter.AnnualProvisions)
}

func TestMintTestSuite(t *testing.T) {
	suite.Run(t, new(MintTestSuite))
}
