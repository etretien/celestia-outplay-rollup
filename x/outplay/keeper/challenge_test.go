package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "outplay/testutil/keeper"
	"outplay/testutil/nullify"
	"outplay/x/outplay/keeper"
	"outplay/x/outplay/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNChallenge(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Challenge {
	items := make([]types.Challenge, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetChallenge(ctx, items[i])
	}
	return items
}

func TestChallengeGet(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	items := createNChallenge(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChallenge(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestChallengeRemove(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	items := createNChallenge(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChallenge(ctx,
			item.Index,
		)
		_, found := keeper.GetChallenge(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestChallengeGetAll(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	items := createNChallenge(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllChallenge(ctx)),
	)
}
