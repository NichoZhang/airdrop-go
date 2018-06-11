package main

import (
	"flag"
	"fmt"
	"os"

	"airdrop/account"
	"airdrop/config"
	"airdrop/transaction"
)

func main() {
	gopath := os.Getenv("GOPATH")

	var confPath string
	var mode string
	var network string
	var amount int64

	// createAccount create accounts
	// sendToSub     send eth to subaddress
	// subToAirdrop  send 0 eth to airdrop address
	// withdrawToken withdraw the token from subaddress
	flag.StringVar(&mode, "m", "getTokenTotal", "conf path")
	flag.Int64Var(&amount, "amount", 0, "tx amount")
	flag.StringVar(&network, "network", "ropsten", "network")
	flag.StringVar(&confPath, "c", gopath+"/src/airdrop/conf_", "conf path")
	flag.Parse()

	config.Init(confPath + network + ".json")

	switch mode {
	case "testContract":
		//		transaction.TestContract()
	case "createAccount":
		account.Create()
	case "sendToSub":
		transaction.SendToSub()
	case "subToAirDrop":
		transaction.SubToAirDrop()
	case "getTokenTotal":
		transaction.GetTokenTotal()
	case "getSubBalance":
		transaction.GetAddrsBalanceTotal()
	case "withdrawToken":
		transaction.WithdrawToken()
	case "sendSmartContract":
		transaction.SendSmartContract()
	default:
		fmt.Println("default")
	}
}
