package sims

import (
	"context"
	"math/rand"
	"testing"
	"time"

	types2 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"

	"cosmossdk.io/core/header"
	"cosmossdk.io/errors"
	txsigning "cosmossdk.io/x/tx/signing"

	"github.com/depinnetwork/depin-sdk/baseapp"
	"github.com/depinnetwork/depin-sdk/client"
	codectypes "github.com/depinnetwork/depin-sdk/codec/types"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
	"github.com/depinnetwork/depin-sdk/types/simulation"
	"github.com/depinnetwork/depin-sdk/types/tx/signing"
	authsign "github.com/depinnetwork/depin-sdk/x/auth/signing"
)

// GenSignedMockTx generates a signed mock transaction.
func GenSignedMockTx(
	r *rand.Rand,
	txConfig client.TxConfig,
	msgs []sdk.Msg,
	feeAmt sdk.Coins,
	gas uint64,
	chainID string,
	accNums, accSeqs []uint64,
	priv ...cryptotypes.PrivKey,
) (sdk.Tx, error) {
	sigs := make([]signing.SignatureV2, len(priv))

	// create a random length memo
	memo := simulation.RandStringOfLength(r, simulation.RandIntBetween(r, 0, 100))

	// 1st round: set SignatureV2 with empty signatures, to set correct
	// signer infos.
	for i, p := range priv {
		sigs[i] = signing.SignatureV2{
			PubKey: p.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode: txConfig.SignModeHandler().DefaultMode(),
			},
			Sequence: accSeqs[i],
		}
	}

	tx := txConfig.NewTxBuilder()
	err := tx.SetMsgs(msgs...)
	if err != nil {
		return nil, err
	}

	err = tx.SetSignatures(sigs...)
	if err != nil {
		return nil, err
	}
	tx.SetMemo(memo)
	tx.SetFeeAmount(feeAmt)
	tx.SetGasLimit(gas)

	// 2nd round: once all signer infos are set, every signer can sign.
	for i, p := range priv {
		anyPk, err := codectypes.NewAnyWithValue(p.PubKey())
		if err != nil {
			return nil, err
		}

		signerData := txsigning.SignerData{
			Address:       sdk.AccAddress(p.PubKey().Address()).String(),
			ChainID:       chainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
			PubKey:        &anypb.Any{TypeUrl: anyPk.TypeUrl, Value: anyPk.Value},
		}

		signBytes, err := authsign.GetSignBytesAdapter(
			context.Background(), txConfig.SignModeHandler(), txConfig.SignModeHandler().DefaultMode(), signerData,
			tx.GetTx())
		if err != nil {
			panic(err)
		}
		sig, err := p.Sign(signBytes)
		if err != nil {
			panic(err)
		}
		sigs[i].Data.(*signing.SingleSignatureData).Signature = sig
	}
	err = tx.SetSignatures(sigs...)
	if err != nil {
		panic(err)
	}

	return tx.GetTx(), nil
}

// SignCheckDeliver checks a generated signed transaction and simulates a
// block commitment with the given transaction. A test assertion is made using
// the parameter 'expPass' against the result. A corresponding result is
// returned.
func SignCheckDeliver(
	t *testing.T, txCfg client.TxConfig, app *baseapp.BaseApp, header header.Info, msgs []sdk.Msg,
	chainID string, accNums, accSeqs []uint64, expSimPass, expPass bool, priv ...cryptotypes.PrivKey,
) (sdk.GasInfo, *sdk.Result, error) {
	t.Helper()
	tx, err := GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txCfg,
		msgs,
		sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)},
		DefaultGenTxGas,
		chainID,
		accNums,
		accSeqs,
		priv...,
	)
	require.NoError(t, err)
	txBytes, err := txCfg.TxEncoder()(tx)
	require.Nil(t, err)

	// Must simulate now as CheckTx doesn't run Msgs anymore
	_, res, err := app.Simulate(txBytes)

	if expSimPass {
		require.NoError(t, err)
		require.NotNil(t, res)
	} else {
		require.Error(t, err)
		require.Nil(t, res)
	}

	bz, err := txCfg.TxEncoder()(tx)
	require.NoError(t, err)

	resBlock, err := app.FinalizeBlock(&types2.FinalizeBlockRequest{
		Height: header.Height,
		Txs:    [][]byte{bz},
	})
	require.NoError(t, err)

	require.Equal(t, 1, len(resBlock.TxResults))
	txResult := resBlock.TxResults[0]
	finalizeSuccess := txResult.Code == 0
	if expPass {
		require.True(t, finalizeSuccess)
	} else {
		require.False(t, finalizeSuccess)
	}

	_, err = app.Commit()
	require.NoError(t, err)
	gInfo := sdk.GasInfo{GasWanted: uint64(txResult.GasWanted), GasUsed: uint64(txResult.GasUsed)}
	txRes := sdk.Result{Data: txResult.Data, Log: txResult.Log, Events: txResult.Events}
	if finalizeSuccess {
		err = nil
	} else {
		err = errors.ABCIError(txResult.Codespace, txResult.Code, txResult.Log)
	}

	return gInfo, &txRes, err
}
