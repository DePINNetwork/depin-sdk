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

package fixtures

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/DePINNetwork/depin-sdk/config"
)

type baseFixture struct {
	Config config.Global

	binDir      string
	testDataDir string
	testDir     string
	testDirTmp  bool
	instance    Fixture
}

func getTestDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		return path.Join(path.Dir(filename), "..", "..")
	}
	// fallback to the legacy GOPATH location.
	return os.ExpandEnv("${GOPATH}/src/github.com/DePINNetwork/depin-sdk/test/")
}

func (f *baseFixture) initialize(instance Fixture) {
	f.instance = instance
	f.Config = config.Protocol
	f.binDir = os.Getenv("NODEBINDIR")
	if f.binDir == "" {
		f.binDir = os.ExpandEnv("$GOPATH/bin")
	}
	f.testDir = os.Getenv("TESTDIR")
	if f.testDir == "" {
		f.testDir, _ = os.MkdirTemp("", "tmp")
		f.testDirTmp = true
	}
	f.testDataDir = os.Getenv("TESTDATADIR")
	if f.testDataDir == "" {
		f.testDataDir = path.Join(getTestDir(), "testdata")
	}
}

func (f *baseFixture) run(m *testing.M) int {
	return m.Run()
}

func (f *baseFixture) runAndExit(m *testing.M) {
	ret := m.Run()
	preserveData := ret != 0 // If ret != 0, something failed so preserve data
	f.instance.ShutdownImpl(preserveData)
	os.Exit(ret)
}

func (f *baseFixture) failOnError(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf(message, err))
	}
}
