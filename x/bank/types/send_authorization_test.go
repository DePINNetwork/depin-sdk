package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	coregas "cosmossdk.io/core/gas"
	coretesting "cosmossdk.io/core/testing"
	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/x/bank/types"

	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

var (
	coins1000      = sdk.NewCoins(sdk.NewCoin("stake", sdkmath.NewInt(1000)))
	coins500       = sdk.NewCoins(sdk.NewCoin("stake", sdkmath.NewInt(500)))
	toAddr         = sdk.AccAddress("_______to________")
	fromAddrStr    = "cosmos1ta047h6lveex7mfqta047h6ln9jal0"
	toAddrStr      = "cosmos1ta047h6lta0hgm6lta047h6lta0stgm2m3"
	unknownAddrStr = "cosmos1ta047h6lw4hxkmn0wah97h6lta0sml880l"
)

type mockGasMeter struct {
	coregas.Meter
}

func (m mockGasMeter) Consume(_ coregas.Gas, _ string) error {
	return nil
}

func TestSendAuthorization(t *testing.T) {
	ac := codectestutil.CodecOptions{}.GetAddressCodec()
	cfg := coretesting.TestEnvironmentConfig{
		ModuleName: "bank",
		Logger:     log.NewNopLogger(),
	}

	ctx, _ := coretesting.NewTestEnvironment(cfg)
	ctx = ctx.WithGas(coregas.GasConfig{}, mockGasMeter{})

	allowList := make([]sdk.AccAddress, 1)
	allowList[0] = toAddr
	authorization := types.NewSendAuthorization(coins1000, nil, ac)

	t.Log("verify authorization returns valid method name")
	require.Equal(t, authorization.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send := types.NewMsgSend(fromAddrStr, toAddrStr, coins1000)

	t.Log("verify updated authorization returns nil")
	resp, err := authorization.Accept(ctx, send)
	require.NoError(t, err)
	require.True(t, resp.Delete)
	require.Nil(t, resp.Updated)

	authorization = types.NewSendAuthorization(coins1000, nil, ac)
	require.Equal(t, authorization.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddrStr, toAddrStr, coins500)
	require.NoError(t, authorization.ValidateBasic())
	resp, err = authorization.Accept(ctx, send)

	t.Log("verify updated authorization returns remaining spent limit")
	require.NoError(t, err)
	require.False(t, resp.Delete)
	require.NotNil(t, resp.Updated)
	sendAuth := types.NewSendAuthorization(coins500, nil, ac)
	require.Equal(t, sendAuth.String(), resp.Updated.String())

	t.Log("expect updated authorization nil after spending remaining amount")
	resp, err = resp.Updated.(*types.SendAuthorization).Accept(ctx, send)
	require.NoError(t, err)
	require.True(t, resp.Delete)
	require.Nil(t, resp.Updated)

	t.Log("allow list and no address")
	authzWithAllowList := types.NewSendAuthorization(coins1000, allowList, ac)
	require.Equal(t, authzWithAllowList.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddrStr, unknownAddrStr, coins500)
	require.NoError(t, authzWithAllowList.ValidateBasic())
	resp, err = authzWithAllowList.Accept(ctx, send)
	require.False(t, resp.Accept)
	require.False(t, resp.Delete)
	require.Nil(t, resp.Updated)
	require.Error(t, err)
	require.Contains(t, err.Error(), fmt.Sprintf("cannot send to %s address", unknownAddrStr))

	t.Log("send to address in allow list")
	authzWithAllowList = types.NewSendAuthorization(coins1000, allowList, ac)
	require.Equal(t, authzWithAllowList.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddrStr, toAddrStr, coins500)
	require.NoError(t, authzWithAllowList.ValidateBasic())
	resp, err = authzWithAllowList.Accept(ctx, send)
	require.NoError(t, err)
	require.True(t, resp.Accept)
	require.NotNil(t, resp.Updated)
	// coins1000-coins500 = coins500
	require.Equal(t, types.NewSendAuthorization(coins500, allowList, ac).String(), resp.Updated.String())

	t.Log("send everything to address not in allow list")
	authzWithAllowList = types.NewSendAuthorization(coins1000, allowList, ac)
	require.Equal(t, authzWithAllowList.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddrStr, unknownAddrStr, coins1000)
	require.NoError(t, authzWithAllowList.ValidateBasic())
	resp, err = authzWithAllowList.Accept(ctx, send)
	require.Error(t, err)
	require.Contains(t, err.Error(), fmt.Sprintf("cannot send to %s address", unknownAddrStr))

	t.Log("send everything to address in allow list")
	authzWithAllowList = types.NewSendAuthorization(coins1000, allowList, ac)
	require.Equal(t, authzWithAllowList.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddrStr, toAddrStr, coins1000)
	require.NoError(t, authzWithAllowList.ValidateBasic())
	resp, err = authzWithAllowList.Accept(ctx, send)
	require.NoError(t, err)
	require.True(t, resp.Accept)
	require.Nil(t, resp.Updated)
}
