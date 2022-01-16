package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jtremback/cns/x/cns/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GRecordAll(c context.Context, req *types.QueryAllGRecordRequest) (*types.QueryAllGRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var gRecords []types.GRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	gRecordStore := prefix.NewStore(store, types.KeyPrefix(types.GRecordKeyPrefix))

	pageRes, err := query.Paginate(gRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var gRecord types.GRecord
		if err := k.cdc.Unmarshal(value, &gRecord); err != nil {
			return err
		}

		gRecords = append(gRecords, gRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGRecordResponse{GRecord: gRecords, Pagination: pageRes}, nil
}

func (k Keeper) GRecord(c context.Context, req *types.QueryGetGRecordRequest) (*types.QueryGetGRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGRecord(
		ctx,
		req.Key,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetGRecordResponse{GRecord: val}, nil
}
