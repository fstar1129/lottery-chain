package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lottery module sentinel errors
var (
	ErrLotteryAlreadyOwned = sdkerrors.Register(ModuleName, 1000, "Lottery already has an owner")
)
