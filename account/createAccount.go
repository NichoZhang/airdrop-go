package account

import (
	"airdrop/config"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

var wg sync.WaitGroup

// Create used to create subaccount
func Create() {
	// filename := "account_test.txt"
	// f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	accountCount := config.Conf.AccountCount
	for i := 0; i < accountCount; i++ {
		wg.Add(1)
		go func() {
			key := keystore.NewKeyForDirectICAP(rand.Reader)
			fmt.Println(key.Address.Hex())
			fmt.Println(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)))
			wg.Done()
		}()
	}
	wg.Wait()
}

// func writeToFile(i int, f *os.File, w *sync.WaitGroup) {
// 	w.Done()
// }
