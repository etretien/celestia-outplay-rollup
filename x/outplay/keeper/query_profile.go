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

func (k Keeper) ProfileAll(goCtx context.Context, req *types.QueryAllProfileRequest) (*types.QueryAllProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var profiles []types.Profile
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	profileStore := prefix.NewStore(store, types.KeyPrefix(types.ProfileKeyPrefix))

	pageRes, err := query.Paginate(profileStore, req.Pagination, func(key []byte, value []byte) error {
		var profile types.Profile
		if err := k.cdc.Unmarshal(value, &profile); err != nil {
			return err
		}

		profiles = append(profiles, profile)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProfileResponse{Profile: profiles, Pagination: pageRes}, nil
}

func (k Keeper) Profile(goCtx context.Context, req *types.QueryGetProfileRequest) (*types.QueryGetProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetProfile(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProfileResponse{Profile: val}, nil
}
