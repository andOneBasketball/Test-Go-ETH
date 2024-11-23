package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const Confirmations = 6 // 等待 6 个确认区块

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个用于监听新区块的通道
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// 缓存区块及其确认数
	blockCache := make(map[uint64]*types.Block)

	fmt.Println("Listening for new blocks...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			blockNumber := block.Number().Uint64()
			fmt.Printf("New block received: %d\n", blockNumber)

			// 将新接收到的区块缓存
			blockCache[blockNumber] = block

			// 检查之前未确认的区块是否已经得到了6个确认
			for blockNumber, cachedBlock := range blockCache {
				latestBlockNumber, err := client.BlockNumber(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				if latestBlockNumber >= blockNumber+Confirmations {
					// 该区块已确认6个区块，可以安全处理
					fmt.Printf("Block %d has been confirmed after 6 blocks\n", blockNumber)
					// 在这里处理业务逻辑
					delete(blockCache, blockNumber) // 移除已处理的区块
				}
			}
		}
		// 暂停短时间，以防止频繁查询最新区块号
		time.Sleep(time.Second * 2)
	}
}
