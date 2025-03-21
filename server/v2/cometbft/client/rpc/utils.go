package rpc

import (
	"fmt"

	cmttypes "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	coretypes "github.com/depinnetwork/por-consensus/rpc/core/types"
	gogoproto "github.com/cosmos/gogoproto/proto"
)

// formatBlockResults parses the indexed blocks into a slice of BlockResponse objects.
func formatBlockResults(resBlocks []*coretypes.ResultBlock) ([]*cmttypes.Block, error) {
	var (
		err error
		out = make([]*cmttypes.Block, len(resBlocks))
	)
	for i := range resBlocks {
		out[i], err = responseResultBlock(resBlocks[i])
		if err != nil {
			return nil, fmt.Errorf("unable to create response block from comet result block: %v: %w", resBlocks[i], err)
		}
		if out[i] == nil {
			return nil, fmt.Errorf("unable to create response block from comet result block: %v", resBlocks[i])
		}
	}

	return out, nil
}

// responseResultBlock returns a BlockResponse given a ResultBlock from CometBFT
func responseResultBlock(res *coretypes.ResultBlock) (*cmttypes.Block, error) {
	blkProto, err := res.Block.ToProto()
	if err != nil {
		return nil, err
	}
	blkBz, err := gogoproto.Marshal(blkProto)
	if err != nil {
		return nil, err
	}

	blk := &cmttypes.Block{}
	err = gogoproto.Unmarshal(blkBz, blk)
	if err != nil {
		return nil, err
	}
	return blk, nil
}

// calcTotalPages calculates total pages in an overflow safe manner
func calcTotalPages(totalCount, limit int64) int64 {
	totalPages := int64(0)
	if totalCount != 0 && limit != 0 {
		if totalCount%limit > 0 {
			totalPages = totalCount/limit + 1
		} else {
			totalPages = totalCount / limit
		}
	}
	return totalPages
}
