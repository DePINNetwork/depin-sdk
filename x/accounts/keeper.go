package accounts

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"

	gogoproto "github.com/cosmos/gogoproto/proto"
	gogotypes "github.com/cosmos/gogoproto/types"

	_ "cosmossdk.io/api/cosmos/accounts/defaults/base/v1" // import for side-effects
	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/transaction"
	"cosmossdk.io/x/accounts/accountstd"
	"cosmossdk.io/x/accounts/internal/implementation"
	v1 "cosmossdk.io/x/accounts/v1"
	txdecode "cosmossdk.io/x/tx/decode"

	"github.com/depinnetwork/depin-sdk/codec"
	sdk "github.com/depinnetwork/depin-sdk/types"
	authtypes "github.com/depinnetwork/depin-sdk/x/auth/types"
)

var (
	errAccountTypeNotFound = errors.New("account type not found")
	// ErrUnauthorized is returned when a message sender is not allowed to perform the operation.
	ErrUnauthorized = errors.New("unauthorized")
)

var (
	// AccountTypeKeyPrefix is the prefix for the account type key.
	AccountTypeKeyPrefix = collections.NewPrefix(0)
	// AccountNumberKey is the key for the account number.
	AccountNumberKey = collections.NewPrefix(1)
	// AccountByNumber is the key for the accounts by number.
	AccountByNumber = collections.NewPrefix(2)
)

type InterfaceRegistry interface {
	RegisterInterface(name string, iface any, impls ...gogoproto.Message)
	RegisterImplementations(iface any, impls ...gogoproto.Message)
}

func NewKeeper(
	cdc codec.Codec,
	env appmodule.Environment,
	addressCodec address.Codec,
	ir InterfaceRegistry,
	txDecoder *txdecode.Decoder,
	accounts ...accountstd.AccountCreatorFunc,
) (Keeper, error) {
	sb := collections.NewSchemaBuilder(env.KVStoreService)
	keeper := Keeper{
		Environment:      env,
		txDecoder:        txDecoder,
		addressCodec:     addressCodec,
		codec:            cdc,
		makeSendCoinsMsg: defaultCoinsTransferMsgFunc(addressCodec),
		accounts:         nil,
		Schema:           collections.Schema{},
		AccountNumber:    collections.NewSequence(sb, AccountNumberKey, "account_number"),
		AccountsByType:   collections.NewMap(sb, AccountTypeKeyPrefix, "accounts_by_type", collections.BytesKey.WithName("address"), collections.StringValue.WithName("type")),
		AccountByNumber:  collections.NewMap(sb, AccountByNumber, "account_by_number", collections.BytesKey.WithName("address"), collections.Uint64Value.WithName("number")),
		AccountsState: collections.NewMap(sb, implementation.AccountStatePrefix, "accounts_state", collections.NamedPairKeyCodec(
			"number",
			collections.Uint64Key,
			"key",
			collections.BytesKey,
		), collections.BytesValue),
	}

	schema, err := sb.Build()
	if err != nil {
		return Keeper{}, err
	}
	keeper.Schema = schema
	keeper.accounts, err = implementation.MakeAccountsMap(cdc, keeper.addressCodec, env, accounts)
	if err != nil {
		return Keeper{}, err
	}
	registerToInterfaceRegistry(ir, keeper.accounts)
	return keeper, nil
}

type Keeper struct {
	appmodule.Environment

	txDecoder        *txdecode.Decoder
	addressCodec     address.Codec
	codec            codec.Codec
	makeSendCoinsMsg coinsTransferMsgFunc

	accounts map[string]implementation.Implementation

	// Schema is the schema for the module.
	Schema collections.Schema
	// AccountNumber is the last global account number.
	AccountNumber collections.Sequence
	// AccountsByType maps account address to their implementation.
	AccountsByType collections.Map[[]byte, string]
	// AccountByNumber maps account number to their address.
	AccountByNumber collections.Map[[]byte, uint64]

	// AccountsState keeps track of the state of each account.
	// NOTE: this is only used for genesis import and export.
	// Account set and get their own state but this helps providing a nice mapping
	// between: (account number, account state key) => account state value.
	AccountsState collections.Map[collections.Pair[uint64, []byte], []byte]

	bundlingDisabled bool // if this is set then bundling of txs is disallowed.
}

