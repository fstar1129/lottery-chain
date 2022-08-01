package keeper_test

import (
	keepertest "lottery-chain-nel/testutil/keeper"
	"lottery-chain-nel/x/lottery/keeper"
	"lottery-chain-nel/x/lottery/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestClaimOwner(t *testing.T) {
	k, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgServer := keeper.NewMsgServerImpl(*k)

	/* ClaimOwner when lottery is unowned */
	var message = types.MsgClaimOwner{
		Creator: "pass",
	}

	var expectedResponse = types.MsgClaimOwnerResponse{
		Success: true,
	}

	response, err := msgServer.ClaimOwner(wctx, &message)

	// check correct response, and no error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.NoError(t, err)

	// check that owner was found and stored
	owner, found := k.GetOwner(ctx)
	require.Equal(t, true, found)
	require.Equal(t, "pass", owner.Owner)

	///

	/* ClaimOwner when lottery is owned */
	expectedOwner := createTestOwner(k, ctx)

	message = types.MsgClaimOwner{
		Creator: "fail",
	}

	expectedResponse = types.MsgClaimOwnerResponse{
		Success: false,
	}

	response, err = msgServer.ClaimOwner(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrLotteryAlreadyOwned, err)

	// check that owner value is unchanged
	owner, _ = k.GetOwner(ctx)
	require.Equal(t, expectedOwner.Owner, owner.Owner)
}
