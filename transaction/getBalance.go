package transaction

import (
	"context"
	"fmt"
	"math/big"

	"airdrop/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

// GetTokenTotal get all Token balance total
func GetTokenTotal() {
	_, subPrivates := getSubAddrKey()
	total := big.NewInt(0)
	for _, subPrivateKey := range subPrivates {
		balance, err := getTokenBalance(subPrivateKey)
		if err != nil {
			fmt.Println(err)
		}
		total = total.Add(total, balance)
	}
	fmt.Println(total)
}

func getTokenBalance(privateKey string) (*big.Int, error) {
	contract, err := initContract()
	if err != nil {
		return nil, fmt.Errorf("init-contract: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Conf.Timeout)
	defer cancel()

	mainPK := convertToKeystore(config.Conf.MainPrivateKey)
	subPK := convertToKeystore(privateKey)

	var output *big.Int
	var callOpts = &bind.CallOpts{
		Pending: true,
		From:    mainPK.Address,
		Context: ctx,
	}
	err = contract.Call(callOpts, &output, "balanceOf", subPK.Address)
	if err != nil {
		return nil, fmt.Errorf("get-balance: %s", err.Error())
	}

	return output, nil
}

// GetAddrsBalanceTotal get Addrs balance total
func GetAddrsBalanceTotal() *big.Int {
	i := 0
	_, subPrivates := getSubAddrKey()
	for _, privateKey := range subPrivates {
		if i > 0 {
			// break
		}
		subAddrKeystore := convertToKeystore(privateKey)
		// nonce, _ := getCurrentNonce(subAddrKeystore)
		fmt.Println(GetAddrBalance(subAddrKeystore))
		i++
	}

	// awk '{if(NR%2==0){printf $0 "\n"}else{printf "%s\t", $0}}' file.txt
	// sed '$!N;s/\n/\t/' file.txt
	return common.Big0
}

// GetAddrBalance get balance via private-key
func GetAddrBalance(keystore *keystore.Key) *big.Int {
	GetBalance(keystore.Address)
	return common.Big0
}
