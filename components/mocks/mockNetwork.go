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

package mocks

import (
	"context"
	"errors"
	"net/http"

	"github.com/DePINNetwork/depin-sdk/network"
	"github.com/DePINNetwork/depin-sdk/protocol"
)

// MockNetwork is a dummy network that doesn't do anything
type MockNetwork struct {
	network.GossipNode
	GenesisID string
}

// Broadcast - unused function
func (network *MockNetwork) Broadcast(ctx context.Context, tag protocol.Tag, data []byte, wait bool, except network.Peer) error {
	return nil
}

// Relay - unused function
func (network *MockNetwork) Relay(ctx context.Context, tag protocol.Tag, data []byte, wait bool, except network.Peer) error {
	return nil
}

// Address - unused function
func (network *MockNetwork) Address() (string, bool) {
	return "mock network", true
}

// Start - unused function
func (network *MockNetwork) Start() error {
	return nil
}

// Stop - unused function
func (network *MockNetwork) Stop() {
}

// RequestConnectOutgoing - unused function
func (network *MockNetwork) RequestConnectOutgoing(replace bool, quit <-chan struct{}) {
}

// Disconnect - unused function
func (network *MockNetwork) Disconnect(badpeer network.DisconnectablePeer) {
}

// DisconnectPeers - unused function
func (network *MockNetwork) DisconnectPeers() {
}

// RegisterRPCName - unused function
func (network *MockNetwork) RegisterRPCName(name string, rcvr interface{}) {
}

// GetPeers - unused function
func (network *MockNetwork) GetPeers(options ...network.PeerOption) []network.Peer {
	return nil
}

// Ready - always ready
func (network *MockNetwork) Ready() chan struct{} {
	c := make(chan struct{})
	close(c)
	return c
}

// RegisterHandlers - empty implementation.
func (network *MockNetwork) RegisterHandlers(dispatch []network.TaggedMessageHandler) {
}

// ClearHandlers - empty implementation
func (network *MockNetwork) ClearHandlers() {
}

// RegisterValidatorHandlers - empty implementation.
func (network *MockNetwork) RegisterValidatorHandlers(dispatch []network.TaggedMessageValidatorHandler) {
}

// ClearProcessors - empty implementation
func (network *MockNetwork) ClearProcessors() {
}

// RegisterHTTPHandler - empty implementation
func (network *MockNetwork) RegisterHTTPHandler(path string, handler http.Handler) {
}

// RegisterHTTPHandlerFunc - empty implementation
func (network *MockNetwork) RegisterHTTPHandlerFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
}

// OnNetworkAdvance - empty implementation
func (network *MockNetwork) OnNetworkAdvance() {}

// GetGenesisID - empty implementation
func (network *MockNetwork) GetGenesisID() string {
	if network.GenesisID == "" {
		return "mocknet"
	}
	return network.GenesisID
}

// GetHTTPClient returns a http.Client with a suitable for the network
func (network *MockNetwork) GetHTTPClient(address string) (*http.Client, error) {
	return nil, errors.New("not implemented")
}
