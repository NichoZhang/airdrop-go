package transaction

import (
	"airdrop/config"
	"encoding/json"
	"fmt"
	"time"
)

// SendToSub transfer main account's balance to subaddress
func SendToSub() {
	fmt.Println(config.Conf.MainPrivateKey)
	mainKeyStore := convertToKeystore(config.Conf.MainPrivateKey)

	nonce, _ := getCurrentNonce(mainKeyStore)

	// output the pretty keystore struct on console
	indentMainKey, _ := json.MarshalIndent(mainKeyStore, "", "  ")
	fmt.Println(string(indentMainKey))

	addrKey := getSubAddrKey()
	i := 0
	for _, privateKey := range addrKey {
		if i > 0 {
			// break
		}
		subAddrkeystore := convertToKeystore(privateKey)
		receiptAddr := subAddrkeystore.Address
		fmt.Println(privateKey)        // receipt private key
		fmt.Println(receiptAddr.Hex()) // receipt address
		tx, err := SendTransaction(mainKeyStore, receiptAddr, config.Amount(), "", nonce)
		fmt.Println(err)
		fmt.Println(tx.Hash().Hex())
		time.Sleep(500 * time.Nanosecond)
		nonce++
		i++
	}
}
