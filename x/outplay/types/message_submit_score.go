package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitScore = "submit_score"

var _ sdk.Msg = &MsgSubmitScore{}

func NewMsgSubmitScore(creator string, challengeId string, score string) *MsgSubmitScore {
	return &MsgSubmitScore{
		Creator:     creator,
		ChallengeId: challengeId,
		Score:       score,
	}
}

func (msg *MsgSubmitScore) Route() string {
	return RouterKey
}

func (msg *MsgSubmitScore) Type() string {
	return TypeMsgSubmitScore
}

func (msg *MsgSubmitScore) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitScore) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitScore) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
