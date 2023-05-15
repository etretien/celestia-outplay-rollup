package outplay_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "outplay/testutil/keeper"
	"outplay/testutil/nullify"
	"outplay/x/outplay"
	"outplay/x/outplay/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ProfileList: []types.Profile{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ChallengeList: []types.Challenge{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		MatchList: []types.Match{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OutplayKeeper(t)
	outplay.InitGenesis(ctx, *k, genesisState)
	got := outplay.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProfileList, got.ProfileList)
	require.ElementsMatch(t, genesisState.ChallengeList, got.ChallengeList)
	require.ElementsMatch(t, genesisState.MatchList, got.MatchList)
	// this line is used by starport scaffolding # genesis/test/assert
}
