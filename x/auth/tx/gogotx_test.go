package tx

import (
	"testing"

	"github.com/stretchr/testify/require"

	v1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	"cosmossdk.io/x/tx/decode"

	"github.com/depinnetwork/depin-sdk/codec"
	addresscodec "github.com/depinnetwork/depin-sdk/codec/address"
	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

func TestNewWrapperFromDecodedTxAllowsNilFee(t *testing.T) {
	addrCodec := addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
	cdc := codec.NewProtoCodec(codectestutil.CodecOptions{}.NewInterfaceRegistry())
	_, err := newWrapperFromDecodedTx(addrCodec, cdc, &decode.DecodedTx{
		Tx: &v1beta1.Tx{
			AuthInfo: &v1beta1.AuthInfo{},
		},
	})
	require.Nil(t, err)
}
