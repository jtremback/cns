package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jtremback/cns/testutil/keeper"
	"github.com/jtremback/cns/testutil/nullify"
	"github.com/jtremback/cns/x/cns/keeper"
	"github.com/jtremback/cns/x/cns/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GRecord {
	items := make([]types.GRecord, n)
	for i := range items {
		items[i].Key = strconv.Itoa(i)

		keeper.SetGRecord(ctx, items[i])
	}
	return items
}

func TestGRecordGet(t *testing.T) {
	keeper, ctx := keepertest.CnsKeeper(t)
	items := createNGRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGRecord(ctx,
			item.Key,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.CnsKeeper(t)
	items := createNGRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGRecord(ctx,
			item.Key,
		)
		_, found := keeper.GetGRecord(ctx,
			item.Key,
		)
		require.False(t, found)
	}
}

func TestGRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.CnsKeeper(t)
	items := createNGRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGRecord(ctx)),
	)
}
