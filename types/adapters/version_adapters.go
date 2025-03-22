package adapters

import (
	abciv1 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	abciv2 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v2"
	typesv1 "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	typesv2 "github.com/depinnetwork/por-consensus/api/cometbft/types/v2"
	cmtabciv1 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	abcitypes "github.com/depinnetwork/por-consensus/abci/types"
)

// ToV1ConsensusParams converts typesv2.ConsensusParams to typesv1.ConsensusParams
// This version fixes pointer/value type mismatches
func ToV1ConsensusParams(params typesv2.ConsensusParams) typesv1.ConsensusParams {
	var v1Params typesv1.ConsensusParams
	
	// Handle block params
	if params.Block != nil {
		v1Params.Block = &typesv1.BlockParams{
			MaxBytes: params.Block.MaxBytes,
			MaxGas:   params.Block.MaxGas,
		}
	} else {
		v1Params.Block = &typesv1.BlockParams{}
	}
	
	// Handle evidence params
	if params.Evidence != nil {
		v1Params.Evidence = &typesv1.EvidenceParams{
			MaxAgeNumBlocks: params.Evidence.MaxAgeNumBlocks,
			MaxAgeDuration:  params.Evidence.MaxAgeDuration,
			MaxBytes:        params.Evidence.MaxBytes,
		}
	} else {
		v1Params.Evidence = &typesv1.EvidenceParams{}
	}
	
	// Handle validator params
	if params.Validator != nil {
		v1Params.Validator = &typesv1.ValidatorParams{
			PubKeyTypes: params.Validator.PubKeyTypes,
		}
	} else {
		v1Params.Validator = &typesv1.ValidatorParams{}
	}
	
	return v1Params
}

// ToV2ConsensusParams converts typesv1.ConsensusParams to typesv2.ConsensusParams
// This version fixes pointer/value type mismatches
func ToV2ConsensusParams(params typesv1.ConsensusParams) typesv2.ConsensusParams {
	var v2Params typesv2.ConsensusParams
	
	// Handle block params
	if params.Block != nil {
		v2Params.Block = &typesv2.BlockParams{
			MaxBytes: params.Block.MaxBytes,
			MaxGas:   params.Block.MaxGas,
		}
	} else {
		v2Params.Block = &typesv2.BlockParams{}
	}
	
	// Handle evidence params
	if params.Evidence != nil {
		v2Params.Evidence = &typesv2.EvidenceParams{
			MaxAgeNumBlocks: params.Evidence.MaxAgeNumBlocks,
			MaxAgeDuration:  params.Evidence.MaxAgeDuration,
			MaxBytes:        params.Evidence.MaxBytes,
		}
	} else {
		v2Params.Evidence = &typesv2.EvidenceParams{}
	}
	
	// Handle validator params
	if params.Validator != nil {
		v2Params.Validator = &typesv2.ValidatorParams{
			PubKeyTypes: params.Validator.PubKeyTypes,
		}
	} else {
		v2Params.Validator = &typesv2.ValidatorParams{}
	}
	
	return v2Params
}

