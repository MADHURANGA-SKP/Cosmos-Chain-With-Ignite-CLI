package keeper

import (
	"context"
	"strconv"

	"pasans/x/pasans/types"

	rules "pasans/x/pasans/rules"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)
	newGame := rules.New()
	storedGame := types.StoredGame{
		Index: newIndex,
		Board: newGame.String(),
		Turn:  rules.PieceStrings[newGame.Turn],
		Black: msg.Black,
		Red:   msg.Red,
		Winner:    rules.PieceStrings[rules.NO_PLAYER],
	}

	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredGame(ctx, storedGame)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.GameCreatedEventType,
			sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameCreatedEventGameIndex, newIndex),
			sdk.NewAttribute(types.GameCreatedEventBlack, msg.Black),
			sdk.NewAttribute(types.GameCreatedEventRed, msg.Red),
		),
	)	

	return &types.MsgCreateGameResponse{
		GameIndex: newIndex,
	}, nil
}
