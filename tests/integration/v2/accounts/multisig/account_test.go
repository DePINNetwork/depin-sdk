package multisig

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	basev1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	"cosmossdk.io/core/header"
	"cosmossdk.io/math"
	v1 "cosmossdk.io/x/accounts/defaults/multisig/v1"
	accountsv1 "cosmossdk.io/x/accounts/v1"

	codectypes "github.com/depinnetwork/depin-sdk/codec/types"
	"github.com/depinnetwork/depin-sdk/crypto/keys/secp256k1"
	"github.com/depinnetwork/depin-sdk/tests/integration/v2"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, NewIntegrationTestSuite())
}

// TestSimpleSendProposal creates a multisig account with 1 member, sends a tx, votes and executes it.
func (s *IntegrationTestSuite) TestSimpleSendProposal() {
	ctx := s.app.StateLatestContext(s.T())

	ctx = integration.SetHeaderInfo(ctx, header.Info{Time: time.Now()})

	randAcc := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	addr, err := s.authKeeper.AddressCodec().BytesToString(randAcc)
	s.NoError(err)

	initialMembers := map[string]uint64{
		s.membersAddr[0]: 100,
	}
	accountAddr, accAddrStr := s.initAccount(ctx, s.members[0], initialMembers)

	balance := s.bankKeeper.GetBalance(ctx, randAcc, "stake")
	s.Equal(math.NewInt(0), balance.Amount)

	// do a simple bank send
	msg := &bankv1beta1.MsgSend{
		FromAddress: accAddrStr,
		ToAddress:   addr,
		Amount: []*basev1beta1.Coin{
			{
				Denom:  "stake",
				Amount: "100",
			},
		},
	}
	anyMsg, err := codectypes.NewAnyWithValue(msg)
	s.NoError(err)

	s.createProposal(ctx, accountAddr, s.members[0], anyMsg)

	// now we vote for it
	voteReq := &v1.MsgVote{
		ProposalId: 0,
		Vote:       v1.VoteOption_VOTE_OPTION_YES,
	}

	err = s.executeTx(ctx, voteReq, accountAddr, s.members[0])
	s.NoError(err)

	// now we execute it
	err = s.executeProposal(ctx, accountAddr, s.members[0], 0)
	s.NoError(err)

	foundPropResult := false
	events := integration.EventsFromContext(ctx)
	for _, e := range events {
		if e.Type == "proposal_tally" {
			foundPropResult = true
			attr, found := integration.GetAttribute(e, "status")
			s.True(found)
			s.Equal(v1.ProposalStatus_PROPOSAL_STATUS_PASSED.String(), attr.Value)

			yesVotes, found := integration.GetAttribute(e, "yes_votes")
			s.True(found)
			s.Equal("100", yesVotes.Value)

			noVotes, found := integration.GetAttribute(e, "no_votes")
			s.True(found)
			s.Equal("0", noVotes.Value)

			propID, found := integration.GetAttribute(e, "proposal_id")
			s.True(found)
			s.Equal("0", propID.Value)

			execErr, found := integration.GetAttribute(e, "exec_err")
			s.True(found)
			s.Equal("<nil>", execErr.Value)

			rejectErr, found := integration.GetAttribute(e, "reject_err")
			s.True(found)
			s.Equal("<nil>", rejectErr.Value)

		}
	}
	s.True(foundPropResult)

	balance = s.bankKeeper.GetBalance(ctx, randAcc, "stake")
	s.Equal(int64(100), balance.Amount.Int64())

	// try to execute again, should fail
	err = s.executeProposal(ctx, accountAddr, s.members[0], 0)
	s.Error(err)
}

