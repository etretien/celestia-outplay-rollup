package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAcceptChallenge = "accept_challenge"

var _ sdk.Msg = &MsgAcceptChallenge{}

func NewMsgAcceptChallenge(creator string, challengeId string) *MsgAcceptChallenge {
	return &MsgAcceptChallenge{
		Creator:     creator,
		ChallengeId: challengeId,
	}
}

func (msg *MsgAcceptChallenge) Route() string {
	return RouterKey
}

func (msg *MsgAcceptChallenge) Type() string {
	return TypeMsgAcceptChallenge
}

func (msg *MsgAcceptChallenge) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAcceptChallenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptChallenge) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
