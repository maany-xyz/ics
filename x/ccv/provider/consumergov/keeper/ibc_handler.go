package keeper

// import (
// 	"fmt"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/maany-xyz/ics/v5/x/ccv/provider/consumergov/types"
// )

// // Restrict IBC governance messages to the approved consumer chain
// func (k Keeper) SendConsumerGovIBCMessage(ctx sdk.Context, msg types.GovMsg) error {
//     // params := k.GetParams(ctx)

//     // if msg.ConsumerID != params.AllowedConsumerChainID {
//     //     return fmt.Errorf("IBC governance messages can only be sent to the approved consumer chain")
// 	// 	//sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "IBC governance messages can only be sent to the approved consumer chain")
//     // }

//     // return k.icsKeeper.SendInterchainMsg(ctx, msg)
// 	return nil
// }
