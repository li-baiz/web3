package l2

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types" // 提供以太坊区块和交易类型
	"github.com/ethereum/go-ethereum/ethclient"  // 提供以太坊客户端功能
)

type L209 struct {
}

func (f *L209) L209() {
	// 连接到以太坊节点（这里使用的是 Rinkeby 测试网，通过 Infura 提供的 API 端点）
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	// 创建一个通道，用于接收新区块的头部信息
	headers := make(chan *types.Header)
	// 订阅新区块头部事件
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err) // 如果订阅失败，记录错误并退出程序
	}

	// 进入无限循环，监听新区块事件
	for {
		select {
		case err := <-sub.Err(): // 处理订阅错误
			log.Fatal(err) // 如果发生错误，记录错误并退出程序
		case header := <-headers: // 接收到新区块的头部信息
			// 打印新区块的哈希值
			fmt.Println(header.Hash().Hex()) // 例如：0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			// 根据区块哈希值获取完整的区块信息
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err) // 如果获取区块信息失败，记录错误并退出程序
			}

			// 打印区块的哈希值
			fmt.Println(block.Hash().Hex()) // 例如：0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			// 打印区块的区块号
			fmt.Println(block.Number().Uint64()) // 例如：3477413
			// 打印区块的时间戳
			fmt.Println(block.Time()) // 例如：1529525947
			// 打印区块的 Nonce 值
			fmt.Println(block.Nonce()) // 例如：130524141876765836
			// 打印区块中的交易数量
			fmt.Println(len(block.Transactions())) // 例如：7
		}
	}
}
