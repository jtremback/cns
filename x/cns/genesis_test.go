package cns_test

import (
	"testing"

	keepertest "github.com/jtremback/cns/testutil/keeper"
	"github.com/jtremback/cns/testutil/nullify"
	"github.com/jtremback/cns/x/cns"
	"github.com/jtremback/cns/x/cns/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		GRecordList: []types.GRecord{
			{
				Key: "0",
			},
			{
				Key: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CnsKeeper(t)
	cns.InitGenesis(ctx, *k, genesisState)
	got := cns.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.GRecordList, got.GRecordList)
	// this line is used by starport scaffolding # genesis/test/assert
}
