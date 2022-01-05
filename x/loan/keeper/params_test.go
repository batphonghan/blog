package keeper_test

import (
	"testing"

	testkeeper "github.com/batphonghan/blog/testutil/keeper"
	"github.com/batphonghan/blog/x/loan/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LoanKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