// IsAccountsModuleAccount check if an address belong to a smart account.
func (k Keeper) IsAccountsModuleAccount(
	ctx context.Context,
	accountAddr []byte,
) bool {
	hasAcc, _ := k.AccountByNumber.Has(ctx, accountAddr)
	return hasAcc
}

func (k Keeper) GetAccountNumberLegacy(ctx context.Context) (uint64, error) {
	store := k.KVStoreService.OpenKVStore(ctx)
	b, err := store.Get(authtypes.LegacyGlobalAccountNumberKey)
	if err != nil {
		return 0, fmt.Errorf("failed to get legacy account number: %w", err)
	}
	v := new(gogotypes.UInt64Value)
	if err := v.Unmarshal(b); err != nil {
		return 0, fmt.Errorf("failed to unmarshal legacy account number: %w", err)
	}
	return v.Value, nil
}

func (k Keeper) NextAccountNumber(
	ctx context.Context,
) (accNum uint64, err error) {
	accNum, err = collections.Item[uint64](k.AccountNumber).Get(ctx)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		// This change makes the method works in historical states.
		// Although the behavior is not identical, but semantically compatible.
		//
		// For the state machine, it also does the migration lazily.
		accNum, err = k.GetAccountNumberLegacy(ctx)
	}

	if err != nil {
		return 0, err
	}

	if err := k.AccountNumber.Set(ctx, accNum+1); err != nil {
		return 0, err
	}

	return accNum, nil
}

// InitAccountNumberSeqUnsafe use to set accounts account number tracking.
// Only use for account number migration.
func (k Keeper) InitAccountNumberSeqUnsafe(ctx context.Context, accNum uint64) error {
	currentNum, err := k.AccountNumber.Peek(ctx)
	if err != nil {
		return err
	}
	if currentNum > accNum {
		return fmt.Errorf("cannot set number lower than current account number got %v while current account number is %v", accNum, currentNum)
	}
	return k.AccountNumber.Set(ctx, accNum)
}

// Init creates a new account of the given type.
func (k Keeper) Init(
	ctx context.Context,
	accountType string,
	creator []byte,
	initRequest transaction.Msg,
	funds sdk.Coins,
	addressSeed []byte,
) (transaction.Msg, []byte, error) {
	// get the next account number
	num, err := k.AccountNumber.Next(ctx)
	if err != nil {
		return nil, nil, err
	}
	// create address
	accountAddr, err := k.makeAddress(creator, num, addressSeed)
	if err != nil {
		return nil, nil, err
	}
	initResp, err := k.init(ctx, accountType, creator, num, accountAddr, initRequest, funds)
	if err != nil {
		return nil, nil, err
	}
	return initResp, accountAddr, nil
}

// initFromMsg is a helper which inits an account given a v1.MsgInit.
func (k Keeper) initFromMsg(ctx context.Context, initMsg *v1.MsgInit) (transaction.Msg, []byte, error) {
	creator, err := k.addressCodec.StringToBytes(initMsg.Sender)
	if err != nil {
		return nil, nil, err
	}

	// decode message bytes into the concrete boxed message type
	msg, err := implementation.UnpackAnyRaw(initMsg.Message)
	if err != nil {
		return nil, nil, err
	}

	// run account creation logic
	return k.Init(ctx, initMsg.AccountType, creator, msg, initMsg.Funds, initMsg.AddressSeed)
}

