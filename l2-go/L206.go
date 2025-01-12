package l2

import (
	"context"
	"crypto/ecdsa" // 用于椭圆曲线数字签名算法（ECDSA）
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/sha3" // 提供 Keccak-256 哈希算法

	"github.com/ethereum/go-ethereum"                // 提供以太坊核心功能
	"github.com/ethereum/go-ethereum/common"         // 提供以太坊地址和哈希类型
	"github.com/ethereum/go-ethereum/common/hexutil" // 提供十六进制编码工具
	"github.com/ethereum/go-ethereum/core/types"     // 提供以太坊交易类型
	"github.com/ethereum/go-ethereum/crypto"         // 提供以太坊加密相关功能
	"github.com/ethereum/go-ethereum/ethclient"      // 提供以太坊客户端功能
)

type L206 struct {
}

func (f *L206) L206() {
	// 连接到以太坊节点（这里使用的是 Rinkeby 测试网，通过 Infura 提供的 API 端点）
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

	// 设置转账金额（0 ETH，单位为 wei）
	value := big.NewInt(0) // 0 ETH
	// 获取当前建议的 Gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err) // 如果获取 Gas 价格失败，记录错误并退出程序
	}

	// 设置接收者地址
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	// 设置代币合约地址
	tokenAddress := common.HexToAddress("0x18844c9Eed2952De808746A230F3EA147530Be5b")

	// 定义 transfer 函数的签名（transfer(address,uint256)）
	transferFnSignature := []byte("transfer(address,uint256)")
	// 使用 Keccak-256 哈希算法计算函数签名的哈希值
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	// 取哈希值的前 4 字节作为方法 ID
	methodID := hash.Sum(nil)[:4]

	fmt.Println(hexutil.Encode(methodID)) // 例如：0xa9059cbb

	// 将接收者地址填充到 32 字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 例如：0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	// 设置转账的代币数量（1000 个代币）
	amount := new(big.Int)
	amount.SetString("100000000000000000", 10) // 1000 tokens
	// 将代币数量填充到 32 字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 例如：0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	// 构造交易数据（方法 ID + 接收者地址 + 代币数量）
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 估算交易的 Gas 限制
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err) // 如果估算 Gas 限制失败，记录错误并退出程序
	}
	fmt.Println(gasLimit) // 例如：23256

	gasLimit = 300000
	// 创建交易对象
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	// 获取当前网络的 ChainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err) // 如果获取 ChainID 失败，记录错误并退出程序
	}

	// 使用私钥对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err) // 如果签名失败，记录错误并退出程序
	}

	// 发送已签名的交易到以太坊网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err) // 如果发送交易失败，记录错误并退出程序
	}

	// 打印已发送交易的哈希值
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // 例如：tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
