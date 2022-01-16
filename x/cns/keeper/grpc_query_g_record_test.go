package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/jtremback/cns/testutil/keeper"
	"github.com/jtremback/cns/testutil/nullify"
	"github.com/jtremback/cns/x/cns/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGRecordQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CnsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGRecord(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGRecordRequest
		response *types.QueryGetGRecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGRecordRequest{
				Key: msgs[0].Key,
			},
			response: &types.QueryGetGRecordResponse{GRecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGRecordRequest{
				Key: msgs[1].Key,
			},
			response: &types.QueryGetGRecordResponse{GRecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGRecordRequest{
				Key: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GRecord(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestGRecordQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CnsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGRecord(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGRecordRequest {
		return &types.QueryAllGRecordRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GRecordAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GRecord),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GRecordAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GRecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GRecordAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.GRecord),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GRecordAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
