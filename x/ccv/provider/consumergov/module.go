package consumergov

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keeper "github.com/maany-xyz/ics/v5/x/ccv/provider/consumergov/keeper"
	"github.com/maany-xyz/ics/v5/x/ccv/provider/consumergov/types"
)

var (
	 _ appmodule.HasEndBlocker = AppModule{}
)

// AppModuleBasic defines the basic application module used by the blockrewards module.
type AppModuleBasic struct{}

// Name returns the blockrewards module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// RegisterLegacyAminoCodec registers the blockrewards module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterInterfaces registers the module's protobuf interfaces.
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&types.MsgConsumerGovProposal{},
	)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the blockrewards module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// GetTxCmd returns the root tx command for the blockrewards module.
func (AppModuleBasic) GetTxCmd() *cobra.Command { return nil }

// GetQueryCmd returns the root query command for the blockrewards module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }

// AppModule implements the AppModule interface for the blockrewards module.
type AppModule struct {
	cdc codec.Codec
	AppModuleBasic
	keeper *keeper.Keeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(cdc codec.Codec, k *keeper.Keeper) AppModule {
	return AppModule{
		cdc: cdc,
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}
// RegisterInvariants registers the invariants for the blockrewards module.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// RegisterServices registers the module's services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
		types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))

}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }


// BeginBlock executes all logic for the blockrewards module at the beginning of a block.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestFinalizeBlock) {}

// EndBlock executes all logic for the blockrewards module at the end of a block.
func (am AppModule) EndBlock(ctx context.Context) error {
	return nil
}


// func (am AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
//     // Return the default genesis state marshaled as JSON
//     gen := types.DefaultGenesisState()
//     return cdc.MustMarshalJSON(&gen)
// }

// func (am AppModule) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
//     // Unmarshal the genesis state from the JSON
//     // var genState types.GenesisState
//     // if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
//     //     return fmt.Errorf("failed to unmarshal genesis state: %w", err)
//     // }
//     // Validate the genesis state
//     return nil
// }

// func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
    
//     return nil
// }

// func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
//      gs := am.keeper.ExportGenesis(ctx)
//     return cdc.MustMarshalJSON(gs)
// }