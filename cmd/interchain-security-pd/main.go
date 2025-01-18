package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	appparams "github.com/maany-xyz/ics/v5/app/params"
	app "github.com/maany-xyz/ics/v5/app/provider"
	"github.com/maany-xyz/ics/v5/cmd/interchain-security-pd/cmd"
)

func main() {
	appparams.SetAddressPrefixes("maany")
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}

	// Print Module account address:
	// moduleName := "blockrewards" // The name of your module account
    // moduleAddress := authtypes.NewModuleAddress(moduleName)
    // fmt.Printf("Address for module '%s': %s\n", moduleName, moduleAddress.String())
}
