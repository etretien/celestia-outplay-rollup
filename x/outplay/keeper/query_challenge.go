package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"outplay/x/outplay/types"
)

func (k Keeper) ChallengeAll(goCtx context.Context, req *types.QueryAllChallengeRequest) (*types.QueryAllChallengeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var challenges []types.Challenge
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	challengeStore := prefix.NewStore(store, types.KeyPrefix(types.ChallengeKeyPrefix))

	pageRes, err := query.Paginate(challengeStore, req.Pagination, func(key []byte, value []byte) error {
		var challenge types.Challenge
		if err := k.cdc.Unmarshal(value, &challenge); err != nil {
			return err
		}

		challenges = append(challenges, challenge)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChallengeResponse{Challenge: challenges, Pagination: pageRes}, nil
}

func (k Keeper) Challenge(goCtx context.Context, req *types.QueryGetChallengeRequest) (*types.QueryGetChallengeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetChallenge(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetChallengeResponse{Challenge: val}, nil
}
