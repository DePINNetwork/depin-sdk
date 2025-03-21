package mock

import (
	"github.com/depinnetwork/depin-sdk/types/adapters"
	cmtproto "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	"github.com/depinnetwork/por-consensus/crypto"
	cmttypes "github.com/depinnetwork/por-consensus/types"

	cryptocodec "github.com/depinnetwork/depin-sdk/crypto/codec"
	"github.com/depinnetwork/depin-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
)

var _ cmttypes.PrivValidator = PV{}

// PV implements PrivValidator without any safety or persistence.
// Only use it for testing.
type PV struct {
	PrivKey cryptotypes.PrivKey
}

func NewPV() PV {
	return PV{ed25519.GenPrivKey()}
}

// GetPubKey implements PrivValidator interface
func (pv PV) GetPubKey() (crypto.PubKey, error) {
	return cryptocodec.ToCmtPubKeyInterface(pv.PrivKey.PubKey())
}

// SignVote implements PrivValidator interface
func (pv PV) SignVote(chainID string, vote *cmtproto.Vote, _ bool) error {
	signBytes := cmttypes.VoteSignBytes(chainID, vote)
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	vote.Signature = sig
	return nil
}

// SignProposal implements PrivValidator interface
func (pv PV) SignProposal(chainID string, proposal *cmtproto.Proposal) error {
	signBytes := cmttypes.ProposalSignBytes(chainID, proposal)
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	proposal.Signature = sig
	return nil
}

func (pv PV) SignBytes(bytes []byte) ([]byte, error) {
	return pv.PrivKey.Sign(bytes)
}
