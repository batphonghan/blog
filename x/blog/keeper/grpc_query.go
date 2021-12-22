package keeper

import (
	"github.com/batphonghan/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
