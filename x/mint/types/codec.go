package types

import (
	"github.com/cosmos/cosmos-sdk/v42/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/v42/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
