package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/x/nft"
	nftmodule "cosmossdk.io/x/nft/module"
	"cosmossdk.io/x/nft/simulation"

	addresscodec "github.com/depinnetwork/depin-sdk/codec/address"
	codectestutil "github.com/depinnetwork/depin-sdk/codec/testutil"
	"github.com/depinnetwork/depin-sdk/types/module"
	moduletestutil "github.com/depinnetwork/depin-sdk/types/module/testutil"
	simtypes "github.com/depinnetwork/depin-sdk/types/simulation"
)

func TestRandomizedGenState(t *testing.T) {
	cdcOpts := codectestutil.CodecOptions{}
	encCfg := moduletestutil.MakeTestEncodingConfig(cdcOpts, nftmodule.AppModule{})

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:      make(simtypes.AppParams),
		Cdc:            encCfg.Codec,
		AddressCodec:   cdcOpts.GetAddressCodec(),
		ValidatorCodec: cdcOpts.GetValidatorCodec(),
		Rand:           r,
		NumBonded:      3,
		Accounts:       simtypes.RandomAccounts(r, 3),
		InitialStake:   sdkmath.NewInt(1000),
		GenState:       make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState, addresscodec.NewBech32Codec("cosmos"))
	var nftGenesis nft.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[nft.ModuleName], &nftGenesis)

	require.Len(t, nftGenesis.Classes, len(simState.Accounts)-1)
	require.Len(t, nftGenesis.Entries, len(simState.Accounts)-1)
}