// init initializes the account, given the type, the creator the newly created account number, its address and the
// initialization message.
func (k Keeper) init(
	ctx context.Context,
	accountType string,
	creator []byte,
	accountNum uint64,
	accountAddr []byte,
	initRequest transaction.Msg,
	funds sdk.Coins,
) (transaction.Msg, error) {
	impl, ok := k.accounts[accountType]
	if !ok {
		return nil, fmt.Errorf("%w: not found %s", errAccountTypeNotFound, accountType)
	}

	// check if account exists
	alreadyExists, err := k.AccountsByType.Has(ctx, accountAddr)
	if err != nil {
		return nil, err
	}
	if alreadyExists {
		return nil, ErrAccountAlreadyExists
	}

	// send funds, if provided
	err = k.maybeSendFunds(ctx, creator, accountAddr, funds)
	if err != nil {
		return nil, fmt.Errorf("unable to transfer funds: %w", err)
	}
	// make the context and init the account
	ctx = k.makeAccountContext(ctx, accountNum, accountAddr, creator, funds, false)
	resp, err := impl.Init(ctx, initRequest)
	if err != nil {
		return nil, err
	}

	// map account address to account type
	if err := k.AccountsByType.Set(ctx, accountAddr, accountType); err != nil {
		return nil, err
	}
	// map account number to account address
	if err := k.AccountByNumber.Set(ctx, accountAddr, accountNum); err != nil {
		return nil, err
	}
	return resp, nil
}

// MigrateLegacyAccount is used to migrate a legacy account to x/accounts.
// Concretely speaking this works like Init, but with a custom account number provided,
// Where the creator is the account itself. This can be used by the x/auth module to
// gradually migrate base accounts to x/accounts.
// NOTE: this assumes the calling module checks for account overrides.
func (k Keeper) MigrateLegacyAccount(
	ctx context.Context,
	addr []byte, // The current address of the account
	accNum uint64, // The current account number
	accType string, // The account type to migrate to
	msg transaction.Msg, // The init msg of the account type we're migrating to
) (transaction.Msg, error) {
	return k.init(ctx, accType, addr, accNum, addr, msg, nil)
}

// Execute executes a state transition on the given account.
func (k Keeper) Execute(
	ctx context.Context,
	accountAddr []byte,
	sender []byte,
	execRequest transaction.Msg,
	funds sdk.Coins,
) (transaction.Msg, error) {
	// get account implementation
	impl, err := k.getImplementation(ctx, accountAddr)
	if err != nil {
		// this means the account was initialized with an implementation
		// that the chain does not know about, in theory should never happen,
		// as it might signal that the app-dev stopped supporting an account type.
		return nil, err
	}

	// get account number
	accountNum, err := k.AccountByNumber.Get(ctx, accountAddr)
	if err != nil {
		return nil, err
	}

	err = k.maybeSendFunds(ctx, sender, accountAddr, funds)
	if err != nil {
		return nil, fmt.Errorf("unable to transfer coins to account: %w", err)
	}

	// make the context and execute the account state transition.
	ctx = k.makeAccountContext(ctx, accountNum, accountAddr, sender, funds, false)
	return impl.Execute(ctx, execRequest)
}

// Query queries the given account.
func (k Keeper) Query(
	ctx context.Context,
	accountAddr []byte,
	queryRequest transaction.Msg,
) (transaction.Msg, error) {
	// get account implementation
	impl, err := k.getImplementation(ctx, accountAddr)
	if err != nil {
		// this means the account was initialized with an implementation
		// that the chain does not know about, in theory should never happen,
		// as it might signal that the app-dev stopped supporting an account type.
		return nil, err
	}

	accountNum, err := k.AccountByNumber.Get(ctx, accountAddr)
	if err != nil {
		return nil, err
	}

	// make the context and execute the account query
	ctx = k.makeAccountContext(ctx, accountNum, accountAddr, nil, nil, true)
	return impl.Query(ctx, queryRequest)
}

func (k Keeper) getImplementation(ctx context.Context, addr []byte) (implementation.Implementation, error) {
	accountType, err := k.AccountsByType.Get(ctx, addr)
	if err != nil {
		return implementation.Implementation{}, err
	}
	impl, ok := k.accounts[accountType]
	if !ok {
		return implementation.Implementation{}, fmt.Errorf("%w: %s", errAccountTypeNotFound, accountType)
	}
	return impl, nil
}

