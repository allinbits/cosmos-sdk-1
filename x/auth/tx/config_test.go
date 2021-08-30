package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/v42/codec"
	codectypes "github.com/cosmos/cosmos-sdk/v42/codec/types"
	"github.com/cosmos/cosmos-sdk/v42/std"
	"github.com/cosmos/cosmos-sdk/v42/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/v42/types"
	"github.com/cosmos/cosmos-sdk/v42/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
