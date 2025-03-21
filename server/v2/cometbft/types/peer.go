package types

import abci "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"

// PeerFilter responds to p2p filtering queries from Tendermint
type PeerFilter func(info string) (*abci.QueryResponse, error)
