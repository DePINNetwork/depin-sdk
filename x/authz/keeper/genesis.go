package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/x/authz"

	sdk "github.com/depinnetwork/depin-sdk/types"
)

// InitGenesis initializes new authz genesis
func (k Keeper) InitGenesis(ctx context.Context, data *authz.GenesisState) error {
	now := k.HeaderService.HeaderInfo(ctx).Time
	for _, entry := range data.Authorization {
		// ignore expired authorizations
		if entry.Expiration != nil && entry.Expiration.Before(now) {
			continue
		}

		grantee, err := k.addrCdc.StringToBytes(entry.Grantee)
		if err != nil {
			return err
		}
		granter, err := k.addrCdc.StringToBytes(entry.Granter)
		if err != nil {
			return err
		}

		a, ok := entry.Authorization.GetCachedValue().(authz.Authorization)
		if !ok {
			return errors.New("expected authorization")
		}

		err = k.SaveGrant(ctx, grantee, granter, a, entry.Expiration)
		if err != nil {
			return err
		}
	}
	return nil
}

// ExportGenesis returns a GenesisState for a given context.
func (k Keeper) ExportGenesis(ctx context.Context) (*authz.GenesisState, error) {
	var entries []authz.GrantAuthorization
	err := k.IterateGrants(ctx, func(granter, grantee sdk.AccAddress, grant authz.Grant) (bool, error) {
		granterAddr, err := k.addrCdc.BytesToString(granter)
		if err != nil {
			return false, err
		}
		granteeAddr, err := k.addrCdc.BytesToString(grantee)
		if err != nil {
			return false, err
		}
		entries = append(entries, authz.GrantAuthorization{
			Granter:       granterAddr,
			Grantee:       granteeAddr,
			Expiration:    grant.Expiration,
			Authorization: grant.Authorization,
		})
		return false, nil
	})
	if err != nil {
		return nil, err
	}
	return authz.NewGenesisState(entries), nil
}
