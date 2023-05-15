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

func (k Keeper) MatchAll(goCtx context.Context, req *types.QueryAllMatchRequest) (*types.QueryAllMatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var matchs []types.Match
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	matchStore := prefix.NewStore(store, types.KeyPrefix(types.MatchKeyPrefix))

	pageRes, err := query.Paginate(matchStore, req.Pagination, func(key []byte, value []byte) error {
		var match types.Match
		if err := k.cdc.Unmarshal(value, &match); err != nil {
			return err
		}

		matchs = append(matchs, match)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMatchResponse{Match: matchs, Pagination: pageRes}, nil
}

func (k Keeper) Match(goCtx context.Context, req *types.QueryGetMatchRequest) (*types.QueryGetMatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetMatch(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMatchResponse{Match: val}, nil
}
