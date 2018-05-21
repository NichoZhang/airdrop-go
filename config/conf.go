package config

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"time"
)

// Conf config setting
var Conf confJSON

// confJSON struct used to parse json
type confJSON struct {
	AccountURL       string        `json:"account_url"`
	URL              string        `json:"url"`
	AddrPath         string        `json:"addr_path"`
	Timeout          time.Duration `json:"timeout"`
	MainPrivateKey   string        `json:"main_private_key"`
	MainAddress      string        `json:"main_address"`
	SubPrivateKey    string        `json:"sub_private_key"`
	AirDropAddr      string        `json:"air_drop_addr"`
	AccountCount     int           `json:"account_count"`
	DefaultAmount    int64         `json:"default_amount"`
	DefaultGasPrice  int64         `json:"default_gas_price"`
	DefaultGasLimit  uint64        `json:"default_gas_limit"`
	UseAdditionPrice bool          `json:"use_addition_price"`
	AdditionGasPrice int64         `json:"addition_gas_price"`
}

// Init initial the config
func Init(filePath string) {
	jsonStr, _ := ioutil.ReadFile(filePath)
	json.Unmarshal(jsonStr, &Conf)
}

// GasLimit get gas limit
func GasLimit() uint64 {
	return Conf.DefaultGasLimit
}

// GasPrice get gas price
func GasPrice() *big.Int {
	return MWei(Conf.DefaultGasPrice)
}

// AdditionGasPrice addition gas price
func AdditionGasPrice() *big.Int {
	return MWei(Conf.AdditionGasPrice)
}

// Amount set amount
func Amount() *big.Int {
	return MilliEther(Conf.DefaultAmount)
}

// Wei return (num * 10**exp)
func wei(num int64, exp int64) *big.Int {
	exp10 := new(big.Int).Exp(big.NewInt(10), big.NewInt(exp), nil)
	return new(big.Int).Mul(big.NewInt(num), exp10)
}

// Wei Ethereum Base Unit
func Wei(amount int64) *big.Int {
	return wei(amount, 0)
}

// KWei Babbage = 10^3 Wei
func KWei(amount int64) *big.Int {
	return wei(amount, 3)
}

// MWei Lovelace = 10^6 Wei
func MWei(amount int64) *big.Int {
	return wei(amount, 6)
}

// GWei Shannon = 10^9 Wei
func GWei(amount int64) *big.Int {
	return wei(amount, 9)
}

// MicroEther Szabo = 10^12 Wei
func MicroEther(amount int64) *big.Int {
	return wei(amount, 12)
}

// MilliEther Finney = 10^15 Wei
func MilliEther(amount int64) *big.Int {
	return wei(amount, 15)
}

// Ether = 10^18 Wei
func Ether(amount int64) *big.Int {
	return wei(amount, 18)
}
