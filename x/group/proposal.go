package group

import (
	gogoprotoany "github.com/cosmos/gogoproto/types/any"

	sdk "github.com/depinnetwork/depin-sdk/types"
	"github.com/depinnetwork/depin-sdk/types/tx"
)

// GetMsgs unpacks p.Messages Any's into sdk.Msg's
func (p *Proposal) GetMsgs() ([]sdk.Msg, error) {
	return tx.GetMsgs(p.Messages, "proposal")
}

// SetMsgs packs msgs into Any's
func (p *Proposal) SetMsgs(msgs []sdk.Msg) error {
	anys, err := tx.SetMsgs(msgs)
	if err != nil {
		return err
	}
	p.Messages = anys
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	return tx.UnpackInterfaces(unpacker, p.Messages)
}
