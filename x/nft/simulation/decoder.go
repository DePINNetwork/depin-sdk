package simulation

import (
	"bytes"
	"fmt"

	"cosmossdk.io/core/codec"
	"cosmossdk.io/x/nft"
	"cosmossdk.io/x/nft/keeper"

	sdk "github.com/depinnetwork/depin-sdk/types"
	"github.com/depinnetwork/depin-sdk/types/kv"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding nft type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], keeper.ClassKey):
			var classA, classB nft.Class
			if err := cdc.Unmarshal(kvA.Value, &classA); err != nil {
				panic(err)
			}
			if err := cdc.Unmarshal(kvB.Value, &classB); err != nil {
				panic(err)
			}
			return fmt.Sprintf("%v\n%v", classA, classB)
		case bytes.Equal(kvA.Key[:1], keeper.NFTKey):
			var nftA, nftB nft.NFT
			if err := cdc.Unmarshal(kvA.Value, &nftA); err != nil {
				panic(err)
			}
			if err := cdc.Unmarshal(kvB.Value, &nftB); err != nil {
				panic(err)
			}
			return fmt.Sprintf("%v\n%v", nftA, nftB)
		case bytes.Equal(kvA.Key[:1], keeper.NFTOfClassByOwnerKey):
			return fmt.Sprintf("%v\n%v", kvA.Value, kvB.Value)
		case bytes.Equal(kvA.Key[:1], keeper.OwnerKey):
			var ownerA, ownerB sdk.AccAddress
			ownerA = sdk.AccAddress(kvA.Value)
			ownerB = sdk.AccAddress(kvB.Value)
			return fmt.Sprintf("%v\n%v", ownerA, ownerB)
		case bytes.Equal(kvA.Key[:1], keeper.ClassTotalSupply):
			var supplyA, supplyB uint64
			supplyA = sdk.BigEndianToUint64(kvA.Value)
			supplyB = sdk.BigEndianToUint64(kvB.Value)
			return fmt.Sprintf("%v\n%v", supplyA, supplyB)
		default:
			panic(fmt.Sprintf("invalid nft key %X", kvA.Key))
		}
	}
}
