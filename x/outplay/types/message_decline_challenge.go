package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeclineChallenge = "decline_challenge"

var _ sdk.Msg = &MsgDeclineChallenge{}

func NewMsgDeclineChallenge(creator string, challengeId string) *MsgDeclineChallenge {
	return &MsgDeclineChallenge{
		Creator:     creator,
		ChallengeId: challengeId,
	}
}

func (msg *MsgDeclineChallenge) Route() string {
	return RouterKey
}

func (msg *MsgDeclineChallenge) Type() string {
	return TypeMsgDeclineChallenge
}

func (msg *MsgDeclineChallenge) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeclineChallenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeclineChallenge) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
