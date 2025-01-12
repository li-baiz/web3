package l2

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind" // 提供与智能合约交互的功能
	"github.com/ethereum/go-ethereum/common"            // 提供以太坊地址和哈希类型
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient" // 提供以太坊客户端功能
	"log"
	"math/big"
	store "openWeb/token"
	"strings"
)

const (
	contractAddr = "0x4e4337b7cE4428F0A0FE2Dc551D2bf986f242Ade"
)

type L212 struct{}

func (f *L212) L2121() {
	// 连接到以太坊节点（这里使用的是 Sepolia 测试网，通过 Alchemy 提供的 WebSocket 端点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	storeContract, err := store.NewStore(common.HexToAddress("0x4e4337b7cE4428F0A0FE2Dc551D2bf986f242Ade"), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("560e54255d4eeb85fdd9f9c08c0e26867c16c67d9db61b9149c3aff3189b3671")
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
}

func (f *L212) L2122() {
	// 连接到以太坊节点（这里使用的是 Sepolia 测试网，通过 Alchemy 提供的 WebSocket 端点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	privateKey, err := crypto.HexToECDSA("560e54255d4eeb85fdd9f9c08c0e26867c16c67d9db61b9149c3aff3189b3671")
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 估算 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 准备交易数据
	contractABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		log.Fatal(err)
	}

	methodName := "setItem"
	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key_use_abi"))
	copy(value[:], []byte("demo_save_value_use_abi_11111"))
	input, err := contractABI.Pack(methodName, key, value)

	// 创建交易并签名交易
	chainID := big.NewInt(int64(11155111))
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 100000, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	_, err = waitForReceipt(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// 查询刚刚设置的值
	callInput, err := contractABI.Pack("items", key)
	if err != nil {
		log.Fatal(err)
	}
	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}

	// 解析返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	var unpacked [32]byte
	contractABI.UnpackIntoInterface(&unpacked, "items", result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("is value saving in contract equals to origin value:", unpacked == value)
}

func (f *L212) L2123() {
	// 连接到以太坊节点（这里使用的是 Sepolia 测试网，通过 Alchemy 提供的 WebSocket 端点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/UhL_F3NSNomkbZNG9-ko8Zbokug9OCLs")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并退出程序
	}

	privateKey, err := crypto.HexToECDSA("560e54255d4eeb85fdd9f9c08c0e26867c16c67d9db61b9149c3aff3189b3671")
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 估算 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 准备交易数据
	methodSignature := []byte("setItem(bytes32,bytes32)")
	methodSelector := crypto.Keccak256(methodSignature)[:4]

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key_no_use_abi"))
	copy(value[:], []byte("demo_save_value_no_use_abi_11111"))

	// 组合调用数据
	var input []byte
	input = append(input, methodSelector...)
	input = append(input, key[:]...)
	input = append(input, value[:]...)

	// 创建交易并签名交易
	chainID := big.NewInt(int64(11155111))
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 100000, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	_, err = waitForReceipt(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// 查询刚刚设置的值
	itemsSignature := []byte("items(bytes32)")
	itemsSelector := crypto.Keccak256(itemsSignature)[:4]

	var callInput []byte
	callInput = append(callInput, itemsSelector...)
	callInput = append(callInput, key[:]...)

	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}

	// 解析返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	var unpacked [32]byte
	copy(unpacked[:], result)
	fmt.Println("is value saving in contract equals to origin value:", unpacked == value)
}
