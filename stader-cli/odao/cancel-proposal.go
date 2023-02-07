package odao

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/rocket-pool/rocketpool-go/dao"
	"github.com/rocket-pool/rocketpool-go/types"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func cancelProposal(c *cli.Context) error {

	// Get RP client
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(staderClient)
	if err != nil {
		return err
	}

	// Get oracle DAO proposals
	proposals, err := staderClient.TNDAOProposals()
	if err != nil {
		return err
	}

	// Get wallet status
	wallet, err := staderClient.WalletStatus()
	if err != nil {
		return err
	}

	// Get cancelable proposals
	cancelableProposals := []dao.ProposalDetails{}
	for _, proposal := range proposals.Proposals {
		if bytes.Equal(proposal.ProposerAddress.Bytes(), wallet.AccountAddress.Bytes()) && (proposal.State == types.Pending || proposal.State == types.Active) {
			cancelableProposals = append(cancelableProposals, proposal)
		}
	}

	// Check for cancelable proposals
	if len(cancelableProposals) == 0 {
		fmt.Println("No proposals can be cancelled.")
		return nil
	}

	// Get selected proposal
	var selectedProposal dao.ProposalDetails
	if c.String("proposal") != "" {

		// Get selected proposal ID
		selectedId, err := strconv.ParseUint(c.String("proposal"), 10, 64)
		if err != nil {
			return fmt.Errorf("Invalid proposal ID '%s': %w", c.String("proposal"), err)
		}

		// Get matching proposal
		found := false
		for _, proposal := range cancelableProposals {
			if proposal.ID == selectedId {
				selectedProposal = proposal
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("Proposal %d can not be cancelled.", selectedId)
		}

	} else {

		// Prompt for proposal selection
		options := make([]string, len(cancelableProposals))
		for pi, proposal := range cancelableProposals {
			options[pi] = fmt.Sprintf("proposal %d (message: '%s', payload: %s)", proposal.ID, proposal.Message, proposal.PayloadStr)
		}
		selected, _ := cliutils.Select("Please select a proposal to cancel:", options)
		selectedProposal = cancelableProposals[selected]

	}

	// Display gas estimate
	canResponse, err := staderClient.CanCancelTNDAOProposal(selectedProposal.ID)
	if err != nil {
		return err
	}
	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want to cancel proposal %d?", selectedProposal.ID))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Cancel proposal
	response, err := staderClient.CancelTNDAOProposal(selectedProposal.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Canceling proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully cancelled proposal %d.\n", selectedProposal.ID)
	return nil

}
