package node

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/rocket-pool/rocketpool-go/utils/eth"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func setWithdrawalAddress(c *cli.Context, withdrawalAddressOrENS string) error {

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

	var withdrawalAddress common.Address
	var withdrawalAddressString string
	if strings.Contains(withdrawalAddressOrENS, ".") {
		response, err := staderClient.ResolveEnsName(withdrawalAddressOrENS)
		if err != nil {
			return err
		}
		withdrawalAddress = response.Address
		withdrawalAddressString = fmt.Sprintf("%s (%s)", withdrawalAddressOrENS, withdrawalAddress.Hex())
	} else {
		withdrawalAddress, err = cliutils.ValidateAddress("withdrawal address", withdrawalAddressOrENS)
		if err != nil {
			return err
		}
		withdrawalAddressString = withdrawalAddress.Hex()
	}

	// Print the "pending" disclaimer
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorYellow := "\033[33m"
	var confirm bool
	fmt.Println("You are about to change your withdrawal address. All future ETH rewards/refunds will be sent there.")
	if !c.Bool("force") {
		confirm = false
		fmt.Println("By default, this will put your new withdrawal address into a \"pending\" state.")
		fmt.Println("Stader will continue to use your old withdrawal address until you confirm that you own the new address via the Stader website.")
		fmt.Println("You will need to use a web3-compatible wallet (such as MetaMask) with your new address to confirm it.")
		fmt.Printf("%sIf you cannot use such a wallet, or if you want to bypass this step and force Stader to use the new address immediately, please re-run this command with the \"--force\" flag.\n\n%s", colorYellow, colorReset)
	} else {
		confirm = true
		fmt.Printf("%sYou have specified the \"--force\" option, so your new address will take effect immediately.\n", colorRed)
		fmt.Printf("Please ensure that you have the correct address - if you do not control the new address, you will not be able to change this once set!%s\n\n", colorReset)
	}

	// Check if the withdrawal address can be set
	canResponse, err := staderClient.CanSetNodeWithdrawalAddress(withdrawalAddress, confirm)
	if err != nil {
		return err
	}

	if confirm {
		// Prompt for a test transaction
		if cliutils.Confirm("Would you like to send a test transaction to make sure you have the correct address?") {
			inputAmount := cliutils.Prompt(fmt.Sprintf("Please enter an amount of ETH to send to %s:", withdrawalAddressString), "^\\d+(\\.\\d+)?$", "Invalid amount")
			testAmount, err := strconv.ParseFloat(inputAmount, 64)
			if err != nil {
				return fmt.Errorf("Invalid test amount '%s': %w\n", inputAmount, err)
			}
			amountWei := eth.EthToWei(testAmount)
			canSendResponse, err := staderClient.CanNodeSend(amountWei, "eth")
			if err != nil {
				return err
			}

			// Assign max fees
			err = gas.AssignMaxFeeAndLimit(canSendResponse.GasInfo, staderClient, c.Bool("yes"))
			if err != nil {
				return err
			}

			if !cliutils.Confirm(fmt.Sprintf("Please confirm you want to send %f ETH to %s.", testAmount, withdrawalAddressString)) {
				fmt.Println("Cancelled.")
				return nil
			}

			sendResponse, err := staderClient.NodeSend(amountWei, "eth", withdrawalAddress)
			if err != nil {
				return err
			}

			fmt.Printf("Sending ETH to %s...\n", withdrawalAddressString)
			cliutils.PrintTransactionHash(staderClient, sendResponse.TxHash)
			if _, err = staderClient.WaitForTransaction(sendResponse.TxHash); err != nil {
				return err
			}

			fmt.Printf("Successfully sent the test transaction.\nPlease verify that your withdrawal address received it before confirming it below.\n\n")
		}
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want to set your node's withdrawal address to %s?", withdrawalAddressString))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Set node's withdrawal address
	response, err := staderClient.SetNodeWithdrawalAddress(withdrawalAddress, confirm)
	if err != nil {
		return err
	}

	fmt.Printf("Setting withdrawal address...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	if !c.Bool("force") {
		stakeUrl := ""
		config, _, err := staderClient.LoadConfig()
		if err == nil {
			stakeUrl = config.Smartnode.GetStakeUrl()
		}
		if stakeUrl != "" {
			fmt.Printf("The node's withdrawal address update to %s is now pending.\n"+
				"To confirm it, please visit the Stader website (%s).", withdrawalAddressString, stakeUrl)
		} else {
			fmt.Printf("The node's withdrawal address update to %s is now pending.\n"+
				"To confirm it, please visit the Stader website.", withdrawalAddressString)
		}
	} else {
		fmt.Printf("The node's withdrawal address was successfully set to %s.\n", withdrawalAddressString)
	}
	return nil

}

func confirmWithdrawalAddress(c *cli.Context) error {

	// Get RP client
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Check if the withdrawal address can be confirmed
	canResponse, err := staderClient.CanConfirmNodeWithdrawalAddress()
	if err != nil {
		return err
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to confirm your node's address as the new withdrawal address?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Confirm node's withdrawal address
	response, err := staderClient.ConfirmNodeWithdrawalAddress()
	if err != nil {
		return err
	}

	fmt.Printf("Confirming new withdrawal address...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("The node's withdrawal address was successfully set to the node address.\n")
	return nil

}
