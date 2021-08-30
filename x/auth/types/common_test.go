package types_test

import (
	"github.com/cosmos/cosmos-sdk/v42/simapp"
)

var (
	ecdc                  = simapp.MakeTestEncodingConfig()
	appCodec, legacyAmino = ecdc.Codec, ecdc.Amino
)
