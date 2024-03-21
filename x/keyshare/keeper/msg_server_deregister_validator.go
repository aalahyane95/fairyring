package keeper

import (
	"context"
	"github.com/Fairblock/fairyring/x/keyshare/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DeRegisterValidator remove a validator in the validator set
func (k msgServer) DeRegisterValidator(goCtx context.Context, msg *types.MsgDeRegisterValidator) (*types.MsgDeRegisterValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetValidatorSet(ctx, msg.Creator)

	if !found {
		return nil, types.ErrValidatorNotRegistered.Wrap(msg.Creator)
	}

	if k.GetAuthorizedCount(ctx, msg.Creator) > 0 {
		return nil, types.ErrAuthorizedAnotherAddr
	}

	k.RemoveValidatorSet(ctx, msg.Creator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.DeRegisteredValidatorEventType,
			sdk.NewAttribute(types.DeRegisteredValidatorEventCreator, msg.Creator),
		),
	)

	return &types.MsgDeRegisterValidatorResponse{
		Creator: msg.Creator,
	}, nil
}
