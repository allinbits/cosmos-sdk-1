package params

import (
	"github.com/cosmos/cosmos-sdk/v42/client"
	"github.com/cosmos/cosmos-sdk/v42/codec"
	"github.com/cosmos/cosmos-sdk/v42/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}