// makeAddress creates an address for the given account.
// It uses the creator address to ensure address squatting cannot happen, for example
// assuming creator sends funds to a new account X nobody can front-run that address instantiation
// unless the creator itself sends the tx.
// AddressSeed can be used to create predictable addresses, security guarantees of the above are retained.
// If address seed is not provided, the address is created using the creator and account number.
func (k Keeper) makeAddress(creator []byte, accNum uint64, addressSeed []byte) ([]byte, error) {
	// in case an address seed is provided, we use it to create the address.
	var seed []byte
	if len(addressSeed) > 0 {
		seed = append(creator, addressSeed...)
	} else {
		// otherwise we use the creator and account number to create the address.
		seed = append(creator, binary.BigEndian.AppendUint64(nil, accNum)...)
	}

	moduleAndSeed := append([]byte(ModuleName), seed...)

	addr := sha256.Sum256(moduleAndSeed)

	return addr[:], nil
}

// makeAccountContext makes a new context for the given account.
func (k Keeper) makeAccountContext(ctx context.Context, accountNumber uint64, accountAddr, sender []byte, funds sdk.Coins, isQuery bool) context.Context {
	// if it's not a query we create a context that allows to do anything.
	if !isQuery {
		return implementation.MakeAccountContext(
			ctx,
			k.KVStoreService,
			accountNumber,
			accountAddr,
			sender,
			funds,
			k.SendModuleMessage,
			k.queryModule,
		)
	}

	// if it's a query we create a context that does not allow to execute modules
	// and does not allow to get the sender.
	return implementation.MakeAccountContext(
		ctx,
		k.KVStoreService,
		accountNumber,
		accountAddr,
		nil,
		nil,
		func(ctx context.Context, sender []byte, msg transaction.Msg) (transaction.Msg, error) {
			return nil, errors.New("cannot execute in query context")
		},
		k.queryModule,
	)
}

// sendAnyMessages it a helper function that executes untyped codectypes.Any messages
// The messages must all belong to a module.
func (k Keeper) sendAnyMessages(ctx context.Context, sender []byte, anyMessages []*implementation.Any) ([]*implementation.Any, error) {
	anyResponses := make([]*implementation.Any, len(anyMessages))
	for i := range anyMessages {
		msg, err := implementation.UnpackAnyRaw(anyMessages[i])
		if err != nil {
			return nil, err
		}
		resp, err := k.SendModuleMessage(ctx, sender, msg)
		if err != nil {
			return nil, fmt.Errorf("failed to execute message %d: %s", i, err.Error())
		}
		anyResp, err := implementation.PackAny(resp)
		if err != nil {
			return nil, err
		}
		anyResponses[i] = anyResp
	}
	return anyResponses, nil
}

func (k Keeper) sendManyMessagesReturnAnys(ctx context.Context, sender []byte, msgs []transaction.Msg) ([]*implementation.Any, error) {
	resp, err := k.sendManyMessages(ctx, sender, msgs)
	if err != nil {
		return nil, err
	}
	anys := make([]*implementation.Any, len(resp))
	for i := range resp {
		anypb, err := implementation.PackAny(resp[i])
		if err != nil {
			return nil, err
		}
		anys[i] = anypb
	}
	return anys, nil
}

// sendManyMessages is a helper function that sends many untyped messages on behalf of the sender
// then returns the respective results. Since the function calls into SendModuleMessage
// it is guaranteed to disallow impersonation attacks from the sender.
func (k Keeper) sendManyMessages(ctx context.Context, sender []byte, msgs []transaction.Msg) ([]transaction.Msg, error) {
	resps := make([]transaction.Msg, len(msgs))
	for i, msg := range msgs {
		resp, err := k.SendModuleMessage(ctx, sender, msg)
		if err != nil {
			return nil, fmt.Errorf("failed to execute message %d: %s", i, err.Error())
		}
		resps[i] = resp
	}
	return resps, nil
}

