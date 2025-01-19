package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	app "github.com/maany-xyz/ics/v5/app/dex_consumer"
	appparams "github.com/maany-xyz/ics/v5/app/params"
	"github.com/maany-xyz/ics/v5/cmd/dex-consumer-d/cmd"
)

func main() {
	appparams.SetAddressPrefixes(app.Bech32MainPrefix)

	rootCmd := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}

	// Print Module account address:
	// moduleName := "mintburn" // The name of your module account
    // moduleAddress := authtypes.NewModuleAddress(moduleName)
    // fmt.Printf("Address for module '%s': %s\n", moduleName, moduleAddress.String())
}
