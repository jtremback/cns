package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateGRecord = "create_g_record"
	TypeMsgUpdateGRecord = "update_g_record"
	TypeMsgDeleteGRecord = "delete_g_record"
)

var _ sdk.Msg = &MsgCreateGRecord{}

func NewMsgCreateGRecord(
	creator string,
	key string,
	value string,

) *MsgCreateGRecord {
	return &MsgCreateGRecord{
		Creator: creator,
		Key:     key,
		Value:   value,
	}
}

func (msg *MsgCreateGRecord) Route() string {
	return RouterKey
}

func (msg *MsgCreateGRecord) Type() string {
	return TypeMsgCreateGRecord
}

func (msg *MsgCreateGRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateGRecord{}

func NewMsgUpdateGRecord(
	creator string,
	key string,
	value string,

) *MsgUpdateGRecord {
	return &MsgUpdateGRecord{
		Creator: creator,
		Key:     key,
		Value:   value,
	}
}

func (msg *MsgUpdateGRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateGRecord) Type() string {
	return TypeMsgUpdateGRecord
}

func (msg *MsgUpdateGRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateGRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateGRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteGRecord{}

func NewMsgDeleteGRecord(
	creator string,
	key string,

) *MsgDeleteGRecord {
	return &MsgDeleteGRecord{
		Creator: creator,
		Key:     key,
	}
}
func (msg *MsgDeleteGRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteGRecord) Type() string {
	return TypeMsgDeleteGRecord
}

func (msg *MsgDeleteGRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteGRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteGRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
