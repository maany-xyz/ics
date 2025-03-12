package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"

	"cosmossdk.io/log"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	"github.com/maany-xyz/ics/v5/x/ccv/provider/consumergov/types"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

type Keeper struct {
	cdc		  codec.BinaryCodec
    storeKey  storetypes.StoreKey
    govKeeper govkeeper.Keeper
    channelKeeper types.ChannelKeeper
    scopedKeeper capabilitykeeper.ScopedKeeper
}

func NewKeeper (
    cdc           codec.BinaryCodec,
    storeKey      storetypes.StoreKey,
    govKeeper 	  govkeeper.Keeper,
    channelKeeper types.ChannelKeeper,
        scopedKeeper capabilitykeeper.ScopedKeeper,

) Keeper {
    return Keeper{
        cdc:           cdc,
        storeKey:      storeKey,
        govKeeper: 	   govKeeper,
        channelKeeper: channelKeeper,
        scopedKeeper: scopedKeeper,
    }
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.Logger().With("module", "x/"+types.ModuleName)
}


func (k Keeper) SendCustomIBCMessage(ctx sdk.Context, channelID string, data []byte) error {
    ctx.Logger().Info("inside SendCustomIBCMessage")

    portID := "provider" 

    timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + 60000000000      // 60 seconds timeout
    channelCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
    if !found {
        ctx.Logger().Error("The channelCap not found ")
        return fmt.Errorf("error: channel not found")
    }

    timeoutHeight := clienttypes.NewHeight(0, uint64(ctx.BlockHeight()+100)) // 100 blocks timeout

    res, err := k.channelKeeper.SendPacket(ctx, channelCap, portID, channelID, timeoutHeight, timeoutTimestamp, data)
    if err != nil {
        ctx.Logger().Error("The err is ", "err", err)
        return err
    }

    ctx.Logger().Info("The res is ", "res", res)


    return nil
} 