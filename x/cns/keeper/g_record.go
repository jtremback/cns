package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jtremback/cns/x/cns/types"
)

// SetGRecord set a specific gRecord in the store from its index
func (k Keeper) SetGRecord(ctx sdk.Context, gRecord types.GRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GRecordKeyPrefix))
	b := k.cdc.MustMarshal(&gRecord)
	store.Set(types.GRecordKey(
		gRecord.Key,
	), b)
}

// GetGRecord returns a gRecord from its index
func (k Keeper) GetGRecord(
	ctx sdk.Context,
	key string,

) (val types.GRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GRecordKeyPrefix))

	b := store.Get(types.GRecordKey(
		key,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGRecord removes a gRecord from the store
func (k Keeper) RemoveGRecord(
	ctx sdk.Context,
	key string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GRecordKeyPrefix))
	store.Delete(types.GRecordKey(
		key,
	))
}

// GetAllGRecord returns all gRecord
func (k Keeper) GetAllGRecord(ctx sdk.Context) (list []types.GRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
