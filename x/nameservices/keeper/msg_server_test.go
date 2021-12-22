package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/batphonghan/blog/testutil/keeper"
	"github.com/batphonghan/blog/x/nameservices/keeper"
	"github.com/batphonghan/blog/x/nameservices/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NameservicesKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
