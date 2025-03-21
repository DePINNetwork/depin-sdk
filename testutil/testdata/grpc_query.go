package testdata

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	gogoprotoany "github.com/cosmos/gogoproto/types/any"
	"github.com/cosmos/gogoproto/types/any/test"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"gotest.tools/v3/assert"

	"cosmossdk.io/core/gas"

	"github.com/depinnetwork/depin-sdk/codec/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

// iterCount defines the number of iterations to run on each query to test
// determinism.
var iterCount = 1000

type QueryImpl struct{}

var _ QueryServer = QueryImpl{}

func (e QueryImpl) TestAny(_ context.Context, request *TestAnyRequest) (*TestAnyResponse, error) {
	animal, ok := request.AnyAnimal.GetCachedValue().(test.Animal)
	if !ok {
		return nil, errors.New("expected Animal")
	}

	any, err := types.NewAnyWithValue(animal.(proto.Message))
	if err != nil {
		return nil, err
	}

	return &TestAnyResponse{HasAnimal: &HasAnimal{
		Animal: any,
		X:      10,
	}}, nil
}

func (e QueryImpl) Echo(_ context.Context, req *EchoRequest) (*EchoResponse, error) {
	return &EchoResponse{Message: req.Message}, nil
}

func (e QueryImpl) SayHello(_ context.Context, request *SayHelloRequest) (*SayHelloResponse, error) {
	greeting := fmt.Sprintf("Hello %s!", request.Name)
	return &SayHelloResponse{Greeting: greeting}, nil
}

var _ gogoprotoany.UnpackInterfacesMessage = &TestAnyRequest{}

func (m *TestAnyRequest) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	var animal test.Animal
	return unpacker.UnpackAny(m.AnyAnimal, &animal)
}

var _ gogoprotoany.UnpackInterfacesMessage = &TestAnyResponse{}

func (m *TestAnyResponse) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	return m.HasAnimal.UnpackInterfaces(unpacker)
}

// DeterministicIterations is a function to handle deterministic query requests. It tests 2 things:
// 1. That the response is always the same when calling the
// grpc query `iterCount` times (defaults to 1000).
// 2. That the gas consumption of the query is the same. When
// `gasOverwrite` is set to true, we also check that this consumed
// gas value is equal to the hardcoded `gasConsumed`.
func DeterministicIterations[request, response proto.Message](
	t *testing.T,
	ctx sdk.Context,
	req request,
	grpcFn func(context.Context, request, ...grpc.CallOption) (response, error),
	gasConsumed uint64,
	gasOverwrite bool,
) {
	t.Helper()
	before := ctx.GasMeter().GasConsumed()
	prevRes, err := grpcFn(ctx, req)
	assert.NilError(t, err)
	if gasOverwrite { // to handle regressions, i.e. check that gas consumption didn't change
		gasConsumed = ctx.GasMeter().GasConsumed() - before
	}
	t.Logf("gas consumed: %d", gasConsumed)

	for i := 0; i < iterCount; i++ {
		before := ctx.GasMeter().GasConsumed()
		res, err := grpcFn(ctx, req)
		assert.Equal(t, ctx.GasMeter().GasConsumed()-before, gasConsumed)
		assert.NilError(t, err)
		assert.DeepEqual(t, res, prevRes)
	}
}

func DeterministicIterationsV2[request, response proto.Message](
	t *testing.T,
	req request,
	meterFn func() gas.Meter,
	queryFn func(request) (response, error),
	assertGas func(*testing.T, gas.Gas),
	assertResponse func(*testing.T, response),
) {
	t.Helper()
	prevRes, err := queryFn(req)
	gasMeter := meterFn()
	gasConsumed := gasMeter.Consumed()
	require.NoError(t, err)
	assertGas(t, gasConsumed)

	for i := 0; i < iterCount; i++ {
		res, err := queryFn(req)
		require.NoError(t, err)
		sameGas := gasMeter.Consumed()
		require.Equal(t, gasConsumed, sameGas)
		require.Equal(t, res, prevRes)
		if assertResponse != nil {
			assertResponse(t, res)
		}
	}
}
