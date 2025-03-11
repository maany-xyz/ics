package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/log"
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