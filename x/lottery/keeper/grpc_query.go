package keeper

import (
	"lottery-chain-nel/x/lottery/types"
)

var _ types.QueryServer = Keeper{}
