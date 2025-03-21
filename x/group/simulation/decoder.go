package simulation

import (
	"bytes"
	"fmt"

	"cosmossdk.io/core/codec"
	"cosmossdk.io/x/group"
	"cosmossdk.io/x/group/keeper"

	"github.com/depinnetwork/depin-sdk/types/kv"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding group type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], []byte{keeper.GroupTablePrefix}):
			var groupA, groupB group.GroupInfo

			if err := cdc.Unmarshal(kvA.Value, &groupA); err != nil {
				panic(err)
			}
			if err := cdc.Unmarshal(kvB.Value, &groupB); err != nil {
				panic(err)
			}

			return fmt.Sprintf("%v\n%v", groupA, groupB)
		case bytes.Equal(kvA.Key[:1], []byte{keeper.GroupMemberTablePrefix}):
			var memberA, memberB group.GroupMember

			if err := cdc.Unmarshal(kvA.Value, &memberA); err != nil {
				panic(err)
			}
			if err := cdc.Unmarshal(kvB.Value, &memberB); err != nil {
				panic(err)
			}

			return fmt.Sprintf("%v\n%v", memberA, memberB)
		case bytes.Equal(kvA.Key[:1], []byte{keeper.GroupPolicyTablePrefix}):
			var accA, accB group.GroupPolicyInfo

			if err := cdc.Unmarshal(kvA.Value, &accA); err != nil {
				panic(err)
			}
			if err := cdc.Unmarshal(kvB.Value, &accB); err != nil {
				panic(err)
			}

			return fmt.Sprintf("%v\n%v", accA, accB)
		case bytes.Equal(kvA.Key[:1], []byte{keeper.ProposalTablePrefix}):
			var propA, propB group.Proposal

			if err := cdc.Unmarshal(kvA.Value, &propA); err != nil {
				panic(err)
			}
			if err := cdc.Unmarshal(kvB.Value, &propB); err != nil {
				panic(err)
			}

			return fmt.Sprintf("%v\n%v", propA, propB)
		case bytes.Equal(kvA.Key[:1], []byte{keeper.VoteTablePrefix}):
			var voteA, voteB group.Vote

			if err := cdc.Unmarshal(kvA.Value, &voteA); err != nil {
				panic(err)
			}
			if err := cdc.Unmarshal(kvB.Value, &voteB); err != nil {
				panic(err)
			}

			return fmt.Sprintf("%v\n%v", voteA, voteB)
		default:
			panic(fmt.Sprintf("invalid group key %X", kvA.Key))
		}
	}
}
