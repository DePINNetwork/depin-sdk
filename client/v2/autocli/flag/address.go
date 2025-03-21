package flag

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	"cosmossdk.io/client/v2/autocli/keyring"
	clientcontext "cosmossdk.io/client/v2/context"
	"cosmossdk.io/core/address"

	"github.com/depinnetwork/depin-sdk/codec"
	"github.com/depinnetwork/depin-sdk/codec/types"
	cryptocodec "github.com/depinnetwork/depin-sdk/crypto/codec"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
)

type addressStringType struct{}

func (a addressStringType) NewValue(ctx *context.Context, b *Builder) Value {
	return &addressValue{addressCodec: b.AddressCodec, ctx: ctx}
}

func (a addressStringType) DefaultValue() string {
	return ""
}

type validatorAddressStringType struct{}

func (a validatorAddressStringType) NewValue(ctx *context.Context, b *Builder) Value {
	return &addressValue{addressCodec: b.ValidatorAddressCodec, ctx: ctx}
}

func (a validatorAddressStringType) DefaultValue() string {
	return ""
}

type addressValue struct {
	ctx          *context.Context
	addressCodec address.Codec

	value string
}

func (a *addressValue) Get(protoreflect.Value) (protoreflect.Value, error) {
	return protoreflect.ValueOfString(a.value), nil
}

func (a *addressValue) String() string {
	return a.value
}

// Set implements the flag.Value interface for addressValue.
func (a *addressValue) Set(s string) error {
	// we get the keyring on set, as in NewValue the context is the parent context (before RunE)
	k := getKeyringFromCtx(a.ctx)
	addr, err := k.LookupAddressByKeyName(s)
	if err == nil {
		addrStr, err := a.addressCodec.BytesToString(addr)
		if err != nil {
			return fmt.Errorf("invalid account address got from keyring: %w", err)
		}

		a.value = addrStr
		return nil
	}

	_, err = a.addressCodec.StringToBytes(s)
	if err != nil {
		return fmt.Errorf("invalid account address or key name: %w", err)
	}

	a.value = s

	return nil
}

func (a *addressValue) Type() string {
	return "account address or key name"
}

type consensusAddressStringType struct{}

func (a consensusAddressStringType) NewValue(ctx *context.Context, b *Builder) Value {
	return &consensusAddressValue{
		addressValue: addressValue{
			addressCodec: b.ConsensusAddressCodec,
			ctx:          ctx,
		},
	}
}

func (a consensusAddressStringType) DefaultValue() string {
	return ""
}

type consensusAddressValue struct {
	addressValue
}

func (a consensusAddressValue) Get(protoreflect.Value) (protoreflect.Value, error) {
	return protoreflect.ValueOfString(a.value), nil
}

func (a consensusAddressValue) String() string {
	return a.value
}

func (a *consensusAddressValue) Set(s string) error {
	// we get the keyring on set, as in NewValue the context is the parent context (before RunE)
	k := getKeyringFromCtx(a.ctx)
	addr, err := k.LookupAddressByKeyName(s)
	if err == nil {
		addrStr, err := a.addressCodec.BytesToString(addr)
		if err != nil {
			return fmt.Errorf("invalid consensus address got from keyring: %w", err)
		}

		a.value = addrStr
		return nil
	}

	_, err = a.addressCodec.StringToBytes(s)
	if err == nil {
		a.value = s
		return nil
	}

	// fallback to pubkey parsing
	registry := types.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	var pk cryptotypes.PubKey
	err2 := cdc.UnmarshalInterfaceJSON([]byte(s), &pk)
	if err2 != nil {
		return fmt.Errorf("input isn't a pubkey (%w) or is an invalid account address (%w)", err, err2)
	}

	a.value, err = a.addressCodec.BytesToString(pk.Address())
	if err != nil {
		return fmt.Errorf("invalid pubkey address: %w", err)
	}

	return nil
}

// getKeyringFromCtx retrieves the keyring from the provided context.
// If the context is nil or does not contain a valid client context,
// it returns a no-op keyring implementation.
func getKeyringFromCtx(ctx *context.Context) keyring.Keyring {
	if *ctx == nil {
		return keyring.NoKeyring{}
	}

	c, err := clientcontext.ClientContextFromGoContext(*ctx)
	if err != nil {
		return keyring.NoKeyring{}
	}

	return c.Keyring
}
