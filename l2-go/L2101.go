package l2

import (
	"context"
	"crypto/ecdsa" // 用于椭圆曲线数字签名算法（ECDSA）
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind" // 提供与智能合约交互的功能
	"github.com/ethereum/go-ethereum/crypto"            // 提供以太坊加密相关功能
	"github.com/ethereum/go-ethereum/ethclient"         // 提供以太坊客户端功能
	store "openWeb/token"                               // 导入本地生成的合约绑定代码
)

type L2101 struct {
}

func (f *L2101) L2101() {
	// 连接到以太坊节点（这里使用的是 Sepolia 测试网，通过 Alchemy 提供的 WebSocket 端点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	// 从十六进制字符串加载私钥
	privateKey, err := crypto.HexToECDSA("560e54255d4eeb85fdd9f9c08c0e26867c16c67d9db61b9149c3aff3189b3671")
	if err != nil {
		log.Fatal(err) // 如果加载私钥失败，记录错误并退出程序
	}

	// 从私钥中提取公钥
	publicKey := privateKey.Public()
	// 将公钥断言为 *ecdsa.PublicKey 类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey") // 如果断言失败，记录错误并退出程序
	}

	// 从公钥生成以太坊地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 获取发送地址的当前 nonce（交易计数）
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err) // 如果获取 nonce 失败，记录错误并退出程序
	}

	// 获取当前建议的 Gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err) // 如果获取 Gas 价格失败，记录错误并退出程序
	}

	// 获取当前网络的 ChainID
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err) // 如果获取 ChainID 失败，记录错误并退出程序
	}

	// 创建一个授权的交易签名器，包含 ChainID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err) // 如果创建签名器失败，记录错误并退出程序
	}
	// 设置交易的 nonce
	auth.Nonce = big.NewInt(int64(nonce))
	// 设置交易的转账金额（0 wei）
	auth.Value = big.NewInt(0) // in wei
	// 设置交易的 Gas 限制
	auth.GasLimit = uint64(300000) // in units
	// 设置交易的 Gas 价格
	auth.GasPrice = gasPrice

	// 设置合约的初始化参数
	input := "1.0"
	// 部署合约
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err) // 如果部署合约失败，记录错误并退出程序
	}

	// 打印合约地址
	fmt.Println(address.Hex())
	// 打印部署交易的哈希值
	fmt.Println(tx.Hash().Hex())

	// 忽略未使用的实例（用于避免编译警告）
	_ = instance
}
