// Code generated by MockGen. DO NOT EDIT.
// Source: x/group/testutil/expected_keepers.go
//
// Generated by this command:
//
//	mockgen -source=x/group/testutil/expected_keepers.go -package testutil -destination x/group/testutil/expected_keepers_mocks.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	address "cosmossdk.io/core/address"
	types "cosmossdk.io/x/bank/types"
	types0 "github.com/depinnetwork/depin-sdk/types"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
	isgomock struct{}
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// AddressCodec mocks base method.
func (m *MockAccountKeeper) AddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// AddressCodec indicates an expected call of AddressCodec.
func (mr *MockAccountKeeperMockRecorder) AddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressCodec", reflect.TypeOf((*MockAccountKeeper)(nil).AddressCodec))
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(arg0 context.Context, arg1 types0.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), arg0, arg1)
}

// NewAccount mocks base method.
func (m *MockAccountKeeper) NewAccount(arg0 context.Context, arg1 types0.AccountI) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0, arg1)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// NewAccount indicates an expected call of NewAccount.
func (mr *MockAccountKeeperMockRecorder) NewAccount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountKeeper)(nil).NewAccount), arg0, arg1)
}

// RemoveAccount mocks base method.
func (m *MockAccountKeeper) RemoveAccount(ctx context.Context, acc types0.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveAccount", ctx, acc)
}

// RemoveAccount indicates an expected call of RemoveAccount.
func (mr *MockAccountKeeperMockRecorder) RemoveAccount(ctx, acc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAccount", reflect.TypeOf((*MockAccountKeeper)(nil).RemoveAccount), ctx, acc)
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(arg0 context.Context, arg1 types0.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", arg0, arg1)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), arg0, arg1)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
	isgomock struct{}
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// Burn mocks base method.
func (m *MockBankKeeper) Burn(arg0 context.Context, arg1 *types.MsgBurn) (*types.MsgBurnResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Burn", arg0, arg1)
	ret0, _ := ret[0].(*types.MsgBurnResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Burn indicates an expected call of Burn.
func (mr *MockBankKeeperMockRecorder) Burn(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Burn", reflect.TypeOf((*MockBankKeeper)(nil).Burn), arg0, arg1)
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx context.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx context.Context, moduleName string, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amt)
}

// MultiSend mocks base method.
func (m *MockBankKeeper) MultiSend(arg0 context.Context, arg1 *types.MsgMultiSend) (*types.MsgMultiSendResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiSend", arg0, arg1)
	ret0, _ := ret[0].(*types.MsgMultiSendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiSend indicates an expected call of MultiSend.
func (mr *MockBankKeeperMockRecorder) MultiSend(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiSend", reflect.TypeOf((*MockBankKeeper)(nil).MultiSend), arg0, arg1)
}

// Send mocks base method.
func (m *MockBankKeeper) Send(arg0 context.Context, arg1 *types.MsgSend) (*types.MsgSendResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*types.MsgSendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockBankKeeperMockRecorder) Send(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBankKeeper)(nil).Send), arg0, arg1)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types0.AccAddress, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SendCoinsFromModuleToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToModule", ctx, senderModule, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToModule indicates an expected call of SendCoinsFromModuleToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToModule), ctx, senderModule, recipientModule, amt)
}

// SetSendEnabled mocks base method.
func (m *MockBankKeeper) SetSendEnabled(arg0 context.Context, arg1 *types.MsgSetSendEnabled) (*types.MsgSetSendEnabledResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSendEnabled", arg0, arg1)
	ret0, _ := ret[0].(*types.MsgSetSendEnabledResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSendEnabled indicates an expected call of SetSendEnabled.
func (mr *MockBankKeeperMockRecorder) SetSendEnabled(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSendEnabled", reflect.TypeOf((*MockBankKeeper)(nil).SetSendEnabled), arg0, arg1)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx context.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// UpdateParams mocks base method.
func (m *MockBankKeeper) UpdateParams(arg0 context.Context, arg1 *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateParams", arg0, arg1)
	ret0, _ := ret[0].(*types.MsgUpdateParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateParams indicates an expected call of UpdateParams.
func (mr *MockBankKeeperMockRecorder) UpdateParams(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateParams", reflect.TypeOf((*MockBankKeeper)(nil).UpdateParams), arg0, arg1)
}

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
	isgomock struct{}
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// BondDenom mocks base method.
func (m *MockStakingKeeper) BondDenom(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BondDenom", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BondDenom indicates an expected call of BondDenom.
func (mr *MockStakingKeeperMockRecorder) BondDenom(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BondDenom", reflect.TypeOf((*MockStakingKeeper)(nil).BondDenom), ctx)
}
