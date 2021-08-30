package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/v42/types/errors"
)

const StoreCodespace = "store"

var (
	ErrInvalidProof = sdkerrors.Register(StoreCodespace, 2, "invalid proof")
)
