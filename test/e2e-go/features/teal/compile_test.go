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

package teal

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DePINNetwork/depin-sdk/config"
	"github.com/DePINNetwork/depin-sdk/crypto"
	"github.com/DePINNetwork/depin-sdk/data/transactions/logic"
	"github.com/DePINNetwork/depin-sdk/test/framework/fixtures"
	"github.com/DePINNetwork/depin-sdk/test/partitiontest"
)

func TestTealCompile(t *testing.T) {
	partitiontest.PartitionTest(t)
	defer fixtures.ShutdownSynchronizedTest(t)

	if testing.Short() {
		t.Skip()
	}
	a := require.New(fixtures.SynchronizedTest(t))

	var fixture fixtures.RestClientFixture
	fixture.SetupNoStart(t, filepath.Join("nettemplates", "OneNodeFuture.json"))

	// Get primary node
	primaryNode, err := fixture.GetNodeController("Primary")
	a.NoError(err)

	fixture.Start()
	defer primaryNode.FullStop()

	// get lib goal client
	libGoalClient := fixture.LibGoalFixture.GetLibGoalClientFromNodeController(primaryNode)

	compiledProgram, _, _, err := libGoalClient.Compile([]byte(""), false)
	a.Nil(compiledProgram)
	a.Equal(err.Error(), "HTTP 404 Not Found: /teal/compile was not enabled in the configuration file by setting the EnableDeveloperAPI to true")

	primaryNode.FullStop()

	// update the configuration file to enable the developer API
	cfg, err := config.LoadConfigFromDisk(primaryNode.GetDataDir())
	a.NoError(err)
	cfg.EnableDeveloperAPI = true
	cfg.SaveToDisk(primaryNode.GetDataDir())

	fixture.Start()

	var hash crypto.Digest
	var srcMap *logic.SourceMap
	compiledProgram, hash, srcMap, err = libGoalClient.Compile([]byte("int 1"), true)
	a.NotNil(compiledProgram)
	a.NoError(err, "A valid v1 program should result in a compilation success")
	a.NotNil(srcMap)
	a.Equal([]byte{0x1, 0x20, 0x1, 0x1, 0x22}, compiledProgram)
	a.Equal("6Z3C3LDVWGMX23BMSYMANACQOSINPFIRF77H7N3AWJZYV6OH6GWQ", hash.String())

	compiledProgram, hash, srcMap, err = libGoalClient.Compile([]byte("#pragma version 2\nint 1"), true)
	a.NotNil(compiledProgram)
	a.NoError(err, "A valid v2 program should result in a compilation success")
	a.NotNil(srcMap)
	a.Equal([]byte{0x2, 0x20, 0x1, 0x1, 0x22}, compiledProgram)
	a.Equal("YOE6C22GHCTKAN3HU4SE5PGIPN5UKXAJTXCQUPJ3KKF5HOAH646A", hash.String())

	compiledProgram, hash, _, err = libGoalClient.Compile([]byte("bad program"), false)
	a.Error(err, "An invalid program should result in a compilation failure")
	a.Nil(compiledProgram)
	a.Equal(crypto.Digest{}, hash)
}
