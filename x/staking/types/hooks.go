package types

import (
	"context"

	sdkmath "cosmossdk.io/math"

	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
)

// combine multiple staking hooks, all hook functions are run in array sequence
var _ StakingHooks = &MultiStakingHooks{}

type MultiStakingHooks []StakingHooks

func NewMultiStakingHooks(hooks ...StakingHooks) MultiStakingHooks {
	return hooks
}

func (h MultiStakingHooks) AfterValidatorCreated(ctx context.Context, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].AfterValidatorCreated(ctx, valAddr); err != nil {
			return err
		}
	}

	return nil
}

func (h MultiStakingHooks) BeforeValidatorModified(ctx context.Context, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].BeforeValidatorModified(ctx, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) AfterValidatorRemoved(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].AfterValidatorRemoved(ctx, consAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) AfterValidatorBonded(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].AfterValidatorBonded(ctx, consAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) AfterValidatorBeginUnbonding(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].AfterValidatorBeginUnbonding(ctx, consAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) BeforeDelegationCreated(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].BeforeDelegationCreated(ctx, delAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) BeforeDelegationSharesModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].BeforeDelegationSharesModified(ctx, delAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) BeforeDelegationRemoved(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].BeforeDelegationRemoved(ctx, delAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) AfterDelegationModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	for i := range h {
		if err := h[i].AfterDelegationModified(ctx, delAddr, valAddr); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) BeforeValidatorSlashed(ctx context.Context, valAddr sdk.ValAddress, fraction sdkmath.LegacyDec) error {
	for i := range h {
		if err := h[i].BeforeValidatorSlashed(ctx, valAddr, fraction); err != nil {
			return err
		}
	}
	return nil
}

func (h MultiStakingHooks) AfterConsensusPubKeyUpdate(ctx context.Context, oldPubKey, newPubKey cryptotypes.PubKey, rotationFee sdk.Coin) error {
	for i := range h {
		if err := h[i].AfterConsensusPubKeyUpdate(ctx, oldPubKey, newPubKey, rotationFee); err != nil {
			return err
		}
	}
	return nil
}
