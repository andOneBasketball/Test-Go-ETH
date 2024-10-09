package main

// 代币交易

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	//"os"
)

func main() {
	// apiKey := os.Getenv("INFURA_API_KEY")
	url := "http://127.0.0.1:8545"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("could not connect to Infura with ethclient: %s", err)
	}
	ctx := context.Background()
	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("get chainID error: %s", err)
	}

	pk, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatalf("load private key error: %s", err)
	}
	address := crypto.PubkeyToAddress(pk.PublicKey)
	log.Printf("account load success, address: %s", crypto.PubkeyToAddress(pk.PublicKey))

	// 获取账户余额
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s balance: %s", address, balance)

	nonce, err := client.NonceAt(ctx, address, nil)
	if err != nil {
		log.Fatalf("get nonce error: %v", err)
	}

	log.Printf("nonce: %d", nonce)
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("get header error: %v", err)
	}
	log.Printf("base fee: %s", header.BaseFee)
	// log.Printf("header data: %#v", header)
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatalf("get SuggestGasTipCap error: %v", err)
	}
	log.Printf("Suggested GasTipCap(maxPriorityFeePerGas): %s", gasTipCap)

	to := common.HexToAddress("0x5639Bc2D96c7bA37EECA625599B183241A2bBE6c")
	balance, err = client.BalanceAt(context.Background(), to, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s balance: %s", to.Hex(), balance)
	amount := big.NewInt(100_000_000_000_000_000) // 0.1 ether

	txData := &types.DynamicFeeTx{
		ChainID: chainId,
		Nonce:   nonce,
		To:      &to,
		Value:   amount,
		Gas:     21000,

		GasFeeCap: header.BaseFee,       //maxFeePerGas 创建交易的用户愿意为每gas支付的最高费用
		GasTipCap: big.NewInt(10000000), //maxPriorityFeePerGas 创建交易的用户在基础费用（base fee）之外愿意支付的最高优先费，必须小于 maxFeePerGas
	}
	signedTx, err := types.SignNewTx(pk, types.LatestSignerForChainID(chainId), txData)
	if err != nil {
		log.Fatalf("sign tx error: %v", err)
	}
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatalf("sign tx error: %v", err)
	}

	// 获取账户余额
	balance, err = client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("after transfer, %s balance: %s", address, balance)

	balance, err = client.BalanceAt(context.Background(), to, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("after transfer, %s balance: %s", to.Hex(), balance)
}
