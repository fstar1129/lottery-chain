package keeper

import (
	"context"

	"lottery-chain-nel/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ClaimOwner(
	goCtx context.Context,
	msg *types.MsgClaimOwner,
) (*types.MsgClaimOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if an owner already exists
	owner, alreadyOwned := k.GetOwner(ctx)
	if alreadyOwned {
		return &types.MsgClaimOwnerResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryAlreadyOwned,
				"Lottery already owned by %v",
				owner.Owner,
			)
	}

	// set owner as message creator
	owner = types.Owner{
		Owner: msg.GetCreator(),
	}

	k.SetOwner(ctx, owner)

	return &types.MsgClaimOwnerResponse{Success: true}, nil
}
