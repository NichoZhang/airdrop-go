### 10-17-2018 10:52:32

# airdrop使用说明
* 参数说明
a. -m 使用模式选择，具体有：
	i.   createAccount     创建帐号
	ii.  sendToSub         将主帐号的eth分别发送到子帐号
	iii. subToAirDrop      使用子帐号进行空投
	iv.  getTokenTotal     获取全部帐号中token的总和
	v.   getSubBalance     获取子帐号余额
	vi.  withdrawToken     将智能合约中的eth进行提现
	vii. sendSmartContract 发布智能合约
b.  -network 选择配置文件
	i. 目前已经参与的空投项目有 a 和 b 两种，其中，a 同样使用在 ropsten 的测试环境中，当选择不同的网络环境时，选择的参数设定不同，例如在测试环境中：
```
$ go run main.go -network 'ropsten_{tokenName}_{AccountTotal}'
```

* 查询所有帐号下的token总和
```
$ cd $GOPATH/src/airdrop
$ go run main.go -network 'mainnet_{tokenName}_{AccountTotal}'
```
* 查询子帐号eth金额
```
$ cd $GOPATH/src/airdrop
$ go run main.go -m 'getSubBalance' -network 'mainnet_{tokenName}_{AccountTotal}'
```


### 05-22-2018 10:23:47
add get balance via address

### 05-19-2018 00:37:59
add dynamically adjusting GasPrice

### 05-15-2018 18:41:22
restructure folder
