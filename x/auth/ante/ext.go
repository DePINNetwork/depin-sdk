package ante

import (
	"context"

	codectypes "github.com/depinnetwork/depin-sdk/codec/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
	sdkerrors "github.com/depinnetwork/depin-sdk/types/errors"
)

type HasExtensionOptionsTx interface {
	GetExtensionOptions() []*codectypes.Any
	GetNonCriticalExtensionOptions() []*codectypes.Any
}

// ExtensionOptionChecker is a function that returns true if the extension option is accepted.
type ExtensionOptionChecker func(*codectypes.Any) bool

// rejectExtensionOption is the default extension check that reject all tx
// extensions.
func rejectExtensionOption(*codectypes.Any) bool {
	return false
}

// RejectExtensionOptionsDecorator is an AnteDecorator that rejects all extension
// options which can optionally be included in protobuf transactions. Users that
// need extension options should create a custom AnteHandler chain that handles
// needed extension options properly and rejects unknown ones.
type RejectExtensionOptionsDecorator struct {
	checker ExtensionOptionChecker
}

// NewExtensionOptionsDecorator creates a new antehandler that rejects all extension
// options which can optionally be included in protobuf transactions that don't pass the checker.
// Users that need extension options should pass a custom checker that returns true for the
// needed extension options.
func NewExtensionOptionsDecorator(checker ExtensionOptionChecker) RejectExtensionOptionsDecorator {
	if checker == nil {
		checker = rejectExtensionOption
	}

	return RejectExtensionOptionsDecorator{checker: checker}
}

var _ sdk.AnteDecorator = RejectExtensionOptionsDecorator{}

func (r RejectExtensionOptionsDecorator) ValidateTx(ctx context.Context, tx sdk.Tx) error {
	return checkExtOpts(tx, r.checker)
}

// AnteHandle implements the AnteDecorator.AnteHandle method
func (r RejectExtensionOptionsDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, _ bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if err := r.ValidateTx(ctx, tx); err != nil {
		return ctx, err
	}

	return next(ctx, tx, false)
}

func checkExtOpts(tx sdk.Tx, checker ExtensionOptionChecker) error {
	if hasExtOptsTx, ok := tx.(HasExtensionOptionsTx); ok {
		for _, opt := range hasExtOptsTx.GetExtensionOptions() {
			if !checker(opt) {
				return sdkerrors.ErrUnknownExtensionOptions
			}
		}
	}

	return nil
}
