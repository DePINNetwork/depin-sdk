package authz

import (
	"cosmossdk.io/x/authz/client/cli"

	"github.com/depinnetwork/depin-sdk/client"
	"github.com/depinnetwork/depin-sdk/testutil"
	clitestutil "github.com/depinnetwork/depin-sdk/testutil/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	return clitestutil.ExecTestCLICmd(clientCtx, cli.NewCmdGrantAuthorization(), args)
}
