package keeper

import (
	"context"
	"fmt"

	"blog/x/blog/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeletePost(goCtx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPost(ctx, msg.Id)
	PostId := msg.Id
	if !found {
		return nil,
			errorsmod.Wrap(sdkerrors.ErrKeyNotFound,
				fmt.Sprintf("Key %d doesn't exist", PostId))
	}
	if msg.Creator != val.Creator {
		return nil,
			errorsmod.Wrap(sdkerrors.ErrUnauthorized,
				"The owner information is incorrect. Please ask for authorization.")
	}

	k.RemovePost(ctx, msg.Id)
	fmt.Printf("Post %v is successfully delete", PostId)
	return &types.MsgDeletePostResponse{}, nil
}
