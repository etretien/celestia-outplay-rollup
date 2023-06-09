package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"outplay/x/outplay/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated profile",
			genState: &types.GenesisState{
				ProfileList: []types.Profile{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated challenge",
			genState: &types.GenesisState{
				ChallengeList: []types.Challenge{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated match",
			genState: &types.GenesisState{
				MatchList: []types.Match{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
