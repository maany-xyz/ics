package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/maany-xyz/ics/v5/x/ccv/provider/consumergov/types"
)

type msgServer struct {
	*Keeper
}

func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) ConsumerGovProposal(goCtx context.Context, msg *types.MsgConsumerGovProposal) (*types.MsgConsumerGovProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Add your governance proposal execution logic here
	k.Logger(ctx).Info("Processing ConsumerGovProposal",
		"title", msg.Title,
		"consumer_id", msg.ConsumerId,
		"action", msg.Action,
	)

	return &types.MsgConsumerGovProposalResponse{}, nil
}
