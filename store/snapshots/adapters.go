package snapshots

import (
	cmtabciv1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	depinabciv1 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
)

// CMTToDepinSnapshot converts CometBFT snapshot to DePIN snapshot
func CMTToDepinSnapshot(snapshot *cmtabciv1.Snapshot) *depinabciv1.Snapshot {
	if snapshot == nil {
		return nil
	}
	
	return &depinabciv1.Snapshot{
		Height:   snapshot.Height,
		Format:   snapshot.Format,
		Chunks:   snapshot.Chunks,
		Hash:     snapshot.Hash,
		Metadata: snapshot.Metadata,
	}
}

// DepinToCMTSnapshot converts DePIN snapshot to CometBFT snapshot
func DepinToCMTSnapshot(snapshot *depinabciv1.Snapshot) *cmtabciv1.Snapshot {
	if snapshot == nil {
		return nil
	}
	
	return &cmtabciv1.Snapshot{
		Height:   snapshot.Height,
		Format:   snapshot.Format,
		Chunks:   snapshot.Chunks,
		Hash:     snapshot.Hash,
		Metadata: snapshot.Metadata,
	}
}

// We need an additional function to convert between snapshots for baseapp
func DepinSnapshotFromABCI(snapshot *depinabciv1.Snapshot) *depinabciv1.Snapshot {
	if snapshot == nil {
		return nil
	}
	
	return &depinabciv1.Snapshot{
		Height:   snapshot.Height,
		Format:   snapshot.Format,
		Chunks:   snapshot.Chunks,
		Hash:     snapshot.Hash,
		Metadata: snapshot.Metadata,
	}
}
