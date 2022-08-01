package keeper

import (
	"loan-with-cosmos/x/loan/types"
)

var _ types.QueryServer = Keeper{}
