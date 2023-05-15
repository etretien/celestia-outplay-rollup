package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "outplay/testutil/keeper"
	"outplay/testutil/nullify"
	"outplay/x/outplay/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMatchQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMatch(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMatchRequest
		response *types.QueryGetMatchResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMatchRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetMatchResponse{Match: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMatchRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetMatchResponse{Match: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMatchRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Match(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestMatchQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.OutplayKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMatch(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMatchRequest {
		return &types.QueryAllMatchRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MatchAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Match), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Match),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MatchAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Match), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Match),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MatchAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Match),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MatchAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
