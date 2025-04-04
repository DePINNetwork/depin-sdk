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
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DePINNetwork/depin-sdk/config"
)

func BenchmarkAlgodStartup(b *testing.B) {
	tmpDir := b.TempDir()
	genesisFile, err := os.ReadFile("../../installer/genesis/devnet/genesis.json")
	require.NoError(b, err)

	dataDirectory = &tmpDir
	bInitAndExit := true
	initAndExit = &bInitAndExit
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		err := os.WriteFile(filepath.Join(tmpDir, config.GenesisJSONFile), genesisFile, 0766)
		require.NoError(b, err)
		fmt.Printf("file %s was written\n", filepath.Join(tmpDir, config.GenesisJSONFile))
		run()
		os.RemoveAll(tmpDir)
		os.Mkdir(tmpDir, 0766)
	}
}
