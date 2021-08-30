package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/v42/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/v42/types/errors"
	"github.com/cosmos/cosmos-sdk/v42/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/v42/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/v42/x/gov/types"
)

func NewCommunityPoolSpendProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CommunityPoolSpendProposal:
			return keeper.HandleCommunityPoolSpendProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized distr proposal content type: %T", c)
		}
	}
}
