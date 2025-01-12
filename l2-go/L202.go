package l2

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type L202 struct {
}

func (f *L202) L202() {
	// 连接到以太坊节点（这里使用的是 Sepolia 测试网，通过 Alchemy 提供的 API 端点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	// 获取当前网络的 ChainID
	chainID, err := client.ChainID(context.Background())
	fmt.Println(chainID)
	if err != nil {
		log.Fatal(err) // 如果获取 ChainID 失败，记录错误并退出程序
	}

	// 指定一个区块号（这里使用的是 7372134）
	blockNumber := big.NewInt(7372134)
	// 根据区块号获取区块信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err) // 如果获取区块信息失败，记录错误并退出程序
	}
	fmt.Println(block.Hash().Hex())

	// 遍历区块中的所有交易
	for _, tx := range block.Transactions() {
		// 打印交易的哈希值
		fmt.Println(tx.Hash().Hex()) // 例如：0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		// 打印交易的值（转账的以太坊数量）
		fmt.Println(tx.Value().String()) // 例如：100000000000000000
		// 打印交易的 Gas 限制
		fmt.Println(tx.Gas()) // 例如：21000
		// 打印交易的 Gas 价格
		fmt.Println(tx.GasPrice().Uint64()) // 例如：100000000000
		// 打印交易的 Nonce（发送者的交易计数）
		fmt.Println(tx.Nonce()) // 例如：245132
		// 打印交易的数据（通常用于智能合约调用）
		fmt.Println(tx.Data()) // 例如：[]
		// 打印交易的接收者地址
		fmt.Println(tx.To().Hex()) // 例如：0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587
		fmt.Println(tx.Type())

		// 获取交易的发送者地址
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 例如：0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			//log.Fatal(err) // 如果获取发送者地址失败，记录错误并退出程序
			continue
		}

		// 获取交易的收据（包含交易执行结果）
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			//log.Fatal(err) // 如果获取交易收据失败，记录错误并退出程序
			continue
		}

		// 打印交易的状态（1 表示成功，0 表示失败）
		fmt.Println(receipt.Status) // 例如：1
		// 打印交易的日志（通常用于智能合约事件）
		fmt.Println(receipt.Logs) // 例如：[]
		break                     // 只处理第一个交易，然后退出循环
	}

	// 指定一个区块哈希值
	blockHash := common.HexToHash("0x71b9a55b3eedfd05b89c998f1050c3f9d0765bfc501f8e6014b8924701d2c5da")
	// 获取该区块中的交易数量
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err) // 如果获取交易数量失败，记录错误并退出程序
	}

	// 遍历区块中的所有交易
	for idx := uint(0); idx < count; idx++ {
		// 根据索引获取交易
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err) // 如果获取交易失败，记录错误并退出程序
		}

		// 打印交易的哈希值
		fmt.Println(tx.Hash().Hex()) // 例如：0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break                        // 只处理第一个交易，然后退出循环
	}

	// 指定一个交易哈希值
	txHash := common.HexToHash("0xdfa7db5ce07c5728f456cafbe83510d640da5541b77e41f8cef61942af47505a")
	// 根据交易哈希值获取交易信息
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err) // 如果获取交易信息失败，记录错误并退出程序
	}
	// 打印交易是否处于待处理状态
	fmt.Println(isPending) // 例如：false
	// 打印交易的哈希值
	fmt.Println(tx.Hash().Hex()) // 例如：0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
}
