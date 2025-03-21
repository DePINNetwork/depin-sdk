package lockup

import (
	sdk "github.com/depinnetwork/depin-sdk/types"
	sdkerrors "github.com/depinnetwork/depin-sdk/types/errors"
)

func validateAmount(amount sdk.Coins) error {
	if !amount.IsValid() {
		return sdkerrors.ErrInvalidCoins.Wrap(amount.String())
	}

	if amount.IsZero() {
		return sdkerrors.ErrInvalidCoins.Wrap(amount.String())
	}

	return nil
}
