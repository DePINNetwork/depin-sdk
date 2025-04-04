// Copyright (C) 2019-2025 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DePINNetwork/depin-sdk/config"
	"github.com/DePINNetwork/depin-sdk/data/basics"
	"github.com/DePINNetwork/depin-sdk/data/transactions"
	"github.com/DePINNetwork/depin-sdk/data/transactions/logic"
	"github.com/DePINNetwork/depin-sdk/protocol"
	"github.com/DePINNetwork/depin-sdk/test/partitiontest"
)

// Current implementation uses LegderForCowBase interface to plug into evaluator.
// LedgerForLogic in this case is created inside ledger package, and it is the same
// as used in on-chain evaluation.
// This test ensures TEAL program sees data provided by LegderForCowBase, and sees all
// intermediate changes.
func TestBalanceAdapterStateChanges(t *testing.T) {
	partitiontest.PartitionTest(t)
	t.Parallel()

	a := require.New(t)

	source := `#pragma version 2
// read initial value, must be 1
byte "gkeyint"
app_global_get
int 2
==
// write a new value
byte "gkeyint"
int 3
app_global_put
// read updated value, must be 2
byte "gkeyint"
app_global_get
int 3
==
&&
//
// repeat the same for some local key
//
int 0
byte "lkeyint"
app_local_get
int 1
==
&&
int 0
byte "lkeyint"
int 2
app_local_put
int 0
byte "lkeyint"
app_local_get
int 2
==
&&
`
	ops, err := logic.AssembleString(source)
	a.NoError(err)
	program := ops.Program
	addr, err := basics.UnmarshalChecksumAddress("47YPQTIGQEO7T4Y4RWDYWEKV6RTR2UNBQXBABEEGM72ESWDQNCQ52OPASU")
	a.NoError(err)

	assetIdx := basics.AssetIndex(50)
	appIdx := basics.AppIndex(1001)
	br := makeSampleBalanceRecord(addr, assetIdx, appIdx)
	balances := map[basics.Address]basics.AccountData{
		addr: br.AccountData,
	}

	// make transaction group: app call + sample payment
	txn := transactions.SignedTxn{
		Txn: transactions.Transaction{
			Type: protocol.ApplicationCallTx,
			Header: transactions.Header{
				Sender: addr,
				Fee:    basics.MicroAlgos{Raw: 100},
				Note:   []byte{1, 2, 3},
			},
			ApplicationCallTxnFields: transactions.ApplicationCallTxnFields{
				ApplicationID:   appIdx,
				ApplicationArgs: [][]byte{[]byte("ALGO"), []byte("RAND")},
			},
		},
	}

	ba, _, err := makeBalancesAdapter(
		balances, []transactions.SignedTxn{txn}, 0, string(protocol.ConsensusCurrentVersion),
		100, 102030, appIdx, false, "", "",
	)
	a.NoError(err)

	proto := config.Consensus[protocol.ConsensusCurrentVersion]
	ep := logic.NewAppEvalParams([]transactions.SignedTxnWithAD{{SignedTxn: txn}}, &proto, &transactions.SpecialAddresses{})
	pass, delta, err := ba.StatefulEval(0, ep, appIdx, program)
	a.NoError(err)
	a.True(pass)
	a.Len(delta.GlobalDelta, 1)
	a.Equal(basics.SetUintAction, delta.GlobalDelta["gkeyint"].Action)
	a.Equal(uint64(3), delta.GlobalDelta["gkeyint"].Uint)
	a.Len(delta.LocalDeltas, 1)
	a.Len(delta.LocalDeltas[0], 1)
	a.Equal(basics.SetUintAction, delta.LocalDeltas[0]["lkeyint"].Action)
	a.Equal(uint64(2), delta.LocalDeltas[0]["lkeyint"].Uint)
}
