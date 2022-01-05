package keeper

import (
	"context"

	"github.com/batphonghan/blog/x/scavenge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) SubmitScavenge(goCtx context.Context, msg *types.MsgSubmitScavenge) (*types.MsgSubmitScavengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	scavenge := types.Scavenge{
		Index:        msg.SolutionHash,
		SolutionHash: msg.SolutionHash,
		Description:  msg.Description,
		Reward:       msg.Reward,
		Scavenger:    msg.Creator,
	}

	_, isFound := k.GetScavenge(ctx, scavenge.SolutionHash)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Scavenge already commited")
	}

	// get address of the Scavenge module account
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	// Convert msg creator address from a string into sdk.AccAddress
	scavenger, err := sdk.AccAddressFromBech32(scavenge.Scavenger)
	if err != nil {
		panic(err)
	}

	// Convert token from the scavenge creator to the module account
	reward, err := sdk.ParseCoinsNormalized(scavenge.Reward)
	if err != nil {
		panic(err)
	}

	// Send tokens from the scavenge creator to the module account
	err = k.bankKeeper.SendCoins(ctx, scavenger, moduleAcct, reward)
	if err != nil {
		return nil, err
	}

	// write the scavenge to the store
	k.SetScavenge(ctx, scavenge)

	return &types.MsgSubmitScavengeResponse{}, nil
}
