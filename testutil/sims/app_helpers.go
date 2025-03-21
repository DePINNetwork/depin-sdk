package sims

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	abci "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	cmtproto "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	cmtjson "github.com/depinnetwork/por-consensus/libs/json"
	cmttypes "github.com/depinnetwork/por-consensus/types"

	"cosmossdk.io/core/server"
	corestore "cosmossdk.io/core/store"
	coretesting "cosmossdk.io/core/testing"
	"cosmossdk.io/depinject"
	sdkmath "cosmossdk.io/math"
	banktypes "cosmossdk.io/x/bank/types"
	stakingtypes "cosmossdk.io/x/staking/types"

	"github.com/depinnetwork/depin-sdk/codec"
	codectypes "github.com/depinnetwork/depin-sdk/codec/types"
	cryptocodec "github.com/depinnetwork/depin-sdk/crypto/codec"
	"github.com/depinnetwork/depin-sdk/crypto/keys/secp256k1"
	"github.com/depinnetwork/depin-sdk/runtime"
	"github.com/depinnetwork/depin-sdk/testutil/mock"
	sdk "github.com/depinnetwork/depin-sdk/types"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

const DefaultGenTxGas = 10000000

// DefaultConsensusParams defines the default CometBFT consensus params used in
// SimApp testing.
var DefaultConsensusParams = &cmtproto.ConsensusParams{
	Version: &cmtproto.VersionParams{
		App: 1,
	},
	Block: &cmtproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   100_000_000,
	},
	Evidence: &cmtproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &cmtproto.ValidatorParams{
		PubKeyTypes: []string{
			cmttypes.ABCIPubKeyTypeEd25519,
			cmttypes.ABCIPubKeyTypeSecp256k1,
		},
	},
}

// CreateRandomValidatorSet creates a validator set with one random validator
func CreateRandomValidatorSet() (*cmttypes.ValidatorSet, error) {
	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get pub key: %w", err)
	}

	// create validator set with single validator
	validator := cmttypes.NewValidator(pubKey, 1)

	return cmttypes.NewValidatorSet([]*cmttypes.Validator{validator}), nil
}

type GenesisAccount struct {
	authtypes.GenesisAccount
	Coins sdk.Coins
}

// StartupConfig defines the startup configuration new a test application.
//
// ValidatorSet defines a custom validator set to be validating the app.
// BaseAppOption defines the additional operations that must be run on baseapp before app start.
// AtGenesis defines if the app started should already have produced block or not.
type StartupConfig struct {
	ValidatorSet    func() (*cmttypes.ValidatorSet, error)
	BaseAppOption   runtime.BaseAppOption
	AtGenesis       bool
	GenesisAccounts []GenesisAccount
	DB              corestore.KVStoreWithBatch
}

func DefaultStartUpConfig() StartupConfig {
	priv := secp256k1.GenPrivKey()
	ba := authtypes.NewBaseAccount(priv.PubKey().Address().Bytes(), priv.PubKey(), 0, 0)
	ga := GenesisAccount{ba, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(100000000000000)))}
	return StartupConfig{
		ValidatorSet:    CreateRandomValidatorSet,
		AtGenesis:       false,
		GenesisAccounts: []GenesisAccount{ga},
		DB:              coretesting.NewMemDB(),
	}
}

// Setup initializes a new runtime.App and can inject values into extraOutputs.
// It uses SetupWithConfiguration under the hood.
func Setup(appConfig depinject.Config, extraOutputs ...interface{}) (*runtime.App, error) {
	return SetupWithConfiguration(appConfig, DefaultStartUpConfig(), extraOutputs...)
}

// SetupAtGenesis initializes a new runtime.App at genesis and can inject values into extraOutputs.
// It uses SetupWithConfiguration under the hood.
func SetupAtGenesis(appConfig depinject.Config, extraOutputs ...interface{}) (*runtime.App, error) {
	cfg := DefaultStartUpConfig()
	cfg.AtGenesis = true
	return SetupWithConfiguration(appConfig, cfg, extraOutputs...)
}

// NextBlock starts a new block.
func NextBlock(app *runtime.App, ctx sdk.Context, jumpTime time.Duration) (sdk.Context, error) {
	_, err := app.FinalizeBlock(&abci.FinalizeBlockRequest{Height: ctx.BlockHeight(), Time: ctx.BlockTime()})
	if err != nil {
		return sdk.Context{}, err
	}
	_, err = app.Commit()
	if err != nil {
		return sdk.Context{}, err
	}

	newBlockTime := ctx.BlockTime().Add(jumpTime)

	header := ctx.BlockHeader()
	header.Time = newBlockTime
	header.Height++

	newCtx := app.BaseApp.NewUncachedContext(false, header)

	return newCtx, nil
}

