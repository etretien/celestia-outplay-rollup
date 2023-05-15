package outplay

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"outplay/testutil/sample"
	outplaysimulation "outplay/x/outplay/simulation"
	"outplay/x/outplay/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = outplaysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateProfile = "op_weight_msg_create_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProfile int = 100

	opWeightMsgCreateChallenge = "op_weight_msg_create_challenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateChallenge int = 100

	opWeightMsgAcceptChallenge = "op_weight_msg_accept_challenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAcceptChallenge int = 100

	opWeightMsgDeclineChallenge = "op_weight_msg_decline_challenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeclineChallenge int = 100

	opWeightMsgSubmitScore = "op_weight_msg_submit_score"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitScore int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	outplayGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&outplayGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateProfile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProfile, &weightMsgCreateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProfile = defaultWeightMsgCreateProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProfile,
		outplaysimulation.SimulateMsgCreateProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateChallenge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateChallenge, &weightMsgCreateChallenge, nil,
		func(_ *rand.Rand) {
			weightMsgCreateChallenge = defaultWeightMsgCreateChallenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateChallenge,
		outplaysimulation.SimulateMsgCreateChallenge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAcceptChallenge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAcceptChallenge, &weightMsgAcceptChallenge, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptChallenge = defaultWeightMsgAcceptChallenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAcceptChallenge,
		outplaysimulation.SimulateMsgAcceptChallenge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeclineChallenge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeclineChallenge, &weightMsgDeclineChallenge, nil,
		func(_ *rand.Rand) {
			weightMsgDeclineChallenge = defaultWeightMsgDeclineChallenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeclineChallenge,
		outplaysimulation.SimulateMsgDeclineChallenge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitScore int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitScore, &weightMsgSubmitScore, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitScore = defaultWeightMsgSubmitScore
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitScore,
		outplaysimulation.SimulateMsgSubmitScore(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
