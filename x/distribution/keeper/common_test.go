package keeper_test

import (
	"cosmossdk.io/x/distribution/types"

	simtestutil "github.com/depinnetwork/depin-sdk/testutil/sims"
	sdk "github.com/depinnetwork/depin-sdk/types"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

var (
	PKS = simtestutil.CreateTestPubKeys(5)

	valConsPk0 = PKS[0]
	valConsPk1 = PKS[1]
	valConsPk2 = PKS[2]

	valConsAddr0 = sdk.ConsAddress(valConsPk0.Address())
	valConsAddr1 = sdk.ConsAddress(valConsPk1.Address())
	valConsAddr2 = sdk.ConsAddress(valConsPk2.Address())

	distrAcc = authtypes.NewEmptyModuleAccount(types.ModuleName)
)
