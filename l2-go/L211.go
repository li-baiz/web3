package l2

import (
	"github.com/ethereum/go-ethereum/common"    // 提供以太坊地址和哈希类型
	"github.com/ethereum/go-ethereum/ethclient" // 提供以太坊客户端功能
	"log"
	store "openWeb/token"
)

type L211 struct{}

func (f *L211) L211() {
	// 连接到以太坊节点（这里使用的是 Sepolia 测试网，通过 Alchemy 提供的 WebSocket 端点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	storeContract, err := store.NewStore(common.HexToAddress("0x5eadb131ab0b961d51cd401f2a09da45e2e4b53a"), client)
	if err != nil {
		log.Fatal(err)
	}

	_ = storeContract
}
