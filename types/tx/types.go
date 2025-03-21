package tx

import (
	"errors"

	gogoprotoany "github.com/cosmos/gogoproto/types/any"
	"google.golang.org/protobuf/reflect/protoreflect"

	"cosmossdk.io/core/registry"
	errorsmod "cosmossdk.io/errors"

	"github.com/depinnetwork/depin-sdk/codec"
	cryptotypes "github.com/depinnetwork/depin-sdk/crypto/types"
	sdk "github.com/depinnetwork/depin-sdk/types"
	sdkerrors "github.com/depinnetwork/depin-sdk/types/errors"
)

// MaxGasWanted defines the max gas allowed.
const MaxGasWanted = uint64((1 << 63) - 1)

// Interface implementation checks.
var (
	_, _, _, _ gogoprotoany.UnpackInterfacesMessage = &Tx{}, &TxBody{}, &AuthInfo{}, &SignerInfo{}
)

// GetMsgs implements the GetMsgs method on sdk.Tx.
func (t *Tx) GetMsgs() []sdk.Msg {
	if t == nil || t.Body == nil {
		return nil
	}

	anys := t.Body.Messages
	res, err := GetMsgs(anys, "transaction")
	if err != nil {
		panic(err)
	}
	return res
}

// ValidateBasic implements the ValidateBasic method on sdk.Tx.
func (t *Tx) ValidateBasic() error {
	if t == nil {
		return errors.New("bad Tx")
	}

	body := t.Body
	if body == nil {
		return errors.New("missing TxBody")
	}

	authInfo := t.AuthInfo
	if authInfo == nil {
		return errors.New("missing AuthInfo")
	}

	fee := authInfo.Fee
	if fee == nil {
		return errors.New("missing fee")
	}

	if fee.GasLimit > MaxGasWanted {
		return errorsmod.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid gas supplied; %d > %d", fee.GasLimit, MaxGasWanted,
		)
	}

	if fee.Amount.IsAnyNil() {
		return errorsmod.Wrapf(
			sdkerrors.ErrInsufficientFee,
			"invalid fee provided: null",
		)
	}

	if fee.Amount.IsAnyNegative() {
		return errorsmod.Wrapf(
			sdkerrors.ErrInsufficientFee,
			"invalid fee provided: %s", fee.Amount,
		)
	}

	if fee.Payer != "" {
		_, err := sdk.AccAddressFromBech32(fee.Payer)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid fee payer address (%s)", err)
		}
	}

	sigs := t.Signatures

	if len(sigs) == 0 {
		return sdkerrors.ErrNoSignatures
	}

	return nil
}

// GetSigners retrieves all the signers of a tx.
// This includes all unique signers of the messages (in order),
// as well as the FeePayer (if specified and not already included).
func (t *Tx) GetSigners(cdc codec.Codec) ([][]byte, []protoreflect.Message, error) {
	var signers [][]byte
	seen := map[string]bool{}

	var reflectMsgs []protoreflect.Message
	for _, msg := range t.Body.Messages {
		xs, reflectMsg, err := cdc.GetMsgAnySigners(msg)
		if err != nil {
			return nil, nil, err
		}

		reflectMsgs = append(reflectMsgs, reflectMsg)

		for _, signer := range xs {
			if !seen[string(signer)] {
				signers = append(signers, signer)
				seen[string(signer)] = true
			}
		}
	}

	// ensure any specified fee payer is included in the required signers (at the end)
	feePayer := t.AuthInfo.Fee.Payer
	var feePayerAddr []byte
	if feePayer != "" {
		var err error
		feePayerAddr, err = cdc.InterfaceRegistry().SigningContext().AddressCodec().StringToBytes(feePayer)
		if err != nil {
			return nil, nil, err
		}
	}
	if feePayerAddr != nil && !seen[string(feePayerAddr)] {
		signers = append(signers, feePayerAddr)
		seen[string(feePayerAddr)] = true
	}

	return signers, reflectMsgs, nil
}

func (t *Tx) GetGas() uint64 {
	return t.AuthInfo.Fee.GasLimit
}

func (t *Tx) GetFee() sdk.Coins {
	return t.AuthInfo.Fee.Amount
}

func (t *Tx) FeePayer(cdc codec.Codec) []byte {
	feePayer := t.AuthInfo.Fee.Payer
	if feePayer != "" {
		feePayerAddr, err := cdc.InterfaceRegistry().SigningContext().AddressCodec().StringToBytes(feePayer)
		if err != nil {
			panic(err)
		}
		return feePayerAddr
	}
	// use first signer as default if no payer specified
	signers, _, err := t.GetSigners(cdc)
	if err != nil {
		panic(err)
	}

	return signers[0]
}

func (t *Tx) FeeGranter(cdc codec.Codec) []byte {
	feeGranter := t.AuthInfo.Fee.Granter
	if feeGranter != "" {
		feeGranterAddr, err := cdc.InterfaceRegistry().SigningContext().AddressCodec().StringToBytes(feeGranter)
		if err != nil {
			panic(err)
		}

		return feeGranterAddr
	}
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (t *Tx) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	if t.Body != nil {
		if err := t.Body.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}

	if t.AuthInfo != nil {
		return t.AuthInfo.UnpackInterfaces(unpacker)
	}

	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (m *TxBody) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	if err := UnpackInterfaces(unpacker, m.Messages); err != nil {
		return err
	}

	if err := unpackTxExtensionOptionsI(unpacker, m.ExtensionOptions); err != nil {
		return err
	}

	if err := unpackTxExtensionOptionsI(unpacker, m.NonCriticalExtensionOptions); err != nil {
		return err
	}

	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (m *AuthInfo) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	for _, signerInfo := range m.SignerInfos {
		err := signerInfo.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (m *SignerInfo) UnpackInterfaces(unpacker gogoprotoany.AnyUnpacker) error {
	return unpacker.UnpackAny(m.PublicKey, new(cryptotypes.PubKey))
}

// RegisterInterfaces registers the sdk.Tx and MsgResponse interfaces.
// Note: the registration of sdk.Msg is done in sdk.RegisterInterfaces, but it
// could be moved inside this function.
func RegisterInterfaces(registry registry.InterfaceRegistrar) {
	registry.RegisterInterface(msgResponseInterfaceProtoName, (*MsgResponse)(nil))

	registry.RegisterInterface("cosmos.tx.v1beta1.Tx", (*sdk.HasMsgs)(nil))
	registry.RegisterImplementations((*sdk.HasMsgs)(nil), &Tx{})

	registry.RegisterInterface("cosmos.tx.v1beta1.TxExtensionOptionI", (*TxExtensionOptionI)(nil))
}
