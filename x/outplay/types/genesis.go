package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ProfileList:   []Profile{},
		ChallengeList: []Challenge{},
		MatchList:     []Match{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in profile
	profileIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProfileList {
		index := string(ProfileKey(elem.Index))
		if _, ok := profileIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for profile")
		}
		profileIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in challenge
	challengeIndexMap := make(map[string]struct{})

	for _, elem := range gs.ChallengeList {
		index := string(ChallengeKey(elem.Index))
		if _, ok := challengeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for challenge")
		}
		challengeIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in match
	matchIndexMap := make(map[string]struct{})

	for _, elem := range gs.MatchList {
		index := string(MatchKey(elem.Index))
		if _, ok := matchIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for match")
		}
		matchIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
