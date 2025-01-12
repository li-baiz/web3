package l2

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"    // 提供以太坊地址和哈希类型
	"github.com/ethereum/go-ethereum/ethclient" // 提供以太坊客户端功能
)

type L207 struct {
}

func (f *L207) L207() {
	// 连接到以太坊节点（这里使用的是 Rinkeby 测试网，通过 Infura 提供的 API 端点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	// 将十六进制字符串转换为以太坊地址
	account := common.HexToAddress("0x2D7419E411c895E2A18e86Dca9F1b5A5E91f32a4")

	// 获取当前账户的余额（单位为 wei）
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err) // 如果获取余额失败，记录错误并退出程序
	}
	fmt.Println(balance) // 打印当前余额（单位为 wei）

	// 指定一个区块号（这里使用的是 5532993）
	blockNumber := big.NewInt(7372134)
	// 获取指定区块号时的账户余额（单位为 wei）
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err) // 如果获取余额失败，记录错误并退出程序
	}
	fmt.Println(balanceAt) // 例如：25729324269165216042

	// 将余额从 wei 转换为 ETH
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())                                 // 将余额转换为 big.Float 类型
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18))) // 除以 10^18，将 wei 转换为 ETH
	fmt.Println(ethValue)                                                  // 例如：25.729324269165216041

	// 获取账户的待处理余额（包括未确认的交易）
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err) // 如果获取待处理余额失败，记录错误并退出程序
	}
	fmt.Println(pendingBalance) // 例如：25729324269165216042

	fbalance.SetString(pendingBalance.String())
	fmt.Println(new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18))))
}
