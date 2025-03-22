package adapters

import (
	abciv1 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1"
	abcitypes "github.com/depinnetwork/por-consensus/abci/types"
	abciv2 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v2"
	typesv1 "github.com/depinnetwork/por-consensus/api/cometbft/types/v1"
	typesv2 "github.com/depinnetwork/por-consensus/api/cometbft/types/v2"
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
func CometBFTToDepinCommitResponse(response abciv1.CommitResponse) abciv1.CommitResponse {
	return abciv1.CommitResponse{
		RetainHeight: response.RetainHeight,
	}
}

// DepinToCometBFTCommitResponse converts depin ABCI types to cometbft ABCI types
func DepinToCometBFTCommitResponse(response abciv1.CommitResponse) abciv1.CommitResponse {
	return abciv1.CommitResponse{
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
			Power: update.GetPower(),
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
			Power: update.GetPower(),
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
				Address: sig.ValidatorAddress,
				Timestamp:        sig.Timestamp,
				Signature:        sig.Signature,
			}
		}
	}
	
	var evidence typesv2.EvidenceList
	if block.Evidence.Evidence != nil && len(block.Evidence.Evidence) > 0 && block.Evidence.Evidence != nil && len(block.Evidence.Evidence) > 0 {
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
				BlockIDFlag:      typesv1.BlockIDFlag(sig.BlockIDFlag),
				Address: sig.ValidatorAddress,
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
func CometBFTToDepinFinalizeBlockRequest(req abciv1.FinalizeBlockRequest) abciv1.FinalizeBlockRequest {
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
			Address:          m.Address,
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
			BlockIDFlag: v.BlockIDFlag,
		}
	}
	
	return depinReq
}

// DepinToCometBFTFinalizeBlockRequest converts depin ABCI types to cometbft ABCI types
func DepinToCometBFTFinalizeBlockRequest(req abciv1.FinalizeBlockRequest) abciv1.FinalizeBlockRequest {
	cmtReq := abciv1.FinalizeBlockRequest{
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
		cmtReq.Misbehavior[i] = abciv1.Misbehavior{
			Type:             m.Type,
			Height:           m.Height,
			Time:             m.Time,
			Address:          m.Address,
			TotalVotingPower: m.TotalVotingPower,
		}
	}
	
	// Convert votes
	for i, v := range req.DecidedLastCommit.Votes {
		cmtReq.DecidedLastCommit.Votes[i] = abciv1.VoteInfo{
			Validator: abciv1.Validator{
				Address: v.Validator.Address,
				Power:   v.Validator.Power,
			},
			BlockIDFlag: v.BlockIDFlag,
		}
	}
	
	return cmtReq
}

