package v2

import (
	"encoding/binary"

	"cosmossdk.io/core/address"

	sdk "github.com/depinnetwork/depin-sdk/types"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

// MigrateGenState accepts exported v0.46 x/auth genesis state and migrates it to
// v0.47 x/auth genesis state. The migration includes:
// - If the group module is enabled, replace group policy accounts from module accounts to base accounts.
func MigrateGenState(oldState *authtypes.GenesisState, addressCodec address.Codec) *authtypes.GenesisState {
	newState := *oldState

	accounts, err := authtypes.UnpackAccounts(newState.Accounts)
	if err != nil {
		panic(err)
	}

	groupPolicyAccountCounter := uint64(0)
	for i, acc := range accounts {
		modAcc, ok := acc.(sdk.ModuleAccountI)
		if !ok {
			continue
		}

		modAddr, err := addressCodec.BytesToString(modAcc.GetAddress())
		if err != nil {
			panic(err)
		}

		if modAcc.GetName() != modAddr {
			continue
		}

		// Replace group policy accounts from module accounts to base accounts.
		// These accounts were wrongly created and the address was equal to the module name.
		derivationKey := make([]byte, 8)
		binary.BigEndian.PutUint64(derivationKey, groupPolicyAccountCounter)

		cred, err := authtypes.NewModuleCredential(ModuleName, []byte{GroupPolicyTablePrefix}, derivationKey)
		if err != nil {
			panic(err)
		}
		baseAccount, err := authtypes.NewBaseAccountWithPubKey(cred)
		if err != nil {
			panic(err)
		}

		if err := baseAccount.SetAccountNumber(modAcc.GetAccountNumber()); err != nil {
			panic(err)
		}
		accounts[i] = baseAccount
		groupPolicyAccountCounter++
	}

	packedAccounts, err := authtypes.PackAccounts(accounts)
	if err != nil {
		panic(err)
	}
	newState.Accounts = packedAccounts

	return &newState
}
