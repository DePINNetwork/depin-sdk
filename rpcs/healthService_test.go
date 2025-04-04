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

package rpcs

import (
	"io"
	"net/http"
	"path"
	"testing"

	"github.com/DePINNetwork/depin-sdk/network"
	"github.com/DePINNetwork/depin-sdk/network/addr"
	"github.com/DePINNetwork/depin-sdk/test/partitiontest"
	"github.com/stretchr/testify/require"
)

func TestHealthService_ServeHTTP(t *testing.T) {
	partitiontest.PartitionTest(t)

	nodeA := &basicRPCNode{}
	nodeA.start()
	defer nodeA.stop()

	_ = MakeHealthService(nodeA)

	parsedURL, err := addr.ParseHostOrURL(nodeA.rootURL())
	require.NoError(t, err)

	client := http.Client{}

	parsedURL.Path = path.Join(parsedURL.Path, network.HealthServiceStatusPath)

	response, err := client.Get(parsedURL.String())
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, response.StatusCode)
	bodyData, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	require.Empty(t, bodyData)
}
