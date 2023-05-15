package keeper

import (
	"context"
	"outplay/x/outplay/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) AcceptChallenge(goCtx context.Context, msg *types.MsgAcceptChallenge) (*types.MsgAcceptChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the challenge from the keeper
	challenge, found := k.GetChallenge(ctx, msg.ChallengeId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Challenge not found")
	}

	// Check if the challenge has already been accepted
	if challenge.Status != "open" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Challenge has already been accepted or declined")
	}

	// Lock the challenged user's coins for escrow
	stake, _ := strconv.ParseFloat(challenge.Stake, 64)
	stakeCoins := sdk.Coins{sdk.NewInt64Coin("token", int64(stake))}
	challengedAddress, _ := sdk.AccAddressFromBech32(challenge.Challenged)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	err := k.bankKeeper.SendCoins(ctx, challengedAddress, moduleAcct, stakeCoins)
	if err != nil {
		return nil, err
	}

	// Update the challenge status to "accepted"
	challenge.Status = "accepted"
	k.SetChallenge(ctx, challenge)

	return &types.MsgAcceptChallengeResponse{}, nil
}
