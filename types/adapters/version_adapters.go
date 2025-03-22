package adapters

import (
	abciv1 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	abciv2 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v2"
	typesv1 "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	typesv2 "github.com/depinnetwork/por-consensus/api/cometbft/types/v2"
	cmtabciv1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	abcitypes "github.com/depinnetwork/por-consensus/abci/types"
)

// ToV1ConsensusParams converts typesv2.ConsensusParams to typesv1.ConsensusParams
func ToV1ConsensusParams(params typesv2.ConsensusParams) typesv1.ConsensusParams {
	return typesv1.ConsensusParams{
		Block: typesv1.BlockParams{
			MaxBytes: params.Block.MaxBytes,
			MaxGas:   params.Block.MaxGas,
		},
		Evidence: typesv1.EvidenceParams{
			MaxAgeNumBlocks: params.Evidence.MaxAgeNumBlocks,
			MaxAgeDuration:  params.Evidence.MaxAgeDuration,
			MaxBytes:        params.Evidence.MaxBytes,
		},
		Validator: typesv1.ValidatorParams{
			PubKeyTypes: params.Validator.PubKeyTypes,
		},
	}
}

// ToV2ConsensusParams converts typesv1.ConsensusParams to typesv2.ConsensusParams
func ToV2ConsensusParams(params typesv1.ConsensusParams) typesv2.ConsensusParams {
	return typesv2.ConsensusParams{
		Block: typesv2.BlockParams{
			MaxBytes: params.Block.MaxBytes,
			MaxGas:   params.Block.MaxGas,
		},
		Evidence: typesv2.EvidenceParams{
			MaxAgeNumBlocks: params.Evidence.MaxAgeNumBlocks,
			MaxAgeDuration:  params.Evidence.MaxAgeDuration,
			MaxBytes:        params.Evidence.MaxBytes,
		},
		Validator: typesv2.ValidatorParams{
			PubKeyTypes: params.Validator.PubKeyTypes,
		},
	}
}

// V1ToV2Header converts typesv1.Header to typesv2.Header
func V1ToV2Header(header typesv1.Header) typesv2.Header {
	return typesv2.Header{
		Version: typesv2.Consensus{
			Block: header.Version.Block,
			App:   header.Version.App,
		},
		ChainID:            header.ChainID,
		Height:             header.Height,
		Time:               header.Time,
		LastBlockId:        V1ToV2BlockID(header.LastBlockId),
		LastCommitHash:     header.LastCommitHash,
		DataHash:           header.DataHash,
		ValidatorsHash:     header.ValidatorsHash,
		NextValidatorsHash: header.NextValidatorsHash,
		ConsensusHash:      header.ConsensusHash,
		AppHash:            header.AppHash,
		LastResultsHash:    header.LastResultsHash,
		EvidenceHash:       header.EvidenceHash,
		ProposerAddress:    header.ProposerAddress,
	}
}

// V2ToV1Header converts typesv2.Header to typesv1.Header
func V2ToV1Header(header typesv2.Header) typesv1.Header {
	return typesv1.Header{
		Version: typesv1.Consensus{
			Block: header.Version.Block,
			App:   header.Version.App,
		},
		ChainID:            header.ChainID,
		Height:             header.Height,
		Time:               header.Time,
		LastBlockId:        V2ToV1BlockID(header.LastBlockId),
		LastCommitHash:     header.LastCommitHash,
		DataHash:           header.DataHash,
		ValidatorsHash:     header.ValidatorsHash,
		NextValidatorsHash: header.NextValidatorsHash,
		ConsensusHash:      header.ConsensusHash,
		AppHash:            header.AppHash,
		LastResultsHash:    header.LastResultsHash,
		EvidenceHash:       header.EvidenceHash,
		ProposerAddress:    header.ProposerAddress,
	}
}

// V1ToV2BlockID converts typesv1.BlockID to typesv2.BlockID
func V1ToV2BlockID(blockID typesv1.BlockID) typesv2.BlockID {
	partSetHeader := typesv2.PartSetHeader{
		Total: blockID.PartSetHeader.Total,
		Hash:  blockID.PartSetHeader.Hash,
	}
	
	return typesv2.BlockID{
		Hash:          blockID.Hash,
		PartSetHeader: partSetHeader,
	}
}

// V2ToV1BlockID converts typesv2.BlockID to typesv1.BlockID
func V2ToV1BlockID(blockID typesv2.BlockID) typesv1.BlockID {
	partSetHeader := typesv1.PartSetHeader{
		Total: blockID.PartSetHeader.Total,
		Hash:  blockID.PartSetHeader.Hash,
	}
	
	return typesv1.BlockID{
		Hash:          blockID.Hash,
		PartSetHeader: partSetHeader,
	}
}

