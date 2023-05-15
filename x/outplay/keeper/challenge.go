package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"outplay/x/outplay/types"
)

// SetChallenge set a specific challenge in the store from its index
func (k Keeper) SetChallenge(ctx sdk.Context, challenge types.Challenge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKeyPrefix))
	b := k.cdc.MustMarshal(&challenge)
	store.Set(types.ChallengeKey(
		challenge.Index,
	), b)
}

// GetChallenge returns a challenge from its index
func (k Keeper) GetChallenge(
	ctx sdk.Context,
	index string,

) (val types.Challenge, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKeyPrefix))

	b := store.Get(types.ChallengeKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChallenge removes a challenge from the store
func (k Keeper) RemoveChallenge(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKeyPrefix))
	store.Delete(types.ChallengeKey(
		index,
	))
}

// GetAllChallenge returns all challenge
func (k Keeper) GetAllChallenge(ctx sdk.Context) (list []types.Challenge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Challenge
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
