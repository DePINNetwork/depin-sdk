package signing_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	apisigning "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txsigning "cosmossdk.io/x/tx/signing"

	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	"github.com/depinnetwork/depin-sdk/testutil/testdata"
	moduletestutil "github.com/depinnetwork/depin-sdk/types/module/testutil"
	authsign "github.com/depinnetwork/depin-sdk/x/auth/signing"
)

func TestGetSignBytesAdapterNoPublicKey(t *testing.T) {
	encodingConfig := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{})
	txConfig := encodingConfig.TxConfig
	_, _, addr := testdata.KeyTestPubAddr()
	signerData := txsigning.SignerData{
		Address:       addr.String(),
		ChainID:       "test-chain",
		AccountNumber: 11,
		Sequence:      15,
	}
	w := txConfig.NewTxBuilder()
	_, err := authsign.GetSignBytesAdapter(
		context.Background(),
		txConfig.SignModeHandler(),
		apisigning.SignMode_SIGN_MODE_DIRECT,
		signerData,
		w.GetTx())
	require.NoError(t, err)
}