// TestConfigUpdate creates a multisig with 1 member, adds 2 more members and
// changes the config to require 2/3 majority (also through a proposal).
func (s *IntegrationTestSuite) TestConfigUpdate() {
	ctx := s.app.StateLatestContext(s.T())

	ctx = integration.SetHeaderInfo(ctx, header.Info{Time: time.Now()})

	initialMembers := map[string]uint64{
		s.membersAddr[0]: 100,
	}
	accountAddr, accAddrStr := s.initAccount(ctx, s.members[0], initialMembers)

	// Add 2 members and pass the proposal
	// create proposal
	updateMsg := &v1.MsgUpdateConfig{
		UpdateMembers: []*v1.Member{
			{
				Address: s.membersAddr[1],
				Weight:  100,
			},
			{
				Address: s.membersAddr[2],
				Weight:  100,
			},
		},
		Config: &v1.Config{
			Threshold:      200, // 3 members with 100 power each, 2/3 majority
			Quorum:         200,
			VotingPeriod:   120,
			Revote:         false,
			EarlyExecution: false,
		},
	}

	msgExec := &accountsv1.MsgExecute{
		Sender:  accAddrStr,
		Target:  accAddrStr,
		Message: codectypes.UnsafePackAny(updateMsg),
		Funds:   []sdk.Coin{},
	}

	s.createProposal(ctx, accountAddr, s.members[0], codectypes.UnsafePackAny(msgExec))

	// vote
	voteReq := &v1.MsgVote{
		ProposalId: 0,
		Vote:       v1.VoteOption_VOTE_OPTION_YES,
	}

	err := s.executeTx(ctx, voteReq, accountAddr, s.members[0])
	s.NoError(err)

	err = s.executeProposal(ctx, accountAddr, s.members[0], 0)
	s.NoError(err)

	// get members
	res, err := s.queryAcc(ctx, &v1.QueryConfig{}, accountAddr)
	s.NoError(err)
	resp := res.(*v1.QueryConfigResponse)
	s.Len(resp.Members, 3)
	s.Equal(int64(200), resp.Config.Threshold)

	// Try to remove a member, but it doesn't reach passing threshold
	// create proposal
	msgExec = &accountsv1.MsgExecute{
		Sender: accAddrStr,
		Target: accAddrStr,
		Message: codectypes.UnsafePackAny(&v1.MsgUpdateConfig{
			UpdateMembers: []*v1.Member{
				{
					Address: s.membersAddr[1],
					Weight:  0,
				},
			},
			Config: &v1.Config{
				Threshold:      200, // 3 members with 100 power each, 2/3 majority
				Quorum:         200,
				VotingPeriod:   120,
				Revote:         false,
				EarlyExecution: false,
			},
		}),
		Funds: []sdk.Coin{},
	}

	s.createProposal(ctx, accountAddr, s.members[0], codectypes.UnsafePackAny(msgExec))

	// vote
	voteReq = &v1.MsgVote{
		ProposalId: 1,
		Vote:       v1.VoteOption_VOTE_OPTION_NO,
	}

	err = s.executeTx(ctx, voteReq, accountAddr, s.members[0])
	s.NoError(err)

	// need to wait until voting period is over because we disabled early execution on the last
	// config update
	err = s.executeProposal(ctx, accountAddr, s.members[0], 0)
	s.ErrorContains(err, "voting period has not ended yet, and early execution is not enabled")

	// vote with member 1
	voteReq = &v1.MsgVote{
		ProposalId: 1,
		Vote:       v1.VoteOption_VOTE_OPTION_NO,
	}

	err = s.executeTx(ctx, voteReq, accountAddr, s.members[1])
	s.NoError(err)

	// need to wait until voting period is over because we disabled early execution on the last
	// config update
	err = s.executeProposal(ctx, accountAddr, s.members[0], 1)
	s.ErrorContains(err, "voting period has not ended yet, and early execution is not enabled")

	headerInfo := integration.HeaderInfoFromContext(ctx)
	headerInfo.Time = headerInfo.Time.Add(time.Second * 121)
	ctx = integration.SetHeaderInfo(ctx, headerInfo)

	// now it should work, but the proposal will fail
	err = s.executeProposal(ctx, accountAddr, s.members[0], 1)
	s.NoError(err)

	foundPropResult := false
	events := integration.EventsFromContext(ctx)
	for _, e := range events {
		if e.Type == "proposal_tally" {
			propID, found := integration.GetAttribute(e, "proposal_id")
			s.True(found)

			if propID.Value == "1" {
				foundPropResult = true
				status, found := integration.GetAttribute(e, "status")
				s.True(found)
				s.Equal(v1.ProposalStatus_PROPOSAL_STATUS_REJECTED.String(), status.Value)

				// exec_err is nil because the proposal didn't execute
				execErr, found := integration.GetAttribute(e, "exec_err")
				s.True(found)
				s.Equal("<nil>", execErr.Value)

				rejectErr, found := integration.GetAttribute(e, "reject_err")
				s.True(found)
				s.Equal("threshold not reached", rejectErr.Value)
			}

		}
	}
	s.True(foundPropResult)

	// get members
	res, err = s.queryAcc(ctx, &v1.QueryConfig{}, accountAddr)
	s.NoError(err)
	resp = res.(*v1.QueryConfigResponse)
	s.Len(resp.Members, 3)
	s.Equal(int64(200), resp.Config.Threshold)
}
