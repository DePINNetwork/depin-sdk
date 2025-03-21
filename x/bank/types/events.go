package types

import (
	sdk "github.com/depinnetwork/depin-sdk/types"
)

// bank module event types
const (
	EventTypeTransfer = "transfer"

	AttributeKeyRecipient = "recipient"
	AttributeKeySender    = sdk.AttributeKeySender

	// supply and balance tracking events name and attributes
	EventTypeCoinSpent    = "coin_spent"
	EventTypeCoinReceived = "coin_received"
	EventTypeCoinMint     = "coinbase" // NOTE(fdymylja): using mint clashes with mint module event
	EventTypeCoinBurn     = "burn"

	AttributeKeySpender  = "spender"
	AttributeKeyReceiver = "receiver"
	AttributeKeyMinter   = "minter"
	AttributeKeyBurner   = "burner"
)
