package transaction

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"airdrop/config"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const contractABI = `[{ "constant": false, "inputs": [ { "name": "_to", "type": "address" }, { "name": "_amount", "type": "uint256" } ], "name": "transfer", "outputs": [ { "name": "success", "type": "bool" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" },{ "constant": true, "inputs": [ { "name": "_owner", "type": "address" } ], "name": "balanceOf", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }]`

// inintContract 初始化合约
func initContract() (*bind.BoundContract, error) {
	client, err := InitClient()
	if err != nil {
		return nil, fmt.Errorf("init-client: %s", err.Error())
	}

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, fmt.Errorf("parse-abi: %s", err.Error())
	}

	return bind.NewBoundContract(common.HexToAddress(config.Conf.AirDropAddr), parsedABI, client, client, nil), nil
}

// WithdrawToken withdraw token to specific address
func WithdrawToken() (*types.Transaction, error) {
	_, subPrivates := getSubAddrKey()

	for _, subPrivateKey := range subPrivates {
		tokenBalance, _ := getTokenBalance(subPrivateKey)

		contract, err := initContract()
		if err != nil {
			return nil, fmt.Errorf("init-contract: %s", err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), config.Conf.Timeout)
		defer cancel()

		mainPK := convertToKeystore(config.Conf.MainPrivateKey)
		subPK := convertToKeystore(subPrivateKey)

		currentNonce, _ := getCurrentNonce(subPK)

		auth := bind.NewKeyedTransactor(subPK.PrivateKey)
		auth.Nonce = big.NewInt(int64(currentNonce))
		// auth.GasLimit = config.GasLimit()
		auth.Context = ctx

		fmt.Println(auth.GasPrice)

		fmt.Println(mainPK.Address)
		fmt.Println(common.HexToAddress(config.Conf.MainAddress))

		tx, err := contract.Transact(auth, "transfer", mainPK.Address, tokenBalance)
		break
		fmt.Println(tokenBalance)
		fmt.Println(tx.Hash().Hex())
		fmt.Println(err)
	}

	return nil, nil
}
