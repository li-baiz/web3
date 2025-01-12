package l2

import (
	"crypto/ecdsa" // 用于椭圆曲线数字签名算法（ECDSA）
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil" // 提供十六进制编码工具
	"github.com/ethereum/go-ethereum/crypto"         // 提供以太坊加密相关功能
	"golang.org/x/crypto/sha3"                       // 提供 Keccak-256 哈希算法
)

type L204 struct {
}

func (f *L204) L204() {
	// 生成一个新的 ECDSA 私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err) // 如果生成私钥失败，记录错误并退出程序
	}

	// 将私钥转换为字节格式
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 将私钥字节转换为十六进制字符串，并去掉 '0x' 前缀
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 例如：4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d

	// 从私钥中提取公钥
	publicKey := privateKey.Public()
	// 将公钥断言为 *ecdsa.PublicKey 类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey") // 如果断言失败，记录错误并退出程序
	}

	// 将公钥转换为字节格式
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 将公钥字节转换为十六进制字符串，并去掉 '0x04' 前缀（0x04 是未压缩公钥的前缀）
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 例如：9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	// 从公钥生成以太坊地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 例如：0x96216849c49358B10257cb55b28eA603c874b05E

	// 使用 Keccak-256 哈希算法对公钥进行哈希计算
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:]) // 去掉未压缩公钥的前缀（0x04），然后计算哈希
	// 打印完整的哈希值（32 字节）
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:])) // 例如：0x96216849c49358b10257cb55b28ea603c874b05e4d8e4c7f59f6b1b2c6f1d2e3

	// 取哈希值的后 20 字节作为以太坊地址
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 例如：0x96216849c49358b10257cb55b28ea603c874b05e
}
