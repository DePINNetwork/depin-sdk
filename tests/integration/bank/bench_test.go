package bank_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	abci "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	"github.com/stretchr/testify/require"

	_ "cosmossdk.io/x/bank"
	"cosmossdk.io/x/bank/testutil"
	"cosmossdk.io/x/bank/types"
	stakingtypes "cosmossdk.io/x/staking/types"

	"github.com/depinnetwork/depin-sdk/client"
	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
	simtestutil "github.com/depinnetwork/depin-sdk/testutil/sims"
	sdk "github.com/depinnetwork/depin-sdk/types"
	moduletestutil "github.com/depinnetwork/depin-sdk/types/module/testutil"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

var moduleAccAddr = authtypes.NewModuleAddress(stakingtypes.BondedPoolName)

// GenSequenceOfTxs generates a set of signed transactions of messages, such
// that they differ only by having the sequence numbers incremented between
// every transaction.
func genSequenceOfTxs(txGen client.TxConfig,
	msgs []sdk.Msg,
	accNums []uint64,
	initSeqNums []uint64,
	numToGenerate int,
	priv ...cryptotypes.PrivKey,
) ([]sdk.Tx, error) {
	var err error

	txs := make([]sdk.Tx, numToGenerate)
	for i := 0; i < numToGenerate; i++ {
		txs[i], err = simtestutil.GenSignedMockTx(
			rand.New(rand.NewSource(time.Now().UnixNano())),
			txGen,
			msgs,
			sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)},
			simtestutil.DefaultGenTxGas,
			"",
			accNums,
			initSeqNums,
			priv...,
		)
		if err != nil {
			break
		}

		for i := 0; i < len(initSeqNums); i++ {
			initSeqNums[i]++
		}
	}

	return txs, err
}

func BenchmarkOneBankSendTxPerBlock(b *testing.B) {
	// b.Skip("Skipping benchmark with buggy code reported at https://github.com/cosmos/cosmos-sdk/issues/10023")
	b.ReportAllocs()

	acc := authtypes.BaseAccount{
		Address: addr1.String(),
	}

	// construct genesis state
	genAccs := []authtypes.GenesisAccount{&acc}
	s := createTestSuite(&testing.T{}, genAccs)
	baseApp := s.App.BaseApp
	ctx := baseApp.NewContext(false)

	_, err := baseApp.FinalizeBlock(&abci.FinalizeBlockRequest{Height: 1})
	require.NoError(b, err)

	require.NoError(b, testutil.FundAccount(ctx, s.BankKeeper, addr1, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 100000000000))))

	_, err = baseApp.Commit()
	require.NoError(b, err)

	txGen := moduletestutil.MakeTestTxConfig(codectestutil.CodecOptions{})
	txEncoder := txGen.TxEncoder()

	// pre-compute all txs
	txs, err := genSequenceOfTxs(txGen, []sdk.Msg{sendMsg1}, []uint64{0}, []uint64{uint64(0)}, b.N, priv1)
	require.NoError(b, err)
	b.ResetTimer()

	height := int64(2)

	// Run this with a profiler, so its easy to distinguish what time comes from
	// Committing, and what time comes from Check/Deliver Tx.
	for i := 0; i < b.N; i++ {
		_, _, err := baseApp.SimCheck(txEncoder, txs[i])
		if err != nil {
			panic(fmt.Errorf("failed to simulate tx: %w", err))
		}

		bz, err := txEncoder(txs[i])
		require.NoError(b, err)

		_, err = baseApp.FinalizeBlock(
			&abci.FinalizeBlockRequest{
				Height: height,
				Txs:    [][]byte{bz},
			},
		)
		require.NoError(b, err)

		_, err = baseApp.Commit()
		require.NoError(b, err)

		height++
	}
}

func BenchmarkOneBankMultiSendTxPerBlock(b *testing.B) {
	// b.Skip("Skipping benchmark with buggy code reported at https://github.com/cosmos/cosmos-sdk/issues/10023")
	b.ReportAllocs()

	addr1Str, err := codectestutil.CodecOptions{}.GetAddressCodec().BytesToString(addr1)
	require.NoError(b, err)
	acc := authtypes.BaseAccount{
		Address: addr1.String(),
	}

	addr2Str, err := codectestutil.CodecOptions{}.GetAddressCodec().BytesToString(addr2)
	require.NoError(b, err)

	// construct genesis state
	genAccs := []authtypes.GenesisAccount{&acc}
	s := createTestSuite(&testing.T{}, genAccs)
	baseApp := s.App.BaseApp
	ctx := baseApp.NewContext(false)

	_, err = baseApp.FinalizeBlock(&abci.FinalizeBlockRequest{Height: 1})
	require.NoError(b, err)

	require.NoError(b, testutil.FundAccount(ctx, s.BankKeeper, addr1, sdk.NewCoins(sdk.NewInt64Coin("foocoin", 100000000000))))

	_, err = baseApp.Commit()
	require.NoError(b, err)

	txGen := moduletestutil.MakeTestTxConfig(codectestutil.CodecOptions{})
	txEncoder := txGen.TxEncoder()

	multiSendMsg := &types.MsgMultiSend{
		Inputs:  []types.Input{types.NewInput(addr1Str, coins)},
		Outputs: []types.Output{types.NewOutput(addr2Str, coins)},
	}

	// pre-compute all txs
	txs, err := genSequenceOfTxs(txGen, []sdk.Msg{multiSendMsg}, []uint64{0}, []uint64{uint64(0)}, b.N, priv1)
	require.NoError(b, err)
	b.ResetTimer()

	height := int64(2)

	// Run this with a profiler, so its easy to distinguish what time comes from
	// Committing, and what time comes from Check/Deliver Tx.
	for i := 0; i < b.N; i++ {
		_, _, err := baseApp.SimCheck(txEncoder, txs[i])
		if err != nil {
			panic(fmt.Errorf("failed to simulate tx: %w", err))
		}

		bz, err := txEncoder(txs[i])
		require.NoError(b, err)

		_, err = baseApp.FinalizeBlock(
			&abci.FinalizeBlockRequest{
				Height: height,
				Txs:    [][]byte{bz},
			},
		)
		require.NoError(b, err)

		_, err = baseApp.Commit()
		require.NoError(b, err)

		height++
	}
}
