package events

import (
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
)

type ChainBridgeEvents struct {
	ChainBridge_FungibleTransfer        []EventFungibleTransfer        //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer     []EventNonFungibleTransfer     //nolint:stylecheck,golint
	ChainBridge_GenericTransfer         []EventGenericTransfer         //nolint:stylecheck,golint
	ChainBridge_RelayerThresholdChanged []EventRelayerThresholdChanged //nolint:stylecheck,golint
	ChainBridge_ChainWhitelisted        []EventChainWhitelisted        //nolint:stylecheck,golint
	ChainBridge_RelayerAdded            []EventRelayerAdded            //nolint:stylecheck,golint
	ChainBridge_RelayerRemoved          []EventRelayerRemoved          //nolint:stylecheck,golint
	ChainBridge_VoteFor                 []EventVoteFor                 //nolint:stylecheck,golint
	ChainBridge_VoteAgainst             []EventVoteAgainst             //nolint:stylecheck,golint
	ChainBridge_ProposalApproved        []EventProposalApproved        //nolint:stylecheck,golint
	ChainBridge_ProposalRejected        []EventProposalRejected        //nolint:stylecheck,golint
	ChainBridge_ProposalSucceeded       []EventProposalSucceeded       //nolint:stylecheck,golint
	ChainBridge_ProposalFailed          []EventProposalFailed          //nolint:stylecheck,golint
}

type BridgeTransferEvents struct {
	BridgeTransfer_ProposalFailed []EventProposalFailed //nolint:stylecheck,golint
}

type PhalaEvents struct {
	Phala_CommandPushed         []EventCommandPushed         //nolint:stylecheck,golint
	Phala_TransferToTee         []EventTransferToTee         //nolint:stylecheck,golint
	Phala_TransferToChain       []EventTransferToChain       //nolint:stylecheck,golint
	Phala_WorkerRegistered      []EventWorkerRegistered      //nolint:stylecheck,golint
	Phala_WorkerUnregistered    []EventWorkerUnregistered    //nolint:stylecheck,golint
	Phala_Heartbeat             []EventHeartbeat             //nolint:stylecheck,golint
	Phala_Offline               []EventOffline               //nolint:stylecheck,golint
	Phala_Slash                 []EventSlash                 //nolint:stylecheck,golint
	Phala_WorkerStateUpdated    []EventWorkerStateUpdated    //nolint:stylecheck,golint
	Phala_WhitelistAdded        []EventWhitelistAdded        //nolint:stylecheck,golint
	Phala_WhitelistRemoved      []EventWhitelistRemoved      //nolint:stylecheck,golint
	Phala_RewardSeed            []EventRewardSeed            //nolint:stylecheck,golint
	Phala_WorkerMessageReceived []EventWorkerMessageReceived //nolint:stylecheck,golint
	Phala_MinerStarted          []EventMinerStarted          //nolint:stylecheck,golint
	Phala_MinerStopped          []EventMinerStopped          //nolint:stylecheck,golint
	Phala_NewMiningRound        []EventNewMiningRound        //nolint:stylecheck,golint
	Phala_PayoutMissed          []EventPayoutMissed          //nolint:stylecheck,golint
	Phala_WorkerReset           []EventWorkerReset           //nolint:stylecheck,golint
	Phala_PayoutReward          []EventPayoutReward          //nolint:stylecheck,golint
}

type MiningStakingEvents struct {
	MiningStaking_PendingStakeApplied []EventPendingStakeApplied //nolint:stylecheck,golint
	MiningStaking_PendingUnstakeAdded []EventPendingUnstakeAdded //nolint:stylecheck,golint
	MiningStaking_PendingStakeAdded   []EventPendingStakeAdded   //nolint:stylecheck,golint
}

// pallet chain-bridge
type EventFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Amount       types.U256
	Recipient    types.Bytes
	Topics       []types.Hash
}

type EventNonFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	TokenId      types.Bytes
	Recipient    types.Bytes
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventGenericTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventRelayerThresholdChanged struct {
	Phase     types.Phase
	Threshold types.U32
	Topics    []types.Hash
}

type EventChainWhitelisted struct {
	Phase   types.Phase
	ChainId types.U8
	Topics  []types.Hash
}

type EventRelayerAdded struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventRelayerRemoved struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventVoteFor struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventVoteAgainst struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventProposalApproved struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalRejected struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalSucceeded struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

// pallet bridge-transfer
type EventProposalFailed struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventBridgeTransferRemark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// pallet phala
type EventCommandPushed struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.U32
	Arg2   types.Bytes
	Arg3   types.U64
	Topics []types.Hash
}

type EventTransferToTee struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.U128
	Topics []types.Hash
}

type EventTransferToChain struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.U128
	Arg2   types.U64
	Topics []types.Hash
}

type EventWorkerRegistered struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.Bytes
	Arg2   types.Bytes
	Topics []types.Hash
}

type EventWorkerUnregistered struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.Bytes
	Topics []types.Hash
}

type EventHeartbeat struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.U32
	Topics []types.Hash
}

type EventOffline struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Topics []types.Hash
}

type EventSlash struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.AccountID
	Arg2   types.U128
	Arg3   types.AccountID
	Arg4   types.U128
	Topics []types.Hash
}

type EventWorkerStateUpdated struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Topics []types.Hash
}

type EventWhitelistAdded struct {
	Phase  types.Phase
	Arg0   types.Bytes
	Topics []types.Hash
}

type EventWhitelistRemoved struct {
	Phase  types.Phase
	Arg0   types.Bytes
	Topics []types.Hash
}

type EventRewardSeed struct {
	Phase           types.Phase
	BlockRewardInfo BlockRewardInfo
	Topics          []types.Hash
}

type EventWorkerMessageReceived struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.Bytes
	Arg2   types.U64
	Topics []types.Hash
}

type EventMinerStarted struct {
	Phase  types.Phase
	Arg0   types.U32
	Arg1   types.AccountID
	Topics []types.Hash
}

type EventMinerStopped struct {
	Phase  types.Phase
	Arg0   types.U32
	Arg1   types.AccountID
	Topics []types.Hash
}

type EventNewMiningRound struct {
	Phase  types.Phase
	Arg0   types.U32
	Topics []types.Hash
}

type EventPayoutMissed struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.AccountID
	Topics []types.Hash
}

type EventWorkerReset struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.Bytes
	Topics []types.Hash
}

type EventPayoutReward struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.U128
	Arg2   types.U128
	Arg3   types.U8
	Topics []types.Hash
}

// pallet mining-staking
type EventPendingStakeApplied struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EventPendingUnstakeAdded struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.AccountID
	Arg2   types.U128
	Topics []types.Hash
}

type EventPendingStakeAdded struct {
	Phase  types.Phase
	Arg0   types.AccountID
	Arg1   types.AccountID
	Arg2   types.U128
	Topics []types.Hash
}

// pallet claim
// DEPRECATED
