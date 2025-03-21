package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/depinnetwork/depin-sdk/codec"
	addresscodec "github.com/depinnetwork/depin-sdk/codec/address"
	"github.com/depinnetwork/depin-sdk/codec/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

func TestMsgUnjailGetSignBytes(t *testing.T) {
	addr := sdk.AccAddress("abcd")
	valAddrStr, err := addresscodec.NewBech32Codec("cosmosvaloper").BytesToString(addr)
	require.NoError(t, err)
	msg := NewMsgUnjail(valAddrStr)
	pc := codec.NewProtoCodec(types.NewInterfaceRegistry())
	bytes, err := pc.MarshalAminoJSON(msg)
	require.NoError(t, err)
	require.Equal(
		t,
		`{"type":"cosmos-sdk/MsgUnjail","value":{"address":"cosmosvaloper1v93xxeqhg9nn6"}}`,
		string(bytes),
	)
}