// SendModuleMessage can be used to send a message towards a module.
// It should be used when the response type is not known by the caller.
func (k Keeper) SendModuleMessage(ctx context.Context, sender []byte, msg transaction.Msg) (transaction.Msg, error) {
	// do sender assertions.
	wantSenders, _, err := k.codec.GetMsgSigners(msg)
	if err != nil {
		return nil, fmt.Errorf("cannot get signers: %w", err)
	}
	if len(wantSenders) != 1 {
		return nil, fmt.Errorf("expected only one signer, got %d", len(wantSenders))
	}
	if !bytes.Equal(sender, wantSenders[0]) {
		return nil, fmt.Errorf("%w: sender does not match expected sender", ErrUnauthorized)
	}
	resp, err := k.MsgRouterService.Invoke(ctx, msg)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// sendModuleMessage can be used to send a message towards a module. It expects the
// response type to be known by the caller. It will also assert the sender has the right
// is not trying to impersonate another account.
func (k Keeper) sendModuleMessage(ctx context.Context, sender []byte, msg transaction.Msg) (transaction.Msg, error) {
	// do sender assertions.
	wantSenders, _, err := k.codec.GetMsgSigners(msg)
	if err != nil {
		return nil, fmt.Errorf("cannot get signers: %w", err)
	}
	if len(wantSenders) != 1 {
		return nil, fmt.Errorf("expected only one signer, got %d", len(wantSenders))
	}
	if !bytes.Equal(sender, wantSenders[0]) {
		return nil, fmt.Errorf("%w: sender does not match expected sender", ErrUnauthorized)
	}

	return k.MsgRouterService.Invoke(ctx, msg)
}

// queryModule is the entrypoint for an account to query a module.
// It will try to find the query handler for the given query and execute it.
// If multiple query handlers are found, it will return an error.
func (k Keeper) queryModule(ctx context.Context, queryReq transaction.Msg) (transaction.Msg, error) {
	return k.QueryRouterService.Invoke(ctx, queryReq)
}

// maybeSendFunds will send the provided coins between the provided addresses, if amt
// is not empty.
func (k Keeper) maybeSendFunds(ctx context.Context, from, to []byte, amt sdk.Coins) error {
	if amt.IsZero() {
		return nil
	}

	msg, err := k.makeSendCoinsMsg(from, to, amt)
	if err != nil {
		return err
	}

	// send module message ensures that "from" cannot impersonate.
	_, err = k.sendModuleMessage(ctx, from, msg)
	if err != nil {
		return err
	}

	return nil
}

func (k *Keeper) DisableTxBundling() {
	k.bundlingDisabled = true
}

const msgInterfaceName = "cosmos.accounts.v1.MsgInterface"

// creates a new interface type which is an alias of the proto message interface to avoid conflicts with sdk.Msg
type msgInterface transaction.Msg

var msgInterfaceType = (*msgInterface)(nil)

// registerToInterfaceRegistry registers all the interfaces of the accounts to the
// global interface registry. This is required for the SDK to correctly decode
// the google.Protobuf.Any used in x/accounts.
func registerToInterfaceRegistry(ir InterfaceRegistry, accMap map[string]implementation.Implementation) {
	ir.RegisterInterface(msgInterfaceName, msgInterfaceType)

	for _, acc := range accMap {
		// register init
		ir.RegisterImplementations(msgInterfaceType, acc.InitHandlerSchema.RequestSchema.New(), acc.InitHandlerSchema.ResponseSchema.New())
		// register exec
		for _, exec := range acc.ExecuteHandlersSchema {
			ir.RegisterImplementations(msgInterfaceType, exec.RequestSchema.New(), exec.ResponseSchema.New())
		}
		// register query
		for _, query := range acc.QueryHandlersSchema {
			ir.RegisterImplementations(msgInterfaceType, query.RequestSchema.New(), query.ResponseSchema.New())
		}
	}
}
