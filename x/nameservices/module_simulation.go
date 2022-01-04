package nameservices

import (
	"math/rand"

	"github.com/batphonghan/blog/testutil/sample"
	nameservicessimulation "github.com/batphonghan/blog/x/nameservices/simulation"
	"github.com/batphonghan/blog/x/nameservices/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = nameservicessimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgByName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgByName int = 100

	opWeightMsgSetName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetName int = 100

	opWeightMsgDeleteName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteName int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nameservicesGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nameservicesGenesis)
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

	var weightMsgByName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgByName, &weightMsgByName, nil,
		func(_ *rand.Rand) {
			weightMsgByName = defaultWeightMsgByName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgByName,
		nameservicessimulation.SimulateMsgByName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetName, &weightMsgSetName, nil,
		func(_ *rand.Rand) {
			weightMsgSetName = defaultWeightMsgSetName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetName,
		nameservicessimulation.SimulateMsgSetName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteName, &weightMsgDeleteName, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteName = defaultWeightMsgDeleteName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteName,
		nameservicessimulation.SimulateMsgDeleteName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
