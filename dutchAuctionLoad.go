package main

// 带确认版的钱包账户转账交易，日志流水打印转账前后的账户余额，以及事件的gas消耗情况

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/andOneBasketball/Test-Go-ETH/contracts/dutchAuction" // for demo
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545/")
	if err != nil {
		log.Fatal(err)
	}
	//account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	address := common.HexToAddress("0x9d4454b023096f34b160d6b654540c56a1f81688")
	instance, err := dutchAuction.NewDutchAuction(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	tokenPrice, err := instance.GetPrice(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("token price is:", tokenPrice)

	// 获取最新区块的哈希
	latestBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get the latest block: %v", err)
	}

	// 获取区块的时间戳
	blockTime := time.Unix(int64(latestBlock.Time()), 0)

	fmt.Printf("Current blockchain time: %s\n", blockTime.Format(time.RFC3339))

	blockNumber := big.NewInt(37)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	blockTime = time.Unix(int64(block.Time()), 0)
	fmt.Printf("34 block in blockchain time: %s\n", blockTime.Format(time.RFC3339))
}
