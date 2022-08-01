package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "lottery-chain-nel/testutil/keeper"
	"lottery-chain-nel/x/lottery/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LotteryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
