package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/andOneBasketball/Test-Go-ETH/contracts/AKMToken"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545/")
	if err != nil {
		log.Fatal(err)
	}

	// 私钥初始化
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
	}

	// 公私钥账户地址转换
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fmt.Printf("publicKey is: %x\n", crypto.CompressPubkey(publicKeyECDSA))
	fmt.Printf("privateKey is: %x\n", crypto.FromECDSA(privateKey))

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("address is: %x\n", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("nonce: %d\n", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	// auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	account := common.HexToAddress("0x5639Bc2D96c7bA37EECA625599B183241A2bBE6c")
	address, tx, instance, err := AKMToken.DeployAKMToken(auth, client, fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex())
	log.Printf("tx hash: %s", tx.Hash().Hex())
	log.Printf("chainID: %s", tx.ChainId())
	log.Printf("type: %d", tx.Type())
	log.Printf("gas limit: %d", tx.Gas())
	log.Printf("gas price: %d", tx.GasPrice())
	log.Printf("gas fee cap: %s", tx.GasFeeCap())
	log.Printf("gas tip cap: %s", tx.GasTipCap())

	// 需要重新获取 nonce 才能发交易
	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("second nonce: %d\n", nonce)
	auth.Nonce = big.NewInt(int64(nonce))

	accountB := common.HexToAddress("0xde05927035b51C5f6dE27b427e4649123723e141")
	tx, err = instance.Mint(auth, accountB, big.NewInt(100))
	if err != nil {
		log.Fatal("mint failed: ", err)
	}
	log.Printf("mint tx hash: %s", tx.Hash().Hex())
	log.Printf("mint chainID: %s", tx.ChainId())
	log.Printf("mint type: %d", tx.Type())
	log.Printf("mint gas limit: %d", tx.Gas())
	log.Printf("mint gas price: %d", tx.GasPrice())
	log.Printf("mint gas fee cap: %s", tx.GasFeeCap())
	log.Printf("mint gas tip cap: %s", tx.GasTipCap())

	balance, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("account %s balance is: %v\n", account.Hex(), balance) // 25893180161173005034

	balance, err = instance.BalanceOf(&bind.CallOpts{}, accountB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("accountB %s balance is: %v\n", accountB.Hex(), balance) // 25893180161173005034
}
