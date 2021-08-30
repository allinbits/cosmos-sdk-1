package rosetta

import (
	"github.com/cosmos/cosmos-sdk/v42/codec"
	codectypes "github.com/cosmos/cosmos-sdk/v42/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/v42/crypto/codec"
	authcodec "github.com/cosmos/cosmos-sdk/v42/x/auth/types"
	bankcodec "github.com/cosmos/cosmos-sdk/v42/x/bank/types"
)

// MakeCodec generates the codec required to interact
// with the cosmos APIs used by the rosetta gateway
func MakeCodec() (*codec.ProtoCodec, codectypes.InterfaceRegistry) {
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	authcodec.RegisterInterfaces(ir)
	bankcodec.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)

	return cdc, ir
}
