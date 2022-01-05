package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/batphonghan/blog/testutil/keeper"
	"github.com/batphonghan/blog/x/loan/keeper"
	"github.com/batphonghan/blog/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LoanKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
