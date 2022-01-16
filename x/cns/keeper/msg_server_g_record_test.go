package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/jtremback/cns/testutil/keeper"
	"github.com/jtremback/cns/x/cns/keeper"
	"github.com/jtremback/cns/x/cns/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGRecordMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CnsKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateGRecord{Creator: creator,
			Key: strconv.Itoa(i),
		}
		_, err := srv.CreateGRecord(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetGRecord(ctx,
			expected.Key,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestGRecordMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateGRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateGRecord{Creator: creator,
				Key: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateGRecord{Creator: "B",
				Key: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateGRecord{Creator: creator,
				Key: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CnsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateGRecord{Creator: creator,
				Key: strconv.Itoa(0),
			}
			_, err := srv.CreateGRecord(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateGRecord(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetGRecord(ctx,
					expected.Key,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestGRecordMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteGRecord
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteGRecord{Creator: creator,
				Key: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteGRecord{Creator: "B",
				Key: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteGRecord{Creator: creator,
				Key: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CnsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateGRecord(wctx, &types.MsgCreateGRecord{Creator: creator,
				Key: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteGRecord(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetGRecord(ctx,
					tc.request.Key,
				)
				require.False(t, found)
			}
		})
	}
}
