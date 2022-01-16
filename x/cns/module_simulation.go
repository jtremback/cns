package cns

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jtremback/cns/testutil/sample"
	cnssimulation "github.com/jtremback/cns/x/cns/simulation"
	"github.com/jtremback/cns/x/cns/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cnssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateGRecord = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGRecord int = 100

	opWeightMsgUpdateGRecord = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateGRecord int = 100

	opWeightMsgDeleteGRecord = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteGRecord int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cnsGenesis := types.GenesisState{
		GRecordList: []types.GRecord{
			{
				Creator: sample.AccAddress(),
				Key:     "0",
			},
			{
				Creator: sample.AccAddress(),
				Key:     "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cnsGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	cnsParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyWriter), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(cnsParams.Writer))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateGRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateGRecord, &weightMsgCreateGRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGRecord = defaultWeightMsgCreateGRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGRecord,
		cnssimulation.SimulateMsgCreateGRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateGRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateGRecord, &weightMsgUpdateGRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGRecord = defaultWeightMsgUpdateGRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateGRecord,
		cnssimulation.SimulateMsgUpdateGRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteGRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteGRecord, &weightMsgDeleteGRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteGRecord = defaultWeightMsgDeleteGRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteGRecord,
		cnssimulation.SimulateMsgDeleteGRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
