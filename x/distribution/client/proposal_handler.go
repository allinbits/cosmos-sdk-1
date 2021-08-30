package client

import (
	"github.com/cosmos/cosmos-sdk/v42/x/distribution/client/cli"
	govclient "github.com/cosmos/cosmos-sdk/v42/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
