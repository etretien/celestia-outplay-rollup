package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"outplay/x/outplay/types"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) CreateChallenge(goCtx context.Context, msg *types.MsgCreateChallenge) (*types.MsgCreateChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert the stake to a number and check its range
	stake, _ := strconv.ParseFloat(msg.Stake, 64)
	if stake < 0 {
		stake = 0
	}

	// Use the current Unix timestamp to create a unique index for the challenge
	currentTime := time.Now().Unix()
	var currentTimeBytes = []byte(strconv.FormatInt(currentTime, 10))
	var currentTimeHash = sha256.Sum256(currentTimeBytes)
	var currentTimeHashString = hex.EncodeToString(currentTimeHash[:])

	// Create a new challenge
	var challenge = types.Challenge{
		Index:       currentTimeHashString,
		ChallengeId: currentTimeHashString,
		Stake:       strconv.FormatFloat(stake, 'f', 2, 64),
		Challenger:  msg.Creator,
		Challenged:  msg.Opponent,
		Status:      "open",
	}

	// Store the challenge in the keeper
	k.SetChallenge(ctx, challenge)

	// Lock the challenger's coins for escrow
	stakeCoins := sdk.Coins{sdk.NewInt64Coin("token", int64(stake))}
	challengerAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	err := k.bankKeeper.SendCoins(ctx, challengerAddress, moduleAcct, stakeCoins)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateChallengeResponse{}, nil
}
