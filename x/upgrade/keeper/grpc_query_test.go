package keeper_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/header"
	coretesting "cosmossdk.io/core/testing"
	"cosmossdk.io/core/testing/queryclient"
	"cosmossdk.io/x/upgrade"
	"cosmossdk.io/x/upgrade/keeper"
	upgradetestutil "cosmossdk.io/x/upgrade/testutil"
	"cosmossdk.io/x/upgrade/types"

	"github.com/depinnetwork/depin-sdk/codec"
	addresscodec "github.com/depinnetwork/depin-sdk/codec/address"
	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	moduletestutil "github.com/depinnetwork/depin-sdk/types/module/testutil"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

type UpgradeTestSuite struct {
	suite.Suite

	upgradeKeeper    *keeper.Keeper
	ctx              coretesting.TestContext
	env              coretesting.TestEnvironment
	queryClient      types.QueryClient
	encCfg           moduletestutil.TestEncodingConfig
	encodedAuthority string
}

func (suite *UpgradeTestSuite) SetupTest() {
	suite.encCfg = moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, upgrade.AppModule{})

	ctx, env := coretesting.NewTestEnvironment(coretesting.TestEnvironmentConfig{
		ModuleName: types.ModuleName,
		Logger:     coretesting.NewNopLogger(),
	})

	suite.ctx = ctx
	suite.env = env

	skipUpgradeHeights := make(map[int64]bool)
	authority, err := addresscodec.NewBech32Codec("cosmos").BytesToString(authtypes.NewModuleAddress(types.GovModuleName))
	suite.Require().NoError(err)
	suite.encodedAuthority = authority
	ctrl := gomock.NewController(suite.T())
	ck := upgradetestutil.NewMockConsensusKeeper(ctrl)
	suite.upgradeKeeper = keeper.NewKeeper(suite.env.Environment, skipUpgradeHeights, suite.encCfg.Codec, suite.T().TempDir(), nil, authority, ck)
	err = suite.upgradeKeeper.SetModuleVersionMap(suite.ctx, appmodule.VersionMap{
		"bank": 0,
	})
	suite.Require().NoError(err)
	queryHelper := queryclient.NewQueryHelper(codec.NewProtoCodec(suite.encCfg.InterfaceRegistry).GRPCCodec())
	types.RegisterQueryServer(queryHelper, suite.upgradeKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func (suite *UpgradeTestSuite) TestQueryCurrentPlan() {
	var (
		req         *types.QueryCurrentPlanRequest
		expResponse types.QueryCurrentPlanResponse
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"without current upgrade plan",
			func() {
				req = &types.QueryCurrentPlanRequest{}
				expResponse = types.QueryCurrentPlanResponse{}
			},
			true,
		},
		{
			"with current upgrade plan",
			func() {
				plan := types.Plan{Name: "test-plan", Height: 5}
				err := suite.upgradeKeeper.ScheduleUpgrade(suite.ctx, plan)
				suite.Require().NoError(err)
				req = &types.QueryCurrentPlanRequest{}
				expResponse = types.QueryCurrentPlanResponse{Plan: &plan}
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			tc.malleate()

			res, err := suite.queryClient.CurrentPlan(suite.ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(&expResponse, res)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *UpgradeTestSuite) TestAppliedCurrentPlan() {
	var (
		req       *types.QueryAppliedPlanRequest
		expHeight int64
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"with non-existent upgrade plan",
			func() {
				req = &types.QueryAppliedPlanRequest{Name: "foo"}
			},
			true,
		},
		{
			"with applied upgrade plan",
			func() {
				expHeight = 5

				planName := "test-plan"
				plan := types.Plan{Name: planName, Height: expHeight}
				err := suite.upgradeKeeper.ScheduleUpgrade(suite.ctx, plan)
				suite.Require().NoError(err)
				suite.ctx = suite.ctx.WithHeaderInfo(header.Info{Height: expHeight})
				suite.upgradeKeeper.SetUpgradeHandler(planName, func(ctx context.Context, plan types.Plan, vm appmodule.VersionMap) (appmodule.VersionMap, error) {
					return vm, nil
				})
				err = suite.upgradeKeeper.ApplyUpgrade(suite.ctx, plan)
				suite.Require().NoError(err)
				req = &types.QueryAppliedPlanRequest{Name: planName}
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			tc.malleate()

			res, err := suite.queryClient.AppliedPlan(suite.ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(expHeight, res.Height)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *UpgradeTestSuite) TestModuleVersions() {
	testCases := []struct {
		msg     string
		req     types.QueryModuleVersionsRequest
		single  bool
		expPass bool
	}{
		{
			msg:     "test full query",
			req:     types.QueryModuleVersionsRequest{},
			single:  false,
			expPass: true,
		},
		{
			msg:     "test single module",
			req:     types.QueryModuleVersionsRequest{ModuleName: "bank"},
			single:  true,
			expPass: true,
		},
		{
			msg:     "test non-existent module",
			req:     types.QueryModuleVersionsRequest{ModuleName: "abcdefg"},
			single:  true,
			expPass: false,
		},
	}

	vm, err := suite.upgradeKeeper.GetModuleVersionMap(suite.ctx)
	suite.Require().NoError(err)

	mv, err := suite.upgradeKeeper.GetModuleVersions(suite.ctx)
	suite.Require().NoError(err)

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			res, err := suite.queryClient.ModuleVersions(suite.ctx, &tc.req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				if tc.single {
					// test that the single module response is valid
					suite.Require().Len(res.ModuleVersions, 1)
					// make sure we got the right values
					suite.Require().Equal(vm[tc.req.ModuleName], res.ModuleVersions[0].Version)
					suite.Require().Equal(tc.req.ModuleName, res.ModuleVersions[0].Name)
				} else {
					// check that the full response is valid
					suite.Require().NotEmpty(res.ModuleVersions)
					suite.Require().Equal(len(mv), len(res.ModuleVersions))
					for i, v := range res.ModuleVersions {
						suite.Require().Equal(mv[i].Version, v.Version)
						suite.Require().Equal(mv[i].Name, v.Name)
					}
				}
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *UpgradeTestSuite) TestAuthority() {
	res, err := suite.queryClient.Authority(context.Background(), &types.QueryAuthorityRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(suite.encodedAuthority, res.Address)
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}
