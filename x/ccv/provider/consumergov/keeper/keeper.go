package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"

	"cosmossdk.io/errors"
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
    timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + 10000000000      // 60 seconds timeout
    name := host.ChannelCapabilityPath(portID, channelID)
    ctx.Logger().Info("In here with timestamp port and channelId ", "portId", portID, "channekId", channelID, "tmstp", timeoutTimestamp, "capName", name )

    channelCap, found := k.scopedKeeper.GetCapability(ctx, name)
    ctx.Logger().Info("The channel cap is ", "chanelCap", channelCap, "found", found )

    if !found {
        ctx.Logger().Error("The channelCap not found ")
        return errors.New("Error:", 103, "cant find channel capabilites")
    }
    timeoutHeight := clienttypes.NewHeight(0, uint64(ctx.BlockHeight()+10)) // 100 blocks timeout

    ctx.Logger().Info("in here with ", "portID", portID, "channelId", channelID, "timeoutHeight", timeoutHeight, "channelCap", channelCap)

    res, err := k.channelKeeper.SendPacket(ctx, channelCap, portID, channelID, timeoutHeight, timeoutTimestamp, data)
    if err != nil {
        ctx.Logger().Error("The err is ", "err", err)
        return err
    }

    ctx.Logger().Info("The res is ", "res", res)


    return nil
} 