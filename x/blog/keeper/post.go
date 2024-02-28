package keeper

import (
	//"encoding/binary"

	"blog/x/blog/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)
	post.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.Postkey))
	appendedValue := k.cdc.MustMarshal(&post)
	store.Set(GetPostIDBytes(post.id), appendedValue)
	k.SetPostCount(ctx, count+1)
	return count
}
