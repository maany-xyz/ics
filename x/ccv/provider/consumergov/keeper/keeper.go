package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/log"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/maany-xyz/ics/v5/x/ccv/provider/consumergov/types"
)

type Keeper struct {
	cdc		  codec.BinaryCodec
    storeKey  storetypes.StoreKey
    govKeeper govkeeper.Keeper
    channelKeeper types.ChannelKeeper
}

func NewKeeper (
    cdc           codec.BinaryCodec,
    storeKey      storetypes.StoreKey,
    govKeeper 	  govkeeper.Keeper,
    channelKeeper types.ChannelKeeper,

) Keeper {
    return Keeper{
        cdc:           cdc,
        storeKey:      storeKey,
        govKeeper: 	   govKeeper,
        channelKeeper: channelKeeper,
    }
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.Logger().With("module", "x/"+types.ModuleName)
}


func (k Keeper) SendCustomIBCMessage(ctx sdk.Context, channelID string, data []byte) error {
    portID := "consumer" // Make sure this is the correct port

    channel, found := k.channelKeeper.GetChannel(ctx, portID, channelID)
    if !found {
        return fmt.Errorf("channel not found for port %s and channel %s", portID, channelID)
    }
    seq, found := k.channelKeeper.GetNextSequenceSend(ctx, portID, channelID)
    if !found {
                return fmt.Errorf("couldnt sequence channel ", portID, channelID)

    }

    // Define the packet data structure
    packetData := channeltypes.Packet{
        Sequence:           seq,
        SourcePort:         portID,
        SourceChannel:      channelID,
        DestinationPort:    channel.Counterparty.PortId,
        DestinationChannel: channel.Counterparty.ChannelId,
        Data:               data,  // Your encoded message
        TimeoutTimestamp:   uint64(ctx.BlockTime().UnixNano()) + 60000000000, // 60 sec timeout
    }

    k.Logger().Info("dkdkd", "kd", packetData)

    // Send the packet via IBC
    // err := k.channelKeeper.SendPacket(ctx, nil, packetData)
    // if err != nil {
    //     return fmt.Errorf("failed to send IBC packet: %v", err)
    // }

    return nil
} 