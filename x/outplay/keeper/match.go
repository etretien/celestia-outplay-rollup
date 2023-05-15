package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"outplay/x/outplay/types"
)

// SetMatch set a specific match in the store from its index
func (k Keeper) SetMatch(ctx sdk.Context, match types.Match) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MatchKeyPrefix))
	b := k.cdc.MustMarshal(&match)
	store.Set(types.MatchKey(
		match.Index,
	), b)
}

// GetMatch returns a match from its index
func (k Keeper) GetMatch(
	ctx sdk.Context,
	index string,

) (val types.Match, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MatchKeyPrefix))

	b := store.Get(types.MatchKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMatch removes a match from the store
func (k Keeper) RemoveMatch(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MatchKeyPrefix))
	store.Delete(types.MatchKey(
		index,
	))
}

// GetAllMatch returns all match
func (k Keeper) GetAllMatch(ctx sdk.Context) (list []types.Match) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MatchKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Match
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
