package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jtremback/cns/x/cns/types"
)

func (k msgServer) CreateGRecord(goCtx context.Context, msg *types.MsgCreateGRecord) (*types.MsgCreateGRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetGRecord(
		ctx,
		msg.Key,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var gRecord = types.GRecord{
		Creator: msg.Creator,
		Key:     msg.Key,
		Value:   msg.Value,
	}

	k.SetGRecord(
		ctx,
		gRecord,
	)
	return &types.MsgCreateGRecordResponse{}, nil
}

func (k msgServer) UpdateGRecord(goCtx context.Context, msg *types.MsgUpdateGRecord) (*types.MsgUpdateGRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetGRecord(
		ctx,
		msg.Key,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var gRecord = types.GRecord{
		Creator: msg.Creator,
		Key:     msg.Key,
		Value:   msg.Value,
	}

	k.SetGRecord(ctx, gRecord)

	return &types.MsgUpdateGRecordResponse{}, nil
}

func (k msgServer) DeleteGRecord(goCtx context.Context, msg *types.MsgDeleteGRecord) (*types.MsgDeleteGRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetGRecord(
		ctx,
		msg.Key,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveGRecord(
		ctx,
		msg.Key,
	)

	return &types.MsgDeleteGRecordResponse{}, nil
}
