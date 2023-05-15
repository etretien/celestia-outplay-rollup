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

func createNProfile(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Profile {
	items := make([]types.Profile, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetProfile(ctx, items[i])
	}
	return items
}

func TestProfileGet(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	items := createNProfile(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProfile(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProfileRemove(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	items := createNProfile(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProfile(ctx,
			item.Index,
		)
		_, found := keeper.GetProfile(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestProfileGetAll(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	items := createNProfile(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProfile(ctx)),
	)
}
