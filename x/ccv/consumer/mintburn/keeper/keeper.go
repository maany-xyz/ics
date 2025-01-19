package mintburn

import (
	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
)

type Keeper struct {
    bankKeeper bankKeeper.Keeper
    moduleName string
}

func NewKeeper(bankKeeper bankKeeper.Keeper, moduleName string) Keeper {
    return Keeper{
        bankKeeper: bankKeeper,
        moduleName: moduleName,
    }
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/-"+"mintburn")
}

func (k Keeper) MintTokens(ctx sdk.Context, recipient sdk.AccAddress, amount sdk.Coin) error {
    // Mint tokens to the module account
    ctx.Logger().Info("Now in MintTokens in Keeper")
    err := k.bankKeeper.MintCoins(ctx, k.moduleName, sdk.NewCoins(amount))
    if err != nil {
        return err
    }

    // Send minted tokens to the recipient
    return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, k.moduleName, recipient, sdk.NewCoins(amount))
}

func (k Keeper) GetBalances(ctx sdk.Context, address sdk.AccAddress) error {
    ctx.Logger().Info("Now in Get Balance in Keeper")
    balances := k.bankKeeper.GetAllBalances(ctx, address)
    ctx.Logger().Info("The transfer module account balances are: ", "balamces", balances.String())
    return nil
}

func (k Keeper) BurnTokens(ctx sdk.Context, recipient sdk.AccAddress, amount sdk.Coin) error {

    if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, ibctransfertypes.ModuleName, k.moduleName, sdk.NewCoins(amount)); err != nil {
		ctx.Logger().Error("Failed to redirect tokens to module account", "error", err)
	    return err
    }

    return nil
    // Transfer tokens to module account and burn them
    // ctx.Logger().Info("Now in BurnTokens in Keeper")
    // err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, recipient, k.moduleName, sdk.NewCoins(amount))
    // if err != nil {
    //      ctx.Logger().Error("Error sending bridged tokens ", "msg", err)
    //     return err
    // }
    // ctx.Logger().Info("Sent bridged tokens succesfully to module account.")
    // return k.bankKeeper.BurnCoins(ctx, k.moduleName, sdk.NewCoins(amount))
}
