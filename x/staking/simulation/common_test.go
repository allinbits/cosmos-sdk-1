package simulation_test

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/v42/types"
)

func init() {
	sdk.DefaultPowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}
