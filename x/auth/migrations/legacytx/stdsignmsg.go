package legacytx

import (
	gogoprotoany "github.com/cosmos/gogoproto/types/any"

	"github.com/depinnetwork/depin-sdk/codec/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

var _ gogoprotoany.UnpackInterfacesMessage = StdSignMsg{}

// StdSignMsg is a convenience structure for passing along a Msg with the other
// requirements for a StdSignDoc before it is signed. For use in the CLI.
type StdSignMsg struct {
	ChainID       string    `json:"chain_id" yaml:"chain_id"`
	AccountNumber uint64    `json:"account_number" yaml:"account_number"`
	Sequence      uint64    `json:"sequence" yaml:"sequence"`
	TimeoutHeight uint64    `json:"timeout_height" yaml:"timeout_height"`
	Fee           StdFee    `json:"fee" yaml:"fee"`
	Msgs          []sdk.Msg `json:"msgs" yaml:"msgs"`
	Memo          string    `json:"memo" yaml:"memo"`
}

func (msg StdSignMsg) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	for _, m := range msg.Msgs {
		err := types.UnpackInterfaces(m, unpacker)
		if err != nil {
			return err
		}
	}

	return nil
}
