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

//go:build !windows
// +build !windows

package nodecontrol

import (
	"os"

	"github.com/DePINNetwork/depin-sdk/logging"
)

func (kc *KMDController) isDirectorySafe(dirStats os.FileInfo) bool {
	if (dirStats.Mode() & 0077) != 0 {
		logging.Base().Errorf("%s: kmd data dir exists but is too permissive (%o), change to (%o)", kc.kmdDataDir, dirStats.Mode()&0777, DefaultKMDDataDirPerms)
		return false
	}
	return true
}
