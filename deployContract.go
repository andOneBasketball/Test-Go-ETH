package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 合约字节码，假设你有已编译的合约字节码
var contractBytecode = common.FromHex("608060405234801561001057600080fd5b5060405161013838038061013883398181016040528101906100329190610051565b8060009080519060200190610044929190610051565b50610119565b828054610058906100bd565b90600052602060002090601f01602090048101928261007a57600085556100c1565b82601f1061009357805160ff19168380011785556100c1565b828001600101855582156100c1579182015b828111156100c05782518255916020019190600101906100a5565b5b5090506100ce91906100d2565b5090565b5b808211156100eb5760008160009055506001016100d3565b5090565b600080fd5b6100ff61010e565b61010781610109565b50565b5b5600a165627a7a72305820f89cba2e7c907b4b08c263ff6ba7dd0fcf93871db44c930e5be435e8968c27a40029") // 一个简单的合约字节码

// Create2Address 根据创建者地址、盐值和合约字节码计算合约地址
func Create2Address(deployer common.Address, salt [32]byte, bytecode []byte) common.Address {
	// 将数据拼接
	data := append([]byte{0xff}, deployer.Bytes()...)  // 0xff 常量
	data = append(data, salt[:]...)                    // 盐值
	data = append(data, crypto.Keccak256(bytecode)...) // 合约字节码的 keccak256 哈希值

	// 计算 keccak256 并返回最后 20 字节（地址）
	return common.BytesToAddress(crypto.Keccak256(data)[12:])
}

func main() {
	// 创建以太坊客户端
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 使用私钥生成一个新账户（用来部署合约）
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Cannot cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 设置部署合约时的盐值
	var salt [32]byte
	copy(salt[:], []byte("my_salt_value"))

	// 计算合约地址
	contractAddress := Create2Address(fromAddress, salt, contractBytecode)
	fmt.Println("Expected contract address:", contractAddress.Hex())

	// 获取合约 ABI（在这个例子中我们使用一个空的 ABI）
	parsedABI, err := abi.JSON(bytes.NewReader([]byte("[]"))) // 空的 ABI
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// 获取合约部署者的 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// 设置 gasPrice 和 gasLimit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1)) // 使用 Go-Ethereum 的 bind 包
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // 部署合约无需发送 ETH
	auth.GasLimit = uint64(3000000) // Gas limit
	auth.GasPrice = gasPrice

	// 部署合约（使用 create2 部署）
	tx, _, _, err := bind.DeployContract(auth, parsedABI, contractBytecode, client, salt[:])
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed! tx hash: %s\n", tx.Hex())
}
