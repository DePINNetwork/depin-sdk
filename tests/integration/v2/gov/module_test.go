package gov

import (
	"testing"

	"gotest.tools/v3/assert"

	_ "cosmossdk.io/x/accounts"
	"cosmossdk.io/x/gov/types"
	_ "cosmossdk.io/x/mint"
	_ "cosmossdk.io/x/protocolpool"

	"github.com/depinnetwork/depin-sdk/tests/integration/v2"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	suite := createTestSuite(t, integration.Genesis_COMMIT)
	ctx := suite.ctx

	acc := suite.AuthKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	assert.Assert(t, acc != nil)
}
