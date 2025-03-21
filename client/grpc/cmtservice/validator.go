package cmtservice

import (
	"context"

	coretypes "github.com/depinnetwork/por-consensus/rpc/core/types"
)

func getValidators(ctx context.Context, rpc CometRPC, height *int64, page, limit int) (*coretypes.ResultValidators, error) {
	return rpc.Validators(ctx, height, &page, &limit)
}
