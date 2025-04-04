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
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DePINNetwork/depin-sdk/config"
	"github.com/DePINNetwork/depin-sdk/gen"
	"github.com/DePINNetwork/depin-sdk/util"
)

var outDir = flag.String("d", "", "The directory containing the generated ledger and wallet files.")
var netName = flag.String("n", "", "The name of the network for this ledger (will override config file).")
var configFile = flag.String("c", "", "The config file containing the genesis ledger and wallets")
var quiet = flag.Bool("q", false, "Skip verbose informational messages")
var short = flag.Bool("s", false, "Cap the last participation key round to 1500")

func init() {
	flag.Parse()
}

func main() {
	if !*quiet {
		fmt.Println("Network Name: " + *netName)
	}

	cfgFile := *configFile
	if !util.FileExists(cfgFile) {
		reportErrorf("missing configuration file '%s'\n", cfgFile)
	}

	genesisData, err := gen.LoadGenesisData(cfgFile)
	if err != nil {
		reportErrorf("error loading configuration file: %v\n", err)
	}

	if *netName != "" {
		genesisData.NetworkName = *netName
	}

	var verboseOut io.Writer = nil
	if !*quiet {
		verboseOut = os.Stdout
	}

	if *short {
		if genesisData.LastPartKeyRound > 1500 {
			genesisData.LastPartKeyRound = 1500
		}
	}

	err = gen.GenerateGenesisFiles(genesisData, config.Consensus, *outDir, verboseOut)
	if err != nil {
		reportErrorf("Cannot write genesis files: %s", err)
	}
}

func reportErrorf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
