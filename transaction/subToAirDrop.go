package transaction

import (
	"fmt"

	"airdrop/config"

	"github.com/ethereum/go-ethereum/common"
)

// SubToAirDrop send default value to airdrop address
func SubToAirDrop() {
	// get subaddress key
	subAddrs := getSubAddrKey()
	for _, privateKey := range subAddrs {
		subAddrKeystore := convertToKeystore(privateKey)
		nonce, _ := getCurrentNonce(subAddrKeystore)
		tx, err := SendTransaction(subAddrKeystore, common.HexToAddress(config.Conf.AirDropAddr), common.Big0, "", nonce)
		fmt.Println(err)
		fmt.Println(tx.Hash().Hex())
	}
}
