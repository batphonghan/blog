package keeper

import (
	"github.com/batphonghan/blog/x/nameservices/types"
)

var _ types.QueryServer = Keeper{}
