package cometbft

import (
	"context"
	"strings"

	abci "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	crypto "github.com/depinnetwork/por-consensus/api/cometbft/crypto/v1"

	errorsmod "cosmossdk.io/errors/v2"
	cometerrors "cosmossdk.io/server/v2/cometbft/types/errors"
)

func (c *consensus[T]) handleQueryP2P(path []string) (*abci.QueryResponse, error) {
	// "/p2p" prefix for p2p queries
	if len(path) < 4 {
		return nil, errorsmod.Wrap(cometerrors.ErrUnknownRequest, "path should be p2p filter <addr|id> <parameter>")
	}

	cmd, typ, arg := path[1], path[2], path[3]
	if cmd == "filter" {
		switch typ {
		case "addr":
			if c.addrPeerFilter != nil {
				return c.addrPeerFilter(arg)
			}
		case "id":
			if c.idPeerFilter != nil {
				return c.idPeerFilter(arg)
			}
		}
		return &abci.QueryResponse{}, nil
	}

	return nil, errorsmod.Wrap(cometerrors.ErrUnknownRequest, "expected second parameter to be 'filter'")
}

// handleQueryApp handles the query requests for the application.
// It expects the path parameter to have at least two elements.
// The second element of the path can be either 'simulate' or 'version'.
// If the second element is 'simulate', it decodes the request data into a transaction,
// simulates the transaction using the application, and returns the simulation result.
// If the second element is 'version', it returns the version of the application.
// If the second element is neither 'simulate' nor 'version', it returns an error indicating an unknown query.
func (c *consensus[T]) handleQueryApp(ctx context.Context, path []string, req *abci.QueryRequest) (*abci.QueryResponse, error) {
	if len(path) < 2 {
		return nil, errorsmod.Wrap(
			cometerrors.ErrUnknownRequest,
			"expected second parameter to be either 'simulate' or 'version', neither was present",
		)
	}

	switch path[1] {
	case "simulate":
		tx, err := c.appCodecs.TxCodec.Decode(req.Data)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to decode tx")
		}

		txResult, _, err := c.app.Simulate(ctx, tx)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to simulate tx")
		}

		bz, err := intoABCISimulationResponse(
			txResult,
			c.indexedABCIEvents,
			c.cfg.AppTomlConfig.DisableIndexABCIEvents,
		)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to marshal txResult")
		}

		return &abci.QueryResponse{
			Codespace: cometerrors.RootCodespace,
			Value:     bz,
			Height:    req.Height,
		}, nil

	case "version":
		return &abci.QueryResponse{
			Codespace: cometerrors.RootCodespace,
			Value:     []byte(c.version),
			Height:    req.Height,
		}, nil
	}

	return nil, errorsmod.Wrapf(cometerrors.ErrUnknownRequest, "unknown query: %s", path)
}

func (c *consensus[T]) handleQueryStore(path []string, req *abci.QueryRequest) (*abci.QueryResponse, error) {
	req.Path = "/" + strings.Join(path[1:], "/")
	if req.Height <= 1 && req.Prove {
		return nil, errorsmod.Wrap(
			cometerrors.ErrInvalidRequest,
			"cannot query with proof when height <= 1; please provide a valid height",
		)
	}

	// "/store/<storeName>" for store queries
	storeName := path[1]
	storeNameBz := []byte(storeName) // TODO fastpath?
	qRes, err := c.store.Query(storeNameBz, uint64(req.Height), req.Data, req.Prove)
	if err != nil {
		return nil, err
	}

	res := &abci.QueryResponse{
		Codespace: cometerrors.RootCodespace,
		Height:    int64(qRes.Version),
		Key:       qRes.Key,
		Value:     qRes.Value,
	}

	if req.Prove {
		for _, proof := range qRes.ProofOps {
			bz, err := proof.Proof.Marshal()
			if err != nil {
				return nil, errorsmod.Wrap(err, "failed to marshal proof")
			}

			res.ProofOps = &crypto.ProofOps{
				Ops: []crypto.ProofOp{
					{
						Type: proof.Type,
						Key:  proof.Key,
						Data: bz,
					},
				},
			}
		}
	}

	return res, nil
}
