package l2

import (
	"context"
	"crypto/ecdsa" // 用于椭圆曲线数字签名算法（ECDSA）
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type L205 struct {
}

func (f *L205) L205() {
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
	nonce = uint64(40)

	// 设置转账金额（1 ETH，单位为 wei）
	value := big.NewInt(10000000000000000) // 1 ETH = 10^18 wei
	// 设置 Gas 限制（21000 是以太坊转账的标准 Gas 限制）
	gasLimit := uint64(21000) // 单位：Gas
	// 获取当前建议的 Gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err) // 如果获取 Gas 价格失败，记录错误并退出程序
	}

	// 设置接收者地址
	toAddress := common.HexToAddress("0x23Fa7A7ea35972dD3CEe189703394D8251cAE864")
	// 设置交易数据（普通转账交易不需要数据）
	var data []byte
	// 创建交易对象
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, big.NewInt(0).Add(gasPrice, big.NewInt(int64(10000000))), data)

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
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
