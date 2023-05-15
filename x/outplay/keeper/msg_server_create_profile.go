package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"outplay/x/outplay/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateProfile(goCtx context.Context, msg *types.MsgCreateProfile) (*types.MsgCreateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ntrpRating, _ := strconv.ParseFloat(msg.NtrpRating, 64)
	if ntrpRating < 1.0 {
		ntrpRating = 1.0
	} else if ntrpRating > 6.0 {
		ntrpRating = 6.0
	}

	ownerHash := sha256.Sum256([]byte(msg.Creator))

	var profile = types.Profile{
		Index:       hex.EncodeToString(ownerHash[:]),
		Owner:       msg.Creator,
		Name:        msg.Name,
		DateOfBirth: msg.DateOfBirth,
		PlayingHand: msg.PlayingHand,
		NtrpRating:  strconv.FormatFloat(ntrpRating, 'f', 1, 64),
		Elo:         strconv.Itoa(int(400*ntrpRating + 600)),
	}

	k.SetProfile(ctx, profile)

	return &types.MsgCreateProfileResponse{}, nil
}
