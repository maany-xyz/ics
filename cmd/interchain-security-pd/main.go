package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

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
}
