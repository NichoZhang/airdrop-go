package transaction

import (
	"bufio"
	"context"
	"fmt"
	"math/big"
	"os"

	"airdrop/config"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *Client

// Client defined struct
type Client struct {
	ethclient.Client
	NetworkID *big.Int
	GasPrice  *big.Int
	GasLimit  uint64
}

// InitClient initial ethcli
func InitClient() (*Client, error) {
	if client != nil {
		return client, nil
	}

	cli, err := ethclient.Dial(config.Conf.URL)
	if err != nil {
		return nil, fmt.Errorf("ethclient-dial: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Conf.Timeout)
	defer cancel()
	fmt.Println(ctx)

	networkID, err := cli.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get-chainid: %s", err.Error())
	}

	ctx, cancel = context.WithTimeout(context.Background(), config.Conf.Timeout)
	defer cancel()

	gasPrice := config.GasPrice()
	if config.Conf.UseAdditionPrice {
		suggestGasPrice, err := cli.SuggestGasPrice(ctx)
		if err != nil {
			return nil, fmt.Errorf("suggest-gas-price: %s", err.Error())
		}
		additionPrice := big.NewInt(0)
		gasPrice = additionPrice.Add(suggestGasPrice, config.AdditionGasPrice())
	}

	client = &Client{
		Client:    *cli,
		NetworkID: networkID,
		GasPrice:  gasPrice,
		GasLimit:  config.GasLimit(),
	}

	return client, nil
}

// getCurrentNonce get current max nonce
// because nonce or pendding will trigger panic
func getCurrentNonce(from *keystore.Key) (uint64, error) {
	client, err := InitClient()
	if err != nil {
		return 0, fmt.Errorf("init-client: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Conf.Timeout)
	defer cancel()

	nonce, err := client.NonceAt(ctx, from.Address, nil)
	return nonce, nil
}

func getSubAddrKey() (map[string]string, []string) {
	f, err := os.Open(config.Conf.AddrPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewScanner(f)
	i := 0
	tmp := ""
	addrKey := map[string]string{}
	var addrAry []string
	for rd.Scan() {
		if i%2 == 0 {
			tmp = rd.Text()
		} else {
			addrAry = append(addrAry, rd.Text())
			addrKey[tmp] = rd.Text()
		}
		i++
	}
	return addrKey, addrAry
}

// initKeystore initial publickey and privatekey
func convertToKeystore(privateKeyStr string) *keystore.Key {
	privateKey, _ := crypto.ToECDSA((common.FromHex(privateKeyStr)))
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)

	return &keystore.Key{
		Address:    address,
		PrivateKey: privateKey,
	}
}

// GetBalance get addr balance
func GetBalance(addr common.Address) {
	client, err := InitClient()
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.Conf.Timeout)
	defer cancel()

	fmt.Println(addr.Hex())
	balance, _ := client.BalanceAt(ctx, addr, nil)
	fmt.Printf("%018.18f", float64(balance.Int64())/float64(config.Ether(1).Int64()))
}

// SendTransaction transaction function
func SendTransaction(from *keystore.Key, to common.Address, amount *big.Int, data string, nonce uint64) (*types.Transaction, error) {
	client, err := InitClient()
	if err != nil {
		return nil, fmt.Errorf("init-client: %s", err.Error())
	}

	tx := types.NewTransaction(nonce, to, amount, client.GasLimit, client.GasPrice, []byte(data))
	tx, err = types.SignTx(tx, types.FrontierSigner{}, from.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("sign-transaction: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Conf.Timeout)
	defer cancel()

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("send-transaction: %s", err.Error())
	}

	fmt.Println(tx)
	return tx, nil
}
