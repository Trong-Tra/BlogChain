package keeper

import (
	"context"
	"fmt"

	"blog/x/blog/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Post{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
	}
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

	k.SetPost(ctx, post)
	fmt.Printf("Post %v is successfully update", PostId)
	return &types.MsgUpdatePostResponse{}, nil
}
