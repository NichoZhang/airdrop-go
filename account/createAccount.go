package account

import (
	"airdrop/config"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

var wg sync.WaitGroup

// Create used to create subaccount
func Create() {
	nowTime := time.Now().Unix()
	filename := "account_test.txt." + strconv.FormatInt(nowTime, 10)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	accountCount := config.Conf.AccountCount
	logChain := make(chan string, accountCount)
	defer close(logChain)
	for i := 0; i < accountCount; i++ {
		go func() {
			key := keystore.NewKeyForDirectICAP(rand.Reader)
			address := key.Address.Hex()
			privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
			logString := address + "\n" + privateKey + "\n"
			logChain <- logString
		}()
	}
	for i := 0; i < accountCount; i++ {
		f.Write([]byte(<-logChain))
	}
}
