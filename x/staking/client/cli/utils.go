package cli

import (
	"encoding/json"
	"errors"
	"os"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"cosmossdk.io/x/staking/types"

	"github.com/depinnetwork/depin-sdk/codec"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
	sdkerrors "github.com/depinnetwork/depin-sdk/types/errors"
)

// validator struct to define the fields of the validator
type validator struct {
	Amount            sdk.Coin
	PubKey            cryptotypes.PubKey
	Moniker           string
	Identity          string
	Website           string
	Security          string
	Details           string
	Metadata          *types.Metadata
	CommissionRates   types.CommissionRates
	MinSelfDelegation math.Int
}

func parseAndValidateValidatorJSON(cdc codec.Codec, path string) (validator, error) {
	type internalVal struct {
		Amount                    string          `json:"amount"`
		PubKey                    json.RawMessage `json:"pubkey"`
		Moniker                   string          `json:"moniker"`
		Identity                  string          `json:"identity,omitempty"`
		Website                   string          `json:"website,omitempty"`
		Security                  string          `json:"security,omitempty"`
		Details                   string          `json:"details,omitempty"`
		MetadataProfilePicUri     string          `json:"metadata-profile-pic-uri,omitempty"`
		MetadataSocialHandlesUris []string        `json:"metadata-social-handles-uris,omitempty"`
		CommissionRate            string          `json:"commission-rate"`
		CommissionMaxRate         string          `json:"commission-max-rate"`
		CommissionMaxChange       string          `json:"commission-max-change-rate"`
		MinSelfDelegation         string          `json:"min-self-delegation"`
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		return validator{}, err
	}

	var v internalVal
	err = json.Unmarshal(contents, &v)
	if err != nil {
		return validator{}, err
	}

	if v.Amount == "" {
		return validator{}, errors.New("must specify amount of coins to bond")
	}
	amount, err := sdk.ParseCoinNormalized(v.Amount)
	if err != nil {
		return validator{}, err
	}

	if v.PubKey == nil {
		return validator{}, errors.New("must specify the JSON encoded pubkey")
	}
	var pk cryptotypes.PubKey
	if err := cdc.UnmarshalInterfaceJSON(v.PubKey, &pk); err != nil {
		return validator{}, err
	}

	if v.Moniker == "" {
		return validator{}, errors.New("must specify the moniker name")
	}

	commissionRates, err := buildCommissionRates(v.CommissionRate, v.CommissionMaxRate, v.CommissionMaxChange)
	if err != nil {
		return validator{}, err
	}

	if v.MinSelfDelegation == "" {
		return validator{}, errors.New("must specify minimum self delegation")
	}
	minSelfDelegation, ok := math.NewIntFromString(v.MinSelfDelegation)
	if !ok {
		return validator{}, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "minimum self delegation must be a positive integer")
	}

	metadata, err := buildMetadata(v.MetadataProfilePicUri, v.MetadataSocialHandlesUris)
	if err != nil {
		return validator{}, err
	}

	return validator{
		Amount:            amount,
		PubKey:            pk,
		Moniker:           v.Moniker,
		Identity:          v.Identity,
		Website:           v.Website,
		Security:          v.Security,
		Details:           v.Details,
		Metadata:          metadata,
		CommissionRates:   commissionRates,
		MinSelfDelegation: minSelfDelegation,
	}, nil
}

func buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr string) (commission types.CommissionRates, err error) {
	if rateStr == "" || maxRateStr == "" || maxChangeRateStr == "" {
		return commission, errors.New("must specify all validator commission parameters")
	}

	rate, err := math.LegacyNewDecFromStr(rateStr)
	if err != nil {
		return commission, err
	}

	maxRate, err := math.LegacyNewDecFromStr(maxRateStr)
	if err != nil {
		return commission, err
	}

	maxChangeRate, err := math.LegacyNewDecFromStr(maxChangeRateStr)
	if err != nil {
		return commission, err
	}

	commission = types.NewCommissionRates(rate, maxRate, maxChangeRate)

	return commission, nil
}

func buildMetadata(profilePicUri string, socialHandlesUris []string) (*types.Metadata, error) {
	metadata := types.Metadata{
		ProfilePicUri:    profilePicUri,
		SocialHandleUris: socialHandlesUris,
	}

	if err := metadata.Validate(); err != nil {
		return nil, err
	}

	return &metadata, nil
}
