package keeper

import (
	"context"
	"encoding/hex"

	cosmoserrors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/observer/types"
)

// AddBlockHeader handles adding a block header to the store, through majority voting of observers
func (k msgServer) AddBlockHeader(goCtx context.Context, msg *types.MsgAddBlockHeader) (*types.MsgAddBlockHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check authorization for this chain
	chain := common.GetChainFromChainID(msg.ChainId)
	ok, err := k.IsAuthorized(ctx, msg.Creator, chain)
	if !ok {
		return nil, cosmoserrors.Wrap(types.ErrNotAuthorizedPolicy, err.Error())
	}

	// add vote to ballot
	ballot, _, err := k.FindBallot(ctx, msg.Digest(), chain, types.ObservationType_InBoundTx)
	if err != nil {
		return nil, cosmoserrors.Wrap(err, "failed to find ballot")
	}
	ballot, err = k.AddVoteToBallot(ctx, ballot, msg.Creator, types.VoteType_SuccessObservation)
	if err != nil {
		return nil, cosmoserrors.Wrap(err, "failed to add vote to ballot")
	}
	_, isFinalized := k.CheckIfFinalizingVote(ctx, ballot)
	if !isFinalized {
		return &types.MsgAddBlockHeaderResponse{}, nil
	}

	/**
	 * Vote finalized, add block header to store
	 */
	_, found := k.GetBlockHeader(ctx, msg.BlockHash)
	if found {
		return nil, cosmoserrors.Wrap(types.ErrBlockAlreadyExist, hex.EncodeToString(msg.BlockHash))
	}

	// NOTE: error is checked in BasicValidation in msg; check again for extra caution
	pHash, err := msg.Header.ParentHash()
	if err != nil {
		return nil, cosmoserrors.Wrap(types.ErrNoParentHash, err.Error())
	}

	// TODO: add check for parent block header's existence here https://github.com/zeta-chain/node/issues/1133

	bh := common.BlockHeader{
		Header:     msg.Header,
		Height:     msg.Height,
		Hash:       msg.BlockHash,
		ParentHash: pHash,
		ChainId:    msg.ChainId,
	}
	k.SetBlockHeader(ctx, bh)

	return &types.MsgAddBlockHeaderResponse{}, nil
}
