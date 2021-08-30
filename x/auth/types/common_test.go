package types_test

import (
	"github.com/cosmos/cosmos-sdk/v43/simapp"
)

var (
	ecdc                  = simapp.MakeTestEncodingConfig()
	appCodec, legacyAmino = ecdc.Codec, ecdc.Amino
)
