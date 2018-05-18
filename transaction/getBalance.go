package transaction

import (
	"context"
	"fmt"
	"math/big"

	"airdrop/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// GetTokenTotal get all Token balance total
func GetTokenTotal() {
	subAddrs := getSubAddrKey()
	total := big.NewInt(0)
	for _, subPrivateKey := range subAddrs {
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
