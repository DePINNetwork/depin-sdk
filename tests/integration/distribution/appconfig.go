package distribution_test

import (
	_ "cosmossdk.io/x/accounts"     // import as blank for app wiring
	_ "cosmossdk.io/x/bank"         // import as blank for app wiring
	_ "cosmossdk.io/x/consensus"    // import as blank for app wiring
	_ "cosmossdk.io/x/distribution" // import as blank for app wiring
	_ "cosmossdk.io/x/mint"         // import as blank for app wiring
	_ "cosmossdk.io/x/protocolpool" // import as blank for app wiring
	_ "cosmossdk.io/x/staking"      // import as blank for app wiring

	"github.com/depinnetwork/depin-sdk/testutil/configurator"
	_ "github.com/depinnetwork/depin-sdk/x/auth"           // import as blank for app wiring
	_ "github.com/depinnetwork/depin-sdk/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/depinnetwork/depin-sdk/x/genutil"        // import as blank for app wiring
)

var AppConfig = configurator.NewAppConfig(
	configurator.AccountsModule(),
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.StakingModule(),
	configurator.TxModule(),
	configurator.ValidateModule(),
	configurator.ConsensusModule(),
	configurator.GenutilModule(),
	configurator.DistributionModule(),
	configurator.MintModule(),
	configurator.ProtocolPoolModule(),
)