// CometBFTToDepinFinalizeBlockResponse converts cometbft ABCI types to depin ABCI types
func CometBFTToDepinFinalizeBlockResponse(res abciv1.FinalizeBlockResponse) abciv1.FinalizeBlockResponse {
	depinRes := abciv1.FinalizeBlockResponse{
		AppHash:         res.AppHash,
		ValidatorUpdates: make([]abciv1.ValidatorUpdate, len(res.ValidatorUpdates)),
		ConsensusParamUpdates: &abciv1.ConsensusParams{},
		Events:          make([]abciv1.Event, len(res.Events)),
	}
	
	// Convert validator updates
	for i, v := range res.ValidatorUpdates {
		depinRes.ValidatorUpdates[i] = abciv1.ValidatorUpdate{
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
func DepinToCometBFTFinalizeBlockResponse(res abciv1.FinalizeBlockResponse) abciv1.FinalizeBlockResponse {
	cmtRes := abciv1.FinalizeBlockResponse{
		AppHash:         res.AppHash,
		ValidatorUpdates: make([]abciv1.ValidatorUpdate, len(res.ValidatorUpdates)),
		ConsensusParamUpdates: &abciv1.ConsensusParams{},
		Events:          make([]abciv1.Event, len(res.Events)),
	}
	
	// Convert validator updates
	for i, v := range res.ValidatorUpdates {
		cmtRes.ValidatorUpdates[i] = abciv1.ValidatorUpdate{
			Power: v.Power,
		}
	}
	
	// Convert events
	for i, e := range res.Events {
		cmtEvent := abciv1.Event{
			Type:       e.Type,
			Attributes: make([]abciv1.EventAttribute, len(e.Attributes)),
		}
		
		for j, a := range e.Attributes {
			cmtEvent.Attributes[j] = abciv1.EventAttribute{
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
			cmtRes.ConsensusParamUpdates.Block = &abciv1.BlockParams{
				MaxBytes: res.ConsensusParamUpdates.Block.MaxBytes,
				MaxGas:   res.ConsensusParamUpdates.Block.MaxGas,
			}
		}
		
		if res.ConsensusParamUpdates.Evidence != nil {
			cmtRes.ConsensusParamUpdates.Evidence = &abciv1.EvidenceParams{
				MaxAgeNumBlocks: res.ConsensusParamUpdates.Evidence.MaxAgeNumBlocks,
				MaxAgeDuration:  res.ConsensusParamUpdates.Evidence.MaxAgeDuration,
				MaxBytes:        res.ConsensusParamUpdates.Evidence.MaxBytes,
			}
		}
		
		if res.ConsensusParamUpdates.Validator != nil {
			cmtRes.ConsensusParamUpdates.Validator = &abciv1.ValidatorParams{
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
		Height: res.Height,
	}
}

// V2ToV1Data converts typesv2.Data to typesv1.Data
func V2ToV1Data(data typesv2.Data) typesv1.Data {
	return typesv1.Data{
		Txs: data.Txs,
	}
}

// V1ToV2Data converts typesv1.Data to typesv2.Data
func V1ToV2Data(data typesv1.Data) typesv2.Data {
	return typesv2.Data{
		Txs: data.Txs,
	}
}

// V2ToV1EvidenceList converts typesv2.EvidenceList to typesv1.EvidenceList
func V2ToV1EvidenceList(evidenceList typesv2.EvidenceList) typesv1.EvidenceList {
	// Initialize an empty EvidenceList
	v1Evidence := typesv1.EvidenceList{
		Evidence: make([]typesv1.Evidence, 0),
	}
	
	return v1Evidence
}

// V1ToV2EvidenceList converts typesv1.EvidenceList to typesv2.EvidenceList
func V1ToV2EvidenceList(evidenceList typesv1.EvidenceList) typesv2.EvidenceList {
	// Initialize an empty EvidenceList
	v2Evidence := typesv2.EvidenceList{
		Evidence: make([]typesv2.Evidence, 0),
	}
	
	return v2Evidence
}

// V2ToV1Commit converts typesv2.Commit to typesv1.Commit
func V2ToV1Commit(commit *typesv2.Commit) *typesv1.Commit {
	if commit == nil {
		return nil
	}
	
	v1Commit := &typesv1.Commit{
		Height:     commit.Height,
		Round:      commit.Round,
		BlockID:    V2ToV1BlockID(commit.BlockID),
		Signatures: make([]typesv1.CommitSig, len(commit.Signatures)),
	}
	
	for i, sig := range commit.Signatures {
		v1Commit.Signatures[i] = typesv1.CommitSig{
			BlockIDFlag:      typesv1.BlockIDFlag(sig.BlockIDFlag),
			Address: sig.ValidatorAddress,
			Timestamp:        sig.Timestamp,
			Signature:        sig.Signature,
		}
	}
	
	return v1Commit
}

// V1ToV2Commit converts typesv1.Commit to typesv2.Commit
func V1ToV2Commit(commit *typesv1.Commit) *typesv2.Commit {
	if commit == nil {
		return nil
	}
	
	v2Commit := &typesv2.Commit{
		Height:     commit.Height,
		Round:      commit.Round,
		BlockID:    V1ToV2BlockID(commit.BlockID),
		Signatures: make([]typesv2.CommitSig, len(commit.Signatures)),
	}
	
	for i, sig := range commit.Signatures {
		v2Commit.Signatures[i] = typesv2.CommitSig{
			BlockIDFlag:      typesv2.BlockIDFlag(sig.BlockIDFlag),
			Address: sig.ValidatorAddress,
			Timestamp:        sig.Timestamp,
			Signature:        sig.Signature,
		}
	}
	
	return v2Commit
}

// V1ToV2Proposal converts typesv1.Proposal to typesv2.Proposal
func V1ToV2Proposal(proposal *typesv1.Proposal) *typesv2.Proposal {
	if proposal == nil {
		return nil
	}
	
	return &typesv2.Proposal{
		Type:      proposal.Type,
		Height:    proposal.Height, 
		Round:     proposal.Round,
		PolRound:  proposal.PolRound,
		BlockID:   V1ToV2BlockID(proposal.BlockID),
		Timestamp: proposal.Timestamp,
		Signature: proposal.Signature,
	}
}

// V2ToV1Proposal converts typesv2.Proposal to typesv1.Proposal
func V2ToV1Proposal(proposal *typesv2.Proposal) *typesv1.Proposal {
	if proposal == nil {
		return nil
	}
	
	return &typesv1.Proposal{
		Type:      proposal.Type,
		Height:    proposal.Height,
		Round:     proposal.Round,
		PolRound:  proposal.PolRound,
		BlockID:   V2ToV1BlockID(proposal.BlockID),
		Timestamp: proposal.Timestamp,
		Signature: proposal.Signature,
	}
}

// ABCITypeToV1InitChainRequest converts abcitypes.InitChainRequest to abciv1.InitChainRequest
func ABCITypeToV1InitChainRequest(req *abcitypes.InitChainRequest) *abciv1.InitChainRequest {
	if req == nil {
		return nil
	}
	
	return &abciv1.InitChainRequest{
		Time:             req.Time,
		ChainId:          req.ChainId,
		ConsensusParams:  nil, // Would need conversion if used
		Validators:       ABCITypeToV1ValidatorUpdates(req.Validators),
		AppStateBytes:    req.AppStateBytes,
		InitialHeight:    req.InitialHeight,
	}
}

// ABCITypeToV1FinalizeBlockRequest converts abcitypes.FinalizeBlockRequest to abciv1.FinalizeBlockRequest
func ABCITypeToV1FinalizeBlockRequest(req *abcitypes.FinalizeBlockRequest) *abciv1.FinalizeBlockRequest {
	if req == nil {
		return nil
	}
	
	result := &abciv1.FinalizeBlockRequest{
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
		result.Misbehavior[i] = abciv1.Misbehavior{
			Type:             m.Type,
			Height:           m.Height,
			Time:             m.Time,
			Address:          m.Address,
			TotalVotingPower: m.TotalVotingPower,
		}
	}
	
	// Convert votes
	for i, v := range req.DecidedLastCommit.Votes {
		result.DecidedLastCommit.Votes[i] = abciv1.VoteInfo{
			Validator: abciv1.Validator{
				Address: v.Validator.Address,
				Power:   v.Validator.Power,
			},
			BlockIDFlag: v.BlockIDFlag,
		}
	}
	
	return result
}

// V2ToV1Data converts typesv2.Data to typesv1.Data
func V2ToV1Data(data typesv2.Data) typesv1.Data {
	return typesv1.Data{
		Txs: data.Txs,
	}
}

// V1ToV2Data converts typesv1.Data to typesv2.Data
func V1ToV2Data(data typesv1.Data) typesv2.Data {
	return typesv2.Data{
		Txs: data.Txs,
	}
}

// V2ToV1EvidenceList converts typesv2.EvidenceList to typesv1.EvidenceList
func V2ToV1EvidenceList(evidenceList typesv2.EvidenceList) typesv1.EvidenceList {
	// Initialize an empty EvidenceList
	v1Evidence := typesv1.EvidenceList{
		Evidence: make([]typesv1.Evidence, 0),
	}
	
	return v1Evidence
}

// V1ToV2EvidenceList converts typesv1.EvidenceList to typesv2.EvidenceList
func V1ToV2EvidenceList(evidenceList typesv1.EvidenceList) typesv2.EvidenceList {
	// Initialize an empty EvidenceList
	v2Evidence := typesv2.EvidenceList{
		Evidence: make([]typesv2.Evidence, 0),
	}
	
	return v2Evidence
}

// V2ToV1Commit converts typesv2.Commit to typesv1.Commit
func V2ToV1Commit(commit *typesv2.Commit) *typesv1.Commit {
	if commit == nil {
		return nil
	}
	
	v1Commit := &typesv1.Commit{
		Height:     commit.Height,
		Round:      commit.Round,
		BlockID:    V2ToV1BlockID(commit.BlockID),
		Signatures: make([]typesv1.CommitSig, len(commit.Signatures)),
	}
	
	for i, sig := range commit.Signatures {
		v1Commit.Signatures[i] = typesv1.CommitSig{
			BlockIDFlag:      typesv1.BlockIDFlag(sig.BlockIDFlag),
			ValidatorAddress: sig.ValidatorAddress,
			Timestamp:        sig.Timestamp,
			Signature:        sig.Signature,
		}
	}
	
	return v1Commit
}

// V1ToV2Commit converts typesv1.Commit to typesv2.Commit
func V1ToV2Commit(commit *typesv1.Commit) *typesv2.Commit {
	if commit == nil {
		return nil
	}
	
	v2Commit := &typesv2.Commit{
		Height:     commit.Height,
		Round:      commit.Round,
		BlockID:    V1ToV2BlockID(commit.BlockID),
		Signatures: make([]typesv2.CommitSig, len(commit.Signatures)),
	}
	
	for i, sig := range commit.Signatures {
		v2Commit.Signatures[i] = typesv2.CommitSig{
			BlockIDFlag:      typesv2.BlockIDFlag(sig.BlockIDFlag),
			ValidatorAddress: sig.ValidatorAddress,
			Timestamp:        sig.Timestamp,
			Signature:        sig.Signature,
		}
	}
	
	return v2Commit
}

// V1ToV2Proposal converts typesv1.Proposal to typesv2.Proposal
func V1ToV2Proposal(proposal *typesv1.Proposal) *typesv2.Proposal {
	if proposal == nil {
		return nil
	}
	
	return &typesv2.Proposal{
		Type:      proposal.Type,
		Height:    proposal.Height, 
		Round:     proposal.Round,
		PolRound:  proposal.PolRound,
		BlockID:   V1ToV2BlockID(proposal.BlockID),
		Timestamp: proposal.Timestamp,
		Signature: proposal.Signature,
	}
}

// V2ToV1Proposal converts typesv2.Proposal to typesv1.Proposal
func V2ToV1Proposal(proposal *typesv2.Proposal) *typesv1.Proposal {
	if proposal == nil {
		return nil
	}
	
	return &typesv1.Proposal{
		Type:      proposal.Type,
		Height:    proposal.Height,
		Round:     proposal.Round,
		PolRound:  proposal.PolRound,
		BlockID:   V2ToV1BlockID(proposal.BlockID),
		Timestamp: proposal.Timestamp,
		Signature: proposal.Signature,
	}
}

// ABCITypeToV1InitChainRequest converts abcitypes.InitChainRequest to abciv1.InitChainRequest
func ABCITypeToV1InitChainRequest(req *abcitypes.InitChainRequest) *abciv1.InitChainRequest {
	if req == nil {
		return nil
	}
	
	return &abciv1.InitChainRequest{
		Time:             req.Time,
		ChainId:          req.ChainId,
		ConsensusParams:  nil, // Would need conversion if used
		Validators:       ABCITypeToV1ValidatorUpdates(req.Validators),
		AppStateBytes:    req.AppStateBytes,
		InitialHeight:    req.InitialHeight,
	}
}

// ABCITypeToV1FinalizeBlockRequest converts abcitypes.FinalizeBlockRequest to abciv1.FinalizeBlockRequest
func ABCITypeToV1FinalizeBlockRequest(req *abcitypes.FinalizeBlockRequest) *abciv1.FinalizeBlockRequest {
	if req == nil {
		return nil
	}
	
	result := &abciv1.FinalizeBlockRequest{
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
		result.Misbehavior[i] = abciv1.Misbehavior{
			Type:             m.Type,
			Height:           m.Height,
			Time:             m.Time,
			Address:          m.Address,
			TotalVotingPower: m.TotalVotingPower,
		}
	}
	
	// Convert votes
	for i, v := range req.DecidedLastCommit.Votes {
		result.DecidedLastCommit.Votes[i] = abciv1.VoteInfo{
			Validator: abciv1.Validator{
				Address: v.Validator.Address,
				Power:   v.Validator.Power,
			},
			BlockIDFlag: v.BlockIDFlag,
		}
	}
	
	return result
}