// SetupWithConfiguration initializes a new runtime.App. A Nop logger is set in runtime.App.
// appConfig defines the application configuration (f.e. app_config.go).
// extraOutputs defines the extra outputs to be assigned by the dependency injector (depinject).
func SetupWithConfiguration(appConfig depinject.Config, startupConfig StartupConfig, extraOutputs ...interface{}) (*runtime.App, error) {
	// create the app with depinject
	var (
		app        *runtime.App
		appBuilder *runtime.AppBuilder
		codec      codec.Codec
	)

	if err := depinject.Inject(appConfig, append(extraOutputs, &appBuilder, &codec)...); err != nil {
		return nil, fmt.Errorf("failed to inject dependencies: %w", err)
	}

	if startupConfig.BaseAppOption != nil {
		app = appBuilder.Build(startupConfig.DB, nil, startupConfig.BaseAppOption)
	} else {
		app = appBuilder.Build(startupConfig.DB, nil)
	}
	if err := app.Load(true); err != nil {
		return nil, fmt.Errorf("failed to load app: %w", err)
	}

	// create validator set
	valSet, err := startupConfig.ValidatorSet()
	if err != nil {
		return nil, errors.New("failed to create validator set")
	}

	var (
		balances    []banktypes.Balance
		genAccounts []authtypes.GenesisAccount
	)
	for _, ga := range startupConfig.GenesisAccounts {
		genAccounts = append(genAccounts, ga.GenesisAccount)
		balances = append(balances, banktypes.Balance{Address: ga.GenesisAccount.GetAddress().String(), Coins: ga.Coins})
	}

	genesisState, err := GenesisStateWithValSet(codec, app.DefaultGenesis(), valSet, genAccounts, balances...)
	if err != nil {
		return nil, fmt.Errorf("failed to create genesis state: %w", err)
	}

	// init chain must be called to stop deliverState from being nil
	stateBytes, err := cmtjson.MarshalIndent(genesisState, "", " ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal default genesis state: %w", err)
	}

	// init chain will set the validator set and initialize the genesis accounts
	_, err = app.InitChain(&abci.InitChainRequest{
		Validators:      []abci.ValidatorUpdate{},
		ConsensusParams: DefaultConsensusParams,
		AppStateBytes:   stateBytes,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init chain: %w", err)
	}

	// commit genesis changes
	if !startupConfig.AtGenesis {
		_, err = app.FinalizeBlock(&abci.FinalizeBlockRequest{
			Height:             app.LastBlockHeight() + 1,
			NextValidatorsHash: valSet.Hash(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to finalize block: %w", err)
		}
	}

	return app, nil
}

// GenesisStateWithValSet returns a new genesis state with the validator set
func GenesisStateWithValSet(
	codec codec.Codec,
	genesisState map[string]json.RawMessage,
	valSet *cmttypes.ValidatorSet,
	genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) (map[string]json.RawMessage, error) {
	// set genesis accounts
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = codec.MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromCmtPubKeyInterface(val.PubKey)
		if err != nil {
			return nil, fmt.Errorf("failed to convert pubkey: %w", err)
		}

		pkAny, err := codectypes.NewAnyWithValue(pk)
		if err != nil {
			return nil, fmt.Errorf("failed to create new any: %w", err)
		}

		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdkmath.LegacyOneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec()),
			MinSelfDelegation: sdkmath.ZeroInt(),
		}
		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress().String(), sdk.ValAddress(val.Address).String(), sdkmath.LegacyOneDec()))

	}

	// set validators and delegations
	stakingGenesis := stakingtypes.NewGenesisState(stakingtypes.DefaultParams(), validators, delegations)
	genesisState[stakingtypes.ModuleName] = codec.MustMarshalJSON(stakingGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		// add genesis acc tokens to total supply
		totalSupply = totalSupply.Add(b.Coins...)
	}

	for range delegations {
		// add delegated tokens to total supply
		totalSupply = totalSupply.Add(sdk.NewCoin(sdk.DefaultBondDenom, bondAmt))
	}

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, bondAmt.MulRaw(int64(len(delegations))))},
	})

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{}, []banktypes.SendEnabled{})
	genesisState[banktypes.ModuleName] = codec.MustMarshalJSON(bankGenesis)

	return genesisState, nil
}

var _ server.DynamicConfig = EmptyAppOptions{}

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

// GetString implements AppOptions
func (ao EmptyAppOptions) GetString(o string) string {
	return ""
}

// AppOptionsMap is a stub implementing AppOptions which can get data from a map
type AppOptionsMap map[string]interface{}

func (m AppOptionsMap) Get(key string) interface{} {
	v, ok := m[key]
	if !ok {
		return interface{}(nil)
	}

	return v
}

func (m AppOptionsMap) GetString(key string) string {
	v, ok := m[key]
	if !ok {
		return ""
	}

	return v.(string)
}