// ABCIV1ToV2Events converts []abciv1.Event to []abciv2.Event
func ABCIV1ToV2Events(events []abciv1.Event) []abciv2.Event {
	result := make([]abciv2.Event, len(events))
	for i, event := range events {
		attributes := make([]abciv2.EventAttribute, len(event.Attributes))
		for j, attr := range event.Attributes {
			attributes[j] = abciv2.EventAttribute{
				Key:   attr.Key,
				Value: attr.Value,
				Index: attr.Index,
			}
		}
		
		result[i] = abciv2.Event{
			Type:       event.Type,
			Attributes: attributes,
		}
	}
	return result
}

// ABCIV2ToV1Events converts []abciv2.Event to []abciv1.Event
func ABCIV2ToV1Events(events []abciv2.Event) []abciv1.Event {
	result := make([]abciv1.Event, len(events))
	for i, event := range events {
		attributes := make([]abciv1.EventAttribute, len(event.Attributes))
		for j, attr := range event.Attributes {
			attributes[j] = abciv1.EventAttribute{
				Key:   attr.Key,
				Value: attr.Value,
				Index: attr.Index,
			}
		}
		
		result[i] = abciv1.Event{
			Type:       event.Type,
			Attributes: attributes,
		}
	}
	return result
}

// CometBFTToDepinCommitResponse converts cometbft ABCI types to depin ABCI types
func CometBFTToDepinCommitResponse(response cmtabciv1.CommitResponse) abciv1.CommitResponse {
	return abciv1.CommitResponse{
		RetainHeight: response.RetainHeight,
	}
}

// ValidatorUpdatesAdapter converts between validator update types
func ValidatorUpdatesAdapter(updates []abciv1.ValidatorUpdate) []abcitypes.ValidatorUpdate {
	result := make([]abcitypes.ValidatorUpdate, len(updates))
	for i, update := range updates {
		result[i] = abcitypes.ValidatorUpdate{
			PubKey: abcitypes.PubKey{
				Type:  update.PubKey.Type,
				Data:  update.PubKey.Data,
			},
			Power:  update.Power,
		}
	}
	return result
}

// V1ToV2Block converts typesv1.Block to typesv2.Block
func V1ToV2Block(block *typesv1.Block) *typesv2.Block {
	if block == nil {
		return nil
	}
	
	header := V1ToV2Header(block.Header)
	data := typesv2.Data{
		Txs: block.Data.Txs,
	}
	
	var lastCommit *typesv2.Commit
	if block.LastCommit != nil {
		lastCommit = &typesv2.Commit{
			Height:     block.LastCommit.Height,
			Round:      block.LastCommit.Round,
			BlockID:    V1ToV2BlockID(block.LastCommit.BlockID),
			Signatures: make([]typesv2.CommitSig, len(block.LastCommit.Signatures)),
		}
		
		for i, sig := range block.LastCommit.Signatures {
			lastCommit.Signatures[i] = typesv2.CommitSig{
				BlockIDFlag:      typesv2.BlockIDFlag(sig.BlockIDFlag),
				ValidatorAddress: sig.ValidatorAddress,
				Timestamp:        sig.Timestamp,
				Signature:        sig.Signature,
			}
		}
	}
	
	var evidence typesv2.EvidenceList
	if block.Evidence != nil {
		evidence.Evidence = make([]typesv2.Evidence, len(block.Evidence.Evidence))
		// Evidence conversion would need to be implemented based on the evidence structure
	}
	
	return &typesv2.Block{
		Header:     header,
		Data:       data,
		Evidence:   evidence,
		LastCommit: lastCommit,
	}
}

// V2ToV1Block converts typesv2.Block to typesv1.Block
func V2ToV1Block(block *typesv2.Block) *typesv1.Block {
	if block == nil {
		return nil
	}
	
	header := V2ToV1Header(block.Header)
	data := typesv1.Data{
		Txs: block.Data.Txs,
	}
	
	var lastCommit *typesv1.Commit
	if block.LastCommit != nil {
		lastCommit = &typesv1.Commit{
			Height:     block.LastCommit.Height,
			Round:      block.LastCommit.Round,
			BlockID:    V2ToV1BlockID(block.LastCommit.BlockID),
			Signatures: make([]typesv1.CommitSig, len(block.LastCommit.Signatures)),
		}
		
		for i, sig := range block.LastCommit.Signatures {
			lastCommit.Signatures[i] = typesv1.CommitSig{
				BlockIDFlag:      typesv1.BlockIDFlag(sig.BlockIDFlag),
				ValidatorAddress: sig.ValidatorAddress,
				Timestamp:        sig.Timestamp,
				Signature:        sig.Signature,
			}
		}
	}
	
	var evidence typesv1.EvidenceList
	if block.Evidence.Evidence != nil {
		evidence.Evidence = make([]typesv1.Evidence, len(block.Evidence.Evidence))
		// Evidence conversion would need to be implemented based on the evidence structure
	}
	
	return &typesv1.Block{
		Header:     header,
		Data:       data,
		Evidence:   evidence,
		LastCommit: lastCommit,
	}
}
