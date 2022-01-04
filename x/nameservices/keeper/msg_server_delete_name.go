package keeper

import (
	"context"

	"github.com/batphonghan/blog/x/nameservices/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteName(goCtx context.Context, msg *types.MsgDeleteName) (*types.MsgDeleteNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteNameResponse{}, nil
}
