package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"cosmossdk.io/log"
	storetypes "github.com/depinnetwork/depin-sdk/store/types"

	"github.com/depinnetwork/depin-sdk/baseapp"
	"github.com/depinnetwork/depin-sdk/codec"
	"github.com/depinnetwork/depin-sdk/codec/testutil"
	"github.com/depinnetwork/depin-sdk/runtime"
	"github.com/depinnetwork/depin-sdk/testutil/testdata"
	"github.com/depinnetwork/depin-sdk/testutil/x/counter"
	counterkeeper "github.com/depinnetwork/depin-sdk/testutil/x/counter/keeper"
	countertypes "github.com/depinnetwork/depin-sdk/testutil/x/counter/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
	moduletestutil "github.com/depinnetwork/depin-sdk/types/module/testutil"
)

type IntegrationTestSuite struct {
	suite.Suite

	ctx           sdk.Context
	cdc           codec.Codec
	counterClient countertypes.QueryClient
	testClient    testdata.QueryClient
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	logger := log.NewNopLogger()
	keys := storetypes.NewKVStoreKeys(countertypes.StoreKey)
	cms := moduletestutil.CreateMultiStore(keys, logger)
	s.ctx = sdk.NewContext(cms, true, logger)
	cfg := moduletestutil.MakeTestEncodingConfig(testutil.CodecOptions{}, counter.AppModule{})
	s.cdc = cfg.Codec

	queryHelper := baseapp.NewQueryServerTestHelper(s.ctx, cfg.InterfaceRegistry)
	testdata.RegisterQueryServer(queryHelper, testdata.QueryImpl{})
	s.testClient = testdata.NewQueryClient(queryHelper)

	kvs := runtime.NewKVStoreService(keys[countertypes.StoreKey])
	counterKeeper := counterkeeper.NewKeeper(runtime.NewEnvironment(kvs, logger))
	countertypes.RegisterQueryServer(queryHelper, counterKeeper)
	s.counterClient = countertypes.NewQueryClient(queryHelper)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
}

func (s *IntegrationTestSuite) TestGRPCQuery() {
	// gRPC query to test service should work
	testRes, err := s.testClient.Echo(context.Background(), &testdata.EchoRequest{Message: "hello"})
	s.Require().NoError(err)
	s.Require().Equal("hello", testRes.Message)

	var header metadata.MD
	res, err := s.counterClient.GetCount(s.ctx, &countertypes.QueryGetCountRequest{}, grpc.Header(&header))
	s.Require().NoError(err)
	s.Require().Equal(int64(0), res.TotalCount)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
