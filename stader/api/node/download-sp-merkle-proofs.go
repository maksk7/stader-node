package node

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	arr_utils "github.com/stader-labs/stader-node/shared/utils/arr-utils"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/urfave/cli"
	"os"
)

func canDownloadSpMerkleProofs(c *cli.Context) (*api.CanDownloadSpMerkleProofsResponse, error) {
	//if err := services.RequireNodeWallet(c); err != nil {
	//	return nil, err
	//}
	//if err := services.RequireNodeRegistered(c); err != nil {
	//	return nil, err
	//}
	//sp, err := services.GetSocializingPoolContract(c)
	//if err != nil {
	//	return nil, err
	//}
	//cfg, err := services.GetConfig(c)
	//if err != nil {
	//	return nil, err
	//}
	//
	//response := api.CanDownloadSpMerkleProofsResponse{}
	//
	//// check if all cycles are present
	//rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	//if err != nil {
	//	return nil, err
	//}
	//currentIndex := rewardDetails.CurrentIndex.Int64()
	//missingCycles := []int64{}
	//// iterate thru all cycles starting from 1
	//for i := int64(1); i < currentIndex; i++ {
	//	isEligible, err := IsEligibleForCycle(c, big.NewInt(i))
	//	if err != nil {
	//		return nil, err
	//	}
	//	if !isEligible {
	//		continue
	//	}
	//
	//	// download all cycles irrespective if the NO claim or not claimed.
	//	cycleMerkleRewardFile := cfg.StaderNode.GetSpRewardCyclePath(i, true)
	//	// check if file exists or not
	//	_, err = os.Stat(cycleMerkleRewardFile)
	//	if !os.IsNotExist(err) && err != nil {
	//		return nil, err
	//	}
	//	if os.IsNotExist(err) {
	//		missingCycles = append(missingCycles, i)
	//	}
	//}
	//
	//// no missing cycles
	//if len(missingCycles) == 0 {
	//	response.NoMissingCycles = true
	//	return &response, nil
	//}

	response := api.CanDownloadSpMerkleProofsResponse{}

	response.MissingCycles = []int64{1, 2, 3, 4, 5}
	response.CurrentCycle = 5

	return &response, nil
}

func downloadSpMerkleProofs(c *cli.Context) (*api.DownloadSpMerkleProofsResponse, error) {
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	//sp, err := services.GetSocializingPoolContract(c)
	//if err != nil {
	//	return nil, err
	//}
	//rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	//if err != nil {
	//	return nil, err
	//}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response := api.DownloadSpMerkleProofsResponse{}

	//currentIndex := rewardDetails.CurrentIndex.Int64()
	missingCycles := []int64{1, 2, 3, 4, 5}
	//// iterate thru all cycles starting from 1
	//for i := int64(1); i < currentIndex; i++ {
	//	isEligible, err := IsEligibleForCycle(c, big.NewInt(i))
	//	if err != nil {
	//		return nil, err
	//	}
	//	if !isEligible {
	//		continue
	//	}
	//
	//	cycleRewardFile := cfg.StaderNode.GetSpRewardCyclePath(i, true)
	//	// check if file exists or not
	//	_, err = os.Stat(cycleRewardFile)
	//	if !os.IsNotExist(err) && err != nil {
	//		return nil, err
	//	}
	//	if os.IsNotExist(err) {
	//		missingCycles = append(missingCycles, i)
	//	}
	//}

	allMerkleProofs, err := stader.GetAllMerkleProofsForOperator(nodeAccount.Address)
	if err != nil {
		return nil, err
	}

	downloadedCycles := []int64{}

	for _, cycleMerkleProof := range allMerkleProofs {
		if !arr_utils.ElementExistsInNumArray(missingCycles, cycleMerkleProof.Cycle) {
			continue
		}

		cycleMerkleProofFile := cfg.StaderNode.GetSpRewardCyclePath(cycleMerkleProof.Cycle, true)
		absolutePathOfProofFile, err := homedir.Expand(cycleMerkleProofFile)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(absolutePathOfProofFile)
		if err != nil {
			return nil, err
		}

		encoder := json.NewEncoder(file)
		err = encoder.Encode(cycleMerkleProof)
		if err != nil {
			return nil, fmt.Errorf("Error encoding JSON: %v", err)
		}

		downloadedCycles = append(downloadedCycles, cycleMerkleProof.Cycle)
	}

	response.DownloadedCycles = downloadedCycles

	return &response, nil
}
