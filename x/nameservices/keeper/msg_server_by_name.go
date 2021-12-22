package keeper

import (
	"context"

	"github.com/batphonghan/blog/x/nameservices/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ByName(goCtx context.Context, msg *types.MsgByName) (*types.MsgByNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgByNameResponse{}, nil
}
