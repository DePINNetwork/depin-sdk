package cmtservice

import (
	"context"

	coretypes "github.com/depinnetwork/por-consensus/rpc/core/types"
)

// GetNodeStatus returns the status of the node.
func GetNodeStatus(ctx context.Context, rpc CometRPC) (*coretypes.ResultStatus, error) {
	return rpc.Status(ctx)
}
