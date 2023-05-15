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

func createNMatch(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Match {
	items := make([]types.Match, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetMatch(ctx, items[i])
	}
	return items
}

func TestMatchGet(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	items := createNMatch(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMatch(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMatchRemove(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	items := createNMatch(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMatch(ctx,
			item.Index,
		)
		_, found := keeper.GetMatch(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestMatchGetAll(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	items := createNMatch(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMatch(ctx)),
	)
}
