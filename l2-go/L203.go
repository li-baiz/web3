package l2

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type L203 struct {
}

func (f *L203) L203() {
	// 连接到以太坊节点（这里使用的是 Sepolia 测试网，通过 Alchemy 提供的 API 端点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	// 指定一个区块号（这里使用的是 5671744）
	blockNumber := big.NewInt(7372134)
	// 指定一个区块哈希值
	blockHash := common.HexToHash("0x71b9a55b3eedfd05b89c998f1050c3f9d0765bfc501f8e6014b8924701d2c5da")

	// 根据区块哈希值获取区块中的所有交易收据
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err) // 如果获取交易收据失败，记录错误并退出程序
	}

	// 根据区块号获取区块中的所有交易收据
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err) // 如果获取交易收据失败，记录错误并退出程序
	}

	// 比较通过区块哈希值和区块号获取的交易收据是否相同
	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // true

	// 遍历通过区块哈希值获取的交易收据
	for _, receipt := range receiptByHash {
		// 打印交易的状态（1 表示成功，0 表示失败）
		fmt.Println(receipt.Status) // 例如：1
		// 打印交易的日志（通常用于智能合约事件）
		fmt.Println(receipt.Logs) // 例如：[]
		// 打印交易的哈希值
		fmt.Println(receipt.TxHash.Hex()) // 例如：0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		// 打印交易在区块中的索引
		fmt.Println(receipt.TransactionIndex) // 例如：0
		// 打印合约地址（如果是合约创建交易，否则为 0x0）
		fmt.Println(receipt.ContractAddress.Hex()) // 例如：0x0000000000000000000000000000000000000000
		break                                      // 只处理第一个交易收据，然后退出循环
	}

	// 指定一个交易哈希值
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	// 根据交易哈希值获取交易收据
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err) // 如果获取交易收据失败，记录错误并退出程序
	}

	// 打印交易的状态（1 表示成功，0 表示失败）
	fmt.Println(receipt.Status) // 例如：1
	// 打印交易的日志（通常用于智能合约事件）
	fmt.Println(receipt.Logs) // 例如：[]
	// 打印交易的哈希值
	fmt.Println(receipt.TxHash.Hex()) // 例如：0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	// 打印交易在区块中的索引
	fmt.Println(receipt.TransactionIndex) // 例如：0
	// 打印合约地址（如果是合约创建交易，否则为 0x0）
	fmt.Println(receipt.ContractAddress.Hex()) // 例如：0x0000000000000000000000000000000000000000
}
