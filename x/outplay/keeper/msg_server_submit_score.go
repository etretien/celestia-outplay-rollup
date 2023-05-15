package keeper

import (
	"context"
	"fmt"
	"math"
	"outplay/x/outplay/types"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) SubmitScore(goCtx context.Context, msg *types.MsgSubmitScore) (*types.MsgSubmitScoreResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the challenge from the keeper
	challenge, found := k.GetChallenge(ctx, msg.ChallengeId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Challenge not found")
	}

	// Check if the challenge has been accepted
	if challenge.Status != "accepted" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Challenge has not been accepted")
	}

	// Parse the score
	parsedScores, err := parseScore(msg.Score)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid score format")
	}

	// Determine the winner and loser
	challengerWins, challengedWins := 0, 0
	challengerGames, challengedGames := 0, 0
	for _, set := range parsedScores {
		challengerGames += set[0]
		challengedGames += set[1]
		if set[0] > set[1] {
			challengerWins++
		} else {
			challengedWins++
		}
	}
	var winner, loser string
	var winnerScore, loserScore int
	if challengerWins > challengedWins {
		winner = challenge.Challenger
		loser = challenge.Challenged
		winnerScore = challengerGames
		loserScore = challengedGames
	} else {
		winner = challenge.Challenged
		loser = challenge.Challenger
		winnerScore = challengedGames
		loserScore = challengerGames
	}

	// Get the profiles of the winner and the loser
	winnerProfile, _ := k.GetProfile(ctx, winner)
	loserProfile, _ := k.GetProfile(ctx, loser)

	// Update the Elo ratings
	winnerOldRating, _ := strconv.Atoi(winnerProfile.Elo)
	loserOldRating, _ := strconv.Atoi(loserProfile.Elo)
	winnerNewRating, loserNewRating := eloRatingUpdate(winnerOldRating, loserOldRating, float64(loserScore)/float64(winnerScore), 32)

	// Update the profiles with the new Elo ratings
	winnerProfile.Elo = strconv.Itoa(winnerNewRating)
	loserProfile.Elo = strconv.Itoa(loserNewRating)
	k.SetProfile(ctx, winnerProfile)
	k.SetProfile(ctx, loserProfile)

	// Create a new match
	match := types.Match{
		Index:  challenge.ChallengeId,
		Score:  msg.Score,
		Winner: winner,
		Loser:  loser,
	}

	// Save the match to the keeper
	k.SetMatch(ctx, match)

	// Convert the stake to an integer
	stake, err := strconv.Atoi(challenge.Stake)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid stake format")
	}

	// Unlock the reward
	reward := sdk.Coins{sdk.NewInt64Coin("stake", int64(2*stake))}
	winnerAddress, _ := sdk.AccAddressFromBech32(winner)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	err = k.bankKeeper.SendCoins(ctx, moduleAcct, winnerAddress, reward)
	if err != nil {
		return nil, err
	}

	return &types.MsgSubmitScoreResponse{}, nil
}

func parseScore(score string) ([][2]int, error) {
	// Remove spaces
	score = strings.ReplaceAll(score, " ", "")

	// Split the score into sets
	sets := strings.Split(score, ",")

	// Check the number of sets
	if len(sets) < 2 || len(sets) > 3 {
		return nil, fmt.Errorf("Invalid number of sets")
	}

	// Prepare a slice to hold the parsed scores
	parsedScores := make([][2]int, len(sets))

	// Parse each set
	for i, set := range sets {
		// Split the set into games
		games := strings.Split(set, ":")

		// Check the number of games
		if len(games) != 2 {
			return nil, fmt.Errorf("Invalid number of games in set")
		}

		// Parse the number of games won by each player
		challengerGames, err := strconv.Atoi(games[0])
		if err != nil {
			return nil, fmt.Errorf("Invalid number of games won by challenger")
		}
		challengedGames, err := strconv.Atoi(games[1])
		if err != nil {
			return nil, fmt.Errorf("Invalid number of games won by challenged")
		}

		// Store the parsed scores
		parsedScores[i] = [2]int{challengerGames, challengedGames}
	}

	return parsedScores, nil
}

func eloRatingUpdate(winnerOldRating int, loserOldRating int, score float64, kFactor int) (int, int) {
	expectedScoreWinner := 1 / (1 + math.Pow(10, float64(loserOldRating-winnerOldRating)/400))
	expectedScoreLoser := 1 - expectedScoreWinner
	winnerNewRating := winnerOldRating + int(float64(kFactor)*(1-expectedScoreWinner))
	loserNewRating := loserOldRating + int(float64(kFactor)*(score-expectedScoreLoser))
	return winnerNewRating, loserNewRating
}
