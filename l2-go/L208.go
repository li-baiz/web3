package l2

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind" // 提供与智能合约交互的功能
	"github.com/ethereum/go-ethereum/common"            // 提供以太坊地址和哈希类型
	"github.com/ethereum/go-ethereum/ethclient"         // 提供以太坊客户端功能
	token "openWeb/token"                               // 导入本地生成的 ERC-20 合约绑定代码
)

type L208 struct {
}

func (f *L208) L208() {
	// 连接到以太坊节点（这里使用的是 Rinkeby 测试网，通过 Infura 提供的 API 端点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	// 设置代币合约地址（这里使用的是 Golem 代币合约地址）
	tokenAddress := common.HexToAddress("0x18844c9Eed2952De808746A230F3EA147530Be5b")
	// 创建代币合约实例
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err) // 如果创建合约实例失败，记录错误并退出程序
	}

	// 设置查询余额的地址
	address := common.HexToAddress("0x2D7419E411c895E2A18e86Dca9F1b5A5E91f32a4")
	// 查询指定地址的代币余额（单位为 wei）
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err) // 如果查询余额失败，记录错误并退出程序
	}

	// 查询代币的名称
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err) // 如果查询名称失败，记录错误并退出程序
	}

	// 查询代币的符号
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err) // 如果查询符号失败，记录错误并退出程序
	}

	// 查询代币的小数位数
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err) // 如果查询小数位数失败，记录错误并退出程序
	}

	// 打印代币的名称、符号和小数位数
	fmt.Printf("name: %s\n", name)         // 例如："name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // 例如："symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // 例如："decimals: 18"

	// 打印代币余额（单位为 wei）
	fmt.Printf("wei: %s\n", bal) // 例如："wei: 74605500647408739782407023"

	// 将代币余额从 wei 转换为实际值
	fbal := new(big.Float)
	fbal.SetString(bal.String())                                               // 将余额转换为 big.Float 类型
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals)))) // 除以 10^decimals，将 wei 转换为实际值
	fmt.Printf("balance: %f", value)                                           // 例如："balance: 74605500.647409"
}
