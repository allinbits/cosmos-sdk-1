package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/v42/server"
	svrcmd "github.com/cosmos/cosmos-sdk/v42/server/cmd"
	"github.com/cosmos/cosmos-sdk/v42/simapp"
	"github.com/cosmos/cosmos-sdk/v42/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
