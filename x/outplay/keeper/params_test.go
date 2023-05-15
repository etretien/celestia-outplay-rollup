package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "outplay/testutil/keeper"
	"outplay/x/outplay/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OutplayKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
