package client

import (
	govclient "github.com/cosmos/cosmos-sdk/v42/x/gov/client"
	"github.com/cosmos/cosmos-sdk/v42/x/upgrade/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal)
