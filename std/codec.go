package std

import (
	"github.com/cosmos/cosmos-sdk/v42/codec"
	"github.com/cosmos/cosmos-sdk/v42/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/v42/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/v42/types"
	txtypes "github.com/cosmos/cosmos-sdk/v42/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
