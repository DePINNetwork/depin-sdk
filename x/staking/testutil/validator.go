package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/x/staking/types"

	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

// NewValidator is a testing helper method to create validators in tests
func NewValidator(tb testing.TB, operator sdk.ValAddress, pubKey cryptotypes.PubKey) types.Validator {
	tb.Helper()
	operatorAddr, err := codectestutil.CodecOptions{}.GetValidatorCodec().BytesToString(operator)
	require.NoError(tb, err)
	v, err := types.NewValidator(operatorAddr, pubKey, types.Description{})
	require.NoError(tb, err)
	return v
}
