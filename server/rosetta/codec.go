package rosetta

import (
	"github.com/cosmos/cosmos-sdk/v43/codec"
	codectypes "github.com/cosmos/cosmos-sdk/v43/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/v43/crypto/codec"
	authcodec "github.com/cosmos/cosmos-sdk/v43/x/auth/types"
	bankcodec "github.com/cosmos/cosmos-sdk/v43/x/bank/types"
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
