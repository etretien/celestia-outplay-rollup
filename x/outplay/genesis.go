package outplay

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"outplay/x/outplay/keeper"
	"outplay/x/outplay/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the profile
	for _, elem := range genState.ProfileList {
		k.SetProfile(ctx, elem)
	}
	// Set all the challenge
	for _, elem := range genState.ChallengeList {
		k.SetChallenge(ctx, elem)
	}
	// Set all the match
	for _, elem := range genState.MatchList {
		k.SetMatch(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ProfileList = k.GetAllProfile(ctx)
	genesis.ChallengeList = k.GetAllChallenge(ctx)
	genesis.MatchList = k.GetAllMatch(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
