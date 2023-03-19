package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func EstimateAddValidatorKeys(pnr *stader.PermissionlessNodeRegistryContractManager, pubKeys [][]byte, preDepositSignatures [][]byte, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return pnr.PermissionlessNodeRegistryContract.GetTransactionGasInfo(opts, "addValidatorKeys", pubKeys, preDepositSignatures)
}

func AddValidatorKeys(pnr *stader.PermissionlessNodeRegistryContractManager, pubKeys [][]byte, preDepositSignatures [][]byte, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := pnr.PermissionlessNodeRegistry.AddValidatorKeys(opts, pubKeys, preDepositSignatures)
	if err != nil {
		return nil, fmt.Errorf("could not add validator keys: %w", err)
	}

	return tx, nil
}

func GetTotalValidatorKeys(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.GetOperatorTotalKeys(opts, operatorId)
}

func GetValidatorIdByOperatorId(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, validatorIndex *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.ValidatorIdsByOperatorId(opts, operatorId, validatorIndex)
}

func GetValidatorInfo(pnr *stader.PermissionlessNodeRegistryContractManager, validatorId *big.Int, opts *bind.CallOpts) (struct {
	Status               uint8
	Pubkey               []byte
	Signature            []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	InitialBondEth       *big.Int
	DepositTime          *big.Int
	WithdrawnTime        *big.Int
}, error) {
	return pnr.PermissionlessNodeRegistry.ValidatorRegistry(opts, validatorId)
}

func ComputeWithdrawVaultAddress(vfcm *stader.VaultFactoryContractManager, poolType uint8, operatorId *big.Int, validatorCount *big.Int, opts *bind.CallOpts) (common.Address, error) {
	return vfcm.VaultFactory.ComputeWithdrawVaultAddress(opts, poolType, operatorId, validatorCount)
}

func GetValidatorWithdrawalCredential(vfcm *stader.VaultFactoryContractManager, withdrawVaultAddress common.Address, opts *bind.CallOpts) (common.Hash, error) {
	withdrawalCredentials := new(common.Hash)
	if err := vfcm.VaultFactoryContract.Call(opts, withdrawalCredentials, "getValidatorWithdrawCredential", withdrawVaultAddress); err != nil {
		return common.Hash{}, fmt.Errorf("Could not get vault withdrawal credentials: %w", err)
	}

	return *withdrawalCredentials, nil
}
