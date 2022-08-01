package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "lottery-chain-nel/testutil/keeper"
	"lottery-chain-nel/testutil/nullify"
	"lottery-chain-nel/x/lottery/keeper"
	"lottery-chain-nel/x/lottery/types"
)

func createTestOwner(keeper *keeper.Keeper, ctx sdk.Context) types.Owner {
	item := types.Owner{}
	keeper.SetOwner(ctx, item)
	return item
}

func TestOwnerGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestOwner(keeper, ctx)
	rst, found := keeper.GetOwner(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestOwnerRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestOwner(keeper, ctx)
	keeper.RemoveOwner(ctx)
	_, found := keeper.GetOwner(ctx)
	require.False(t, found)
}
