package staking

import (
	"bytes"
	"sort"
	"testing"

	abci "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	cmtproto "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	protoio "github.com/cosmos/gogoproto/io"
	"github.com/cosmos/gogoproto/proto"
	gogotypes "github.com/cosmos/gogoproto/types"
	"gotest.tools/v3/assert"

	"cosmossdk.io/core/comet"
	"cosmossdk.io/core/header"
	"cosmossdk.io/math"
	"cosmossdk.io/x/staking/testutil"
	stakingtypes "cosmossdk.io/x/staking/types"

	"github.com/depinnetwork/depin-sdk/baseapp"
	"github.com/depinnetwork/depin-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
	"github.com/depinnetwork/depin-sdk/tests/integration/v2"
	simtestutil "github.com/depinnetwork/depin-sdk/testutil/sims"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

const chainID = "chain-id-123"

// TestValidateVoteExtensions is a unit test function that tests the validation of vote extensions.
// It sets up the necessary fixtures and validators, generates vote extensions for each validator,
// and validates the vote extensions using the baseapp.ValidateVoteExtensions function.
func TestValidateVoteExtensions(t *testing.T) {
	t.Parallel()
	f := initFixture(t, true)

	// enable vote extensions
	cp := simtestutil.DefaultConsensusParams
	cp.Feature = &cmtproto.FeatureParams{VoteExtensionsEnableHeight: &gogotypes.Int64Value{Value: 1}}

	assert.NilError(t, f.consensusKeeper.ParamsStore.Set(f.ctx, *cp))

	f.ctx = integration.SetHeaderInfo(f.ctx, header.Info{Height: 2, ChainID: chainID})

	// setup the validators
	numVals := 1
	privKeys := []cryptotypes.PrivKey{}
	for i := 0; i < numVals; i++ {
		privKeys = append(privKeys, ed25519.GenPrivKey())
	}

	vals := []stakingtypes.Validator{}
	for _, v := range privKeys {
		valAddr := sdk.ValAddress(v.PubKey().Address())
		acc := f.accountKeeper.NewAccountWithAddress(f.ctx, sdk.AccAddress(v.PubKey().Address()))
		f.accountKeeper.SetAccount(f.ctx, acc)
		simtestutil.AddTestAddrsFromPubKeys(f.bankKeeper, f.stakingKeeper, f.ctx, []cryptotypes.PubKey{v.PubKey()}, math.NewInt(100000000000))
		vals = append(vals, testutil.NewValidator(t, valAddr, v.PubKey()))
	}

	votes := []abci.ExtendedVoteInfo{}

	for i, v := range vals {
		v.Tokens = math.NewInt(1000000)
		v.Status = stakingtypes.Bonded
		assert.NilError(t, f.stakingKeeper.SetValidator(f.ctx, v))
		assert.NilError(t, f.stakingKeeper.SetValidatorByConsAddr(f.ctx, v))
		assert.NilError(t, f.stakingKeeper.SetNewValidatorByPowerIndex(f.ctx, v))
		_, err := f.stakingKeeper.Delegate(f.ctx, sdk.AccAddress(privKeys[i].PubKey().Address()), v.Tokens, stakingtypes.Unbonded, v, true)
		assert.NilError(t, err)

		// each val produces a vote
		voteExt := []byte("something" + v.OperatorAddress)
		cve := cmtproto.CanonicalVoteExtension{
			Extension: voteExt,
			Height:    integration.HeaderInfoFromContext(f.ctx).Height - 1, // the vote extension was signed in the previous height
			Round:     0,
			ChainId:   chainID,
		}

		extSignBytes, err := mashalVoteExt(&cve)
		assert.NilError(t, err)

		sig, err := privKeys[i].Sign(extSignBytes)
		assert.NilError(t, err)

		valbz, err := f.stakingKeeper.ValidatorAddressCodec().StringToBytes(v.GetOperator())
		assert.NilError(t, err)
		ve := abci.ExtendedVoteInfo{
			Validator: abci.Validator{
				Address: valbz,
				Power:   1000,
			},
			VoteExtension:      voteExt,
			ExtensionSignature: sig,
			BlockIDFlag:        cmtproto.BlockIDFlagCommit,
		}
		votes = append(votes, ve)
	}

	eci, ci := extendedCommitToLastCommit(abci.ExtendedCommitInfo{Round: 0, Votes: votes})
	f.ctx = integration.SetCometInfo(f.ctx, ci)

	err := baseapp.ValidateVoteExtensionsWithParams(f.ctx, *cp,
		integration.HeaderInfoFromContext(f.ctx), ci, f.stakingKeeper, eci)
	assert.NilError(t, err)
}

func mashalVoteExt(msg proto.Message) ([]byte, error) {
	var buf bytes.Buffer
	if err := protoio.NewDelimitedWriter(&buf).WriteMsg(msg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func extendedCommitToLastCommit(ec abci.ExtendedCommitInfo) (abci.ExtendedCommitInfo, comet.Info) {
	// sort the extended commit info
	sort.Sort(extendedVoteInfos(ec.Votes))

	// convert the extended commit info to last commit info
	lastCommit := comet.CommitInfo{
		Round: ec.Round,
		Votes: make([]comet.VoteInfo, len(ec.Votes)),
	}

	for i, vote := range ec.Votes {
		lastCommit.Votes[i] = comet.VoteInfo{
			Validator: comet.Validator{
				Address: vote.Validator.Address,
				Power:   vote.Validator.Power,
			},
		}
	}

	return ec, comet.Info{
		LastCommit: lastCommit,
	}
}

type extendedVoteInfos []abci.ExtendedVoteInfo

func (v extendedVoteInfos) Len() int {
	return len(v)
}

func (v extendedVoteInfos) Less(i, j int) bool {
	if v[i].Validator.Power == v[j].Validator.Power {
		return bytes.Compare(v[i].Validator.Address, v[j].Validator.Address) == -1
	}
	return v[i].Validator.Power > v[j].Validator.Power
}

func (v extendedVoteInfos) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
