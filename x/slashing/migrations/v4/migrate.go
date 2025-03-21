package v4

import (
	"context"

	"github.com/bits-and-blooms/bitset"
	gogotypes "github.com/cosmos/gogoproto/types"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/codec"
	"cosmossdk.io/errors"
	storetypes "github.com/depinnetwork/depin-sdk/store/types"
	"cosmossdk.io/x/slashing/types"

	sdk "github.com/depinnetwork/depin-sdk/types"
)

// Migrate migrates state to consensus version 4. Specifically, the migration
// deletes all existing validator bitmap entries and replaces them with a real
// "chunked" bitmap.
func Migrate(ctx context.Context, cdc codec.BinaryCodec, store storetypes.KVStore, params types.Params, addressCodec address.ValidatorAddressCodec) error {
	// Get all the missed blocks for each validator, based on the existing signing
	// info.
	var missedBlocks []types.ValidatorMissedBlocks
	iterateValidatorSigningInfos(ctx, cdc, store, func(addr sdk.ConsAddress, info types.ValidatorSigningInfo) (stop bool) {
		bechAddr, err := addressCodec.BytesToString(addr)
		if err != nil {
			return true
		}
		localMissedBlocks := GetValidatorMissedBlocks(ctx, cdc, store, addr, params)

		missedBlocks = append(missedBlocks, types.ValidatorMissedBlocks{
			Address:      bechAddr,
			MissedBlocks: localMissedBlocks,
		})

		return false
	})

	// For each missed blocks entry, of which there should only be one per validator,
	// we clear all the old entries and insert the new chunked entry.
	for _, mb := range missedBlocks {
		addr, err := sdk.ConsAddressFromBech32(mb.Address)
		if err != nil {
			return err
		}

		deleteValidatorMissedBlockBitArray(ctx, store, addr)

		for _, b := range mb.MissedBlocks {
			// Note: It is not necessary to store entries with missed=false, i.e. where
			// the bit is zero, since when the bitmap is initialized, all non-set bits
			// are already zero.
			if b.Missed {
				if err := setMissedBlockBitmapValue(ctx, store, addr, b.Index, true); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func iterateValidatorSigningInfos(
	_ context.Context,
	cdc codec.BinaryCodec,
	store storetypes.KVStore,
	cb func(address sdk.ConsAddress, info types.ValidatorSigningInfo) (stop bool),
) {
	iter := storetypes.KVStorePrefixIterator(store, ValidatorSigningInfoKeyPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		addr := ValidatorSigningInfoAddress(iter.Key())
		var info types.ValidatorSigningInfo
		if err := cdc.Unmarshal(iter.Value(), &info); err != nil {
			panic(err)
		}

		if cb(addr, info) {
			break
		}
	}
}

func iterateValidatorMissedBlockBitArray(
	_ context.Context,
	cdc codec.BinaryCodec,
	store storetypes.KVStore,
	addr sdk.ConsAddress,
	params types.Params,
	cb func(index int64, missed bool) (stop bool),
) {
	for i := int64(0); i < params.SignedBlocksWindow; i++ {
		var missed gogotypes.BoolValue
		bz := store.Get(ValidatorMissedBlockBitArrayKey(addr, i))
		if bz == nil {
			continue
		}

		if err := cdc.Unmarshal(bz, &missed); err != nil {
			panic(err)
		}
		if cb(i, missed.Value) {
			break
		}
	}
}

func GetValidatorMissedBlocks(
	ctx context.Context,
	cdc codec.BinaryCodec,
	store storetypes.KVStore,
	addr sdk.ConsAddress,
	params types.Params,
) []types.MissedBlock {
	var missedBlocks []types.MissedBlock
	iterateValidatorMissedBlockBitArray(ctx, cdc, store, addr, params, func(index int64, missed bool) (stop bool) {
		missedBlocks = append(missedBlocks, types.NewMissedBlock(index, missed))
		return false
	})

	return missedBlocks
}

func deleteValidatorMissedBlockBitArray(_ context.Context, store storetypes.KVStore, addr sdk.ConsAddress) {
	iter := storetypes.KVStorePrefixIterator(store, validatorMissedBlockBitArrayPrefixKey(addr))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

func setMissedBlockBitmapValue(_ context.Context, store storetypes.KVStore, addr sdk.ConsAddress, index int64, missed bool) error {
	// get the chunk or "word" in the logical bitmap
	chunkIndex := index / MissedBlockBitmapChunkSize
	key := ValidatorMissedBlockBitmapKey(addr, chunkIndex)

	bs := bitset.New(uint(MissedBlockBitmapChunkSize))
	chunk := store.Get(key)
	if chunk != nil {
		if err := bs.UnmarshalBinary(chunk); err != nil {
			return errors.Wrapf(err, "failed to decode bitmap chunk; index: %d", index)
		}
	}

	// get the bit position in the chunk of the logical bitmap
	bitIndex := uint(index % MissedBlockBitmapChunkSize)
	if missed {
		bs.Set(bitIndex)
	} else {
		bs.Clear(bitIndex)
	}

	updatedChunk, err := bs.MarshalBinary()
	if err != nil {
		return errors.Wrapf(err, "failed to encode bitmap chunk; index: %d", index)
	}

	store.Set(key, updatedChunk)
	return nil
}
