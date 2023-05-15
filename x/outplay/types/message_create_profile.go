package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateProfile = "create_profile"

var _ sdk.Msg = &MsgCreateProfile{}

func NewMsgCreateProfile(creator string, name string, dateOfBirth string, playingHand string, ntrpRating string) *MsgCreateProfile {
	return &MsgCreateProfile{
		Creator:     creator,
		Name:        name,
		DateOfBirth: dateOfBirth,
		PlayingHand: playingHand,
		NtrpRating:  ntrpRating,
	}
}

func (msg *MsgCreateProfile) Route() string {
	return RouterKey
}

func (msg *MsgCreateProfile) Type() string {
	return TypeMsgCreateProfile
}

func (msg *MsgCreateProfile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProfile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
