package transaction

import (
	"fmt"

	"airdrop/config"

	"github.com/ethereum/go-ethereum/common"
)

// SubToAirDrop send default value to airdrop address
func SubToAirDrop() {
	// get subaddress key
	_, subPrivates := getSubAddrKey()
	for _, privateKey := range subPrivates {
		subAddrKeystore := convertToKeystore(privateKey)
		nonce, _ := getCurrentNonce(subAddrKeystore)
		fmt.Println(subAddrKeystore.Address.Hex())
		tx, err := SendTransaction(subAddrKeystore, common.HexToAddress(config.Conf.AirDropAddr), common.Big0, "", nonce)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(tx.Hash().Hex())
		}
	}
}