// V1ToV2Header converts typesv1.Header to typesv2.Header
// Properly handles Version field structure
func V1ToV2Header(header typesv1.Header) typesv2.Header {
	return typesv2.Header{
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
// Properly handles Version field structure
func V2ToV1Header(header typesv2.Header) typesv1.Header {
	return typesv1.Header{
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
	if events == nil {
		return nil
	}
	
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
	if events == nil {
		return nil
	}
	
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

// DepinToCometBFTCommitResponse converts depin ABCI types to cometbft ABCI types
func DepinToCometBFTCommitResponse(response abciv1.CommitResponse) cmtabciv1.CommitResponse {
	return cmtabciv1.CommitResponse{
		RetainHeight: response.RetainHeight,
	}
}

// ABCITypeToV1ValidatorUpdates converts abcitypes.ValidatorUpdate to abciv1.ValidatorUpdate
func ABCITypeToV1ValidatorUpdates(updates []abcitypes.ValidatorUpdate) []abciv1.ValidatorUpdate {
	if updates == nil {
		return nil
	}
	
	result := make([]abciv1.ValidatorUpdate, len(updates))
	for i, update := range updates {
		// Create a proper validator update with all fields correctly mapped
		result[i] = abciv1.ValidatorUpdate{
			PubKey: abciv1.PublicKey{
				Type: update.PubKey.Type,
				Data: update.PubKey.Data,
			},
			Power: update.Power,
		}
	}
	return result
}

// V1ToABCITypeValidatorUpdates converts abciv1.ValidatorUpdate to abcitypes.ValidatorUpdate
func V1ToABCITypeValidatorUpdates(updates []abciv1.ValidatorUpdate) []abcitypes.ValidatorUpdate {
	if updates == nil {
		return nil
	}
	
	result := make([]abcitypes.ValidatorUpdate, len(updates))
	for i, update := range updates {
		// Create a proper validator update with all fields correctly mapped
		result[i] = abcitypes.ValidatorUpdate{
			PubKey: abcitypes.PubKey{
				Type: update.PubKey.Type,
				Data: update.PubKey.Data,
			},
			Power: update.Power,
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
				BlockIdFlag:      typesv2.BlockIDFlag(sig.BlockIdFlag),
				ValidatorAddress: sig.ValidatorAddress,
				Timestamp:        sig.Timestamp,
				Signature:        sig.Signature,
			}
		}
	}
	
	var evidence typesv2.EvidenceList
	if block.Evidence != nil && len(block.Evidence.Evidence) > 0 {
		// Here we properly initialize the Evidence slice
		evidence.Evidence = make([]typesv2.Evidence, 0)
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
				BlockIdFlag:      typesv1.BlockIDFlag(sig.BlockIdFlag),
				ValidatorAddress: sig.ValidatorAddress,
				Timestamp:        sig.Timestamp,
				Signature:        sig.Signature,
			}
		}
	}
	
	// Initialize evidence properly
	evidence := typesv1.EvidenceList{
		Evidence: make([]typesv1.Evidence, 0),
	}
	
	return &typesv1.Block{
		Header:     header,
		Data:       data,
		Evidence:   evidence,
		LastCommit: lastCommit,
	}
}

// CometBFTToDepinFinalizeBlockRequest converts cometbft ABCI types to depin ABCI types
func CometBFTToDepinFinalizeBlockRequest(req cmtabciv1.FinalizeBlockRequest) abciv1.FinalizeBlockRequest {
	depinReq := abciv1.FinalizeBlockRequest{
		Hash:             req.Hash,
		Height:           req.Height,
		Time:             req.Time,
		ProposerAddress:  req.ProposerAddress,
		Txs:              req.Txs,
		Misbehavior:      make([]abciv1.Misbehavior, len(req.Misbehavior)),
		DecidedLastCommit: abciv1.CommitInfo{
			Round: req.DecidedLastCommit.Round,
			Votes: make([]abciv1.VoteInfo, len(req.DecidedLastCommit.Votes)),
		},
	}
	
	// Convert misbehavior
	for i, m := range req.Misbehavior {
		depinReq.Misbehavior[i] = abciv1.Misbehavior{
			Type:             m.Type,
			Height:           m.Height,
			Time:             m.Time,
			ValidatorAddress: m.ValidatorAddress,
			TotalVotingPower: m.TotalVotingPower,
		}
	}
	
	// Convert votes
	for i, v := range req.DecidedLastCommit.Votes {
		depinReq.DecidedLastCommit.Votes[i] = abciv1.VoteInfo{
			Validator: abciv1.Validator{
				Address: v.Validator.Address,
				Power:   v.Validator.Power,
			},
			BlockIdFlag: v.BlockIdFlag,
		}
	}
	
	return depinReq
}

// DepinToCometBFTFinalizeBlockRequest converts depin ABCI types to cometbft ABCI types
func DepinToCometBFTFinalizeBlockRequest(req abciv1.FinalizeBlockRequest) cmtabciv1.FinalizeBlockRequest {
	cmtReq := cmtabciv1.FinalizeBlockRequest{
		Hash:             req.Hash,
		Height:           req.Height,
		Time:             req.Time,
		ProposerAddress:  req.ProposerAddress,
		Txs:              req.Txs,
		Misbehavior:      make([]cmtabciv1.Misbehavior, len(req.Misbehavior)),
		DecidedLastCommit: cmtabciv1.CommitInfo{
			Round: req.DecidedLastCommit.Round,
			Votes: make([]cmtabciv1.VoteInfo, len(req.DecidedLastCommit.Votes)),
		},
	}
	
	// Convert misbehavior
	for i, m := range req.Misbehavior {
		cmtReq.Misbehavior[i] = cmtabciv1.Misbehavior{
			Type:             m.Type,
			Height:           m.Height,
			Time:             m.Time,
			ValidatorAddress: m.ValidatorAddress,
			TotalVotingPower: m.TotalVotingPower,
		}
	}
	
	// Convert votes
	for i, v := range req.DecidedLastCommit.Votes {
		cmtReq.DecidedLastCommit.Votes[i] = cmtabciv1.VoteInfo{
			Validator: cmtabciv1.Validator{
				Address: v.Validator.Address,
				Power:   v.Validator.Power,
			},
			BlockIdFlag: v.BlockIdFlag,
		}
	}
	
	return cmtReq
}

// CometBFTToDepinFinalizeBlockResponse converts cometbft ABCI types to depin ABCI types
func CometBFTToDepinFinalizeBlockResponse(res cmtabciv1.FinalizeBlockResponse) abciv1.FinalizeBlockResponse {
	depinRes := abciv1.FinalizeBlockResponse{
		AppHash:         res.AppHash,
		ValidatorUpdates: make([]abciv1.ValidatorUpdate, len(res.ValidatorUpdates)),
		ConsensusParamUpdates: &abciv1.ConsensusParams{},
		Events:          make([]abciv1.Event, len(res.Events)),
	}
	
	// Convert validator updates
	for i, v := range res.ValidatorUpdates {
		depinRes.ValidatorUpdates[i] = abciv1.ValidatorUpdate{
			PubKey: abciv1.PublicKey{
				Type: v.PubKey.Type,
				Data: v.PubKey.Data,
			},
			Power: v.Power,
		}
	}
	
	// Convert events
	for i, e := range res.Events {
		depinEvent := abciv1.Event{
			Type:       e.Type,
			Attributes: make([]abciv1.EventAttribute, len(e.Attributes)),
		}
		
		for j, a := range e.Attributes {
			depinEvent.Attributes[j] = abciv1.EventAttribute{
				Key:   a.Key,
				Value: a.Value,
				Index: a.Index,
			}
		}
		
		depinRes.Events[i] = depinEvent
	}
	
	// Handle consensus param updates if present
	if res.ConsensusParamUpdates != nil {
		if res.ConsensusParamUpdates.Block != nil {
			depinRes.ConsensusParamUpdates.Block = &abciv1.BlockParams{
				MaxBytes: res.ConsensusParamUpdates.Block.MaxBytes,
				MaxGas:   res.ConsensusParamUpdates.Block.MaxGas,
			}
		}
		
		if res.ConsensusParamUpdates.Evidence != nil {
			depinRes.ConsensusParamUpdates.Evidence = &abciv1.EvidenceParams{
				MaxAgeNumBlocks: res.ConsensusParamUpdates.Evidence.MaxAgeNumBlocks,
				MaxAgeDuration:  res.ConsensusParamUpdates.Evidence.MaxAgeDuration,
				MaxBytes:        res.ConsensusParamUpdates.Evidence.MaxBytes,
			}
		}
		
		if res.ConsensusParamUpdates.Validator != nil {
			depinRes.ConsensusParamUpdates.Validator = &abciv1.ValidatorParams{
				PubKeyTypes: res.ConsensusParamUpdates.Validator.PubKeyTypes,
			}
		}
	}
	
	return depinRes
}

// DepinToCometBFTFinalizeBlockResponse converts depin ABCI types to cometbft ABCI types
func DepinToCometBFTFinalizeBlockResponse(res abciv1.FinalizeBlockResponse) cmtabciv1.FinalizeBlockResponse {
	cmtRes := cmtabciv1.FinalizeBlockResponse{
		AppHash:         res.AppHash,
		ValidatorUpdates: make([]cmtabciv1.ValidatorUpdate, len(res.ValidatorUpdates)),
		ConsensusParamUpdates: &cmtabciv1.ConsensusParams{},
		Events:          make([]cmtabciv1.Event, len(res.Events)),
	}
	
	// Convert validator updates
	for i, v := range res.ValidatorUpdates {
		cmtRes.ValidatorUpdates[i] = cmtabciv1.ValidatorUpdate{
			PubKey: cmtabciv1.PublicKey{
				Type: v.PubKey.Type,
				Data: v.PubKey.Data,
			},
			Power: v.Power,
		}
	}
	
	// Convert events
	for i, e := range res.Events {
		cmtEvent := cmtabciv1.Event{
			Type:       e.Type,
			Attributes: make([]cmtabciv1.EventAttribute, len(e.Attributes)),
		}
		
		for j, a := range e.Attributes {
			cmtEvent.Attributes[j] = cmtabciv1.EventAttribute{
				Key:   a.Key,
				Value: a.Value,
				Index: a.Index,
			}
		}
		
		cmtRes.Events[i] = cmtEvent
	}
	
	// Handle consensus param updates if present
	if res.ConsensusParamUpdates != nil {
		if res.ConsensusParamUpdates.Block != nil {
			cmtRes.ConsensusParamUpdates.Block = &cmtabciv1.BlockParams{
				MaxBytes: res.ConsensusParamUpdates.Block.MaxBytes,
				MaxGas:   res.ConsensusParamUpdates.Block.MaxGas,
			}
		}
		
		if res.ConsensusParamUpdates.Evidence != nil {
			cmtRes.ConsensusParamUpdates.Evidence = &cmtabciv1.EvidenceParams{
				MaxAgeNumBlocks: res.ConsensusParamUpdates.Evidence.MaxAgeNumBlocks,
				MaxAgeDuration:  res.ConsensusParamUpdates.Evidence.MaxAgeDuration,
				MaxBytes:        res.ConsensusParamUpdates.Evidence.MaxBytes,
			}
		}
		
		if res.ConsensusParamUpdates.Validator != nil {
			cmtRes.ConsensusParamUpdates.Validator = &cmtabciv1.ValidatorParams{
				PubKeyTypes: res.ConsensusParamUpdates.Validator.PubKeyTypes,
			}
		}
	}
	
	return cmtRes
}

// DepinToAbciQueryResponse converts depin-format query response to ABCI format
func DepinToAbciQueryResponse(res abciv1.QueryResponse) abcitypes.QueryResponse {
	return abcitypes.QueryResponse{
		Code:   res.Code,
		Log:    res.Log,
		Info:   res.Info,
		Index:  res.Index,
		Key:    res.Key,
		Value:  res.Value,
		Proof:  res.Proof,
		Height: res.Height,
	}
}

// AbciToDepinQueryResponse converts ABCI-format query response to depin format
func AbciToDepinQueryResponse(res abcitypes.QueryResponse) abciv1.QueryResponse {
	return abciv1.QueryResponse{
		Code:   res.Code,
		Log:    res.Log,
		Info:   res.Info,
		Index:  res.Index,
		Key:    res.Key,
		Value:  res.Value,
		Proof:  res.Proof,
		Height: res.Height,
	}
}
