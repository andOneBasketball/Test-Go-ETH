package main

// 分析 Uniswap 交易对合约的事件日志和交易详情

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/andOneBasketball/Test-Go-ETH/contracts/UniswapV2/UniswapV2Factory"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const uniswapV2PairABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount0In","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1In","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount0Out","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1Out","type":"uint256"},{"indexed":true,"internalType":"address","name":"to","type":"address"}],"name":"Swap","type":"event"}]`

func main() {
	// 连接到以太坊客户端（例如 Infura 或本地节点）
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/hU4c850dMHnn1PfF2SRLLjtTps_kMmmG")
	if err != nil {
		log.Fatal(err)
	}

	uniswapV2FactoryBlockBegin := int64(10000835)
	uniswapV2FactoryAddress := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	uniswapV2FactoryInstance, err := UniswapV2Factory.NewUniswapV2Factory(uniswapV2FactoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	feeTo, err := uniswapV2FactoryInstance.FeeTo(&bind.CallOpts{})
	if err != nil {
		log.Fatal("get feeTo failed, err: %v", err)
	}
	log.Printf("FeeTo: %s", feeTo.String())

	feeToSetter, err := uniswapV2FactoryInstance.FeeToSetter(&bind.CallOpts{})
	if err != nil {
		log.Fatal("get feeToSetter failed, err: %v", err)
	}
	log.Printf("FeeToSetter: %s", feeToSetter.String())

	// 获取所有交易对地址
	allPairLength, err := uniswapV2FactoryInstance.AllPairsLength(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("All Pairs Length: %v", allPairLength.Int64())
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("current block number is %d", header.Number.Uint64())

	blockBatch := int64(2000)
	addressBatch := int64(500)
	j := int64(0)
	k := int64(0)
	addresses := make([]common.Address, allPairLength.Int64())
	for i := int64(0); i < allPairLength.Int64(); i++ {
		pairAddress, err := uniswapV2FactoryInstance.AllPairs(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			log.Fatal(err)
			continue
		}
		addresses[i] = pairAddress

		if i != 0 && i%addressBatch == 0 {
			log.Printf("the %dth Pair Address: %s", i, pairAddress.String())
			for {
				header, err := client.HeaderByNumber(context.Background(), nil)
				if err != nil {
					log.Fatal(err)
				}
				if uniswapV2FactoryBlockBegin+(k+1)*blockBatch > int64(header.Number.Uint64()) {
					log.Printf("the current block number is %d", header.Number.Uint64())
					break
				}

				beginBlock := uniswapV2FactoryBlockBegin + k*blockBatch
				endBlock := uniswapV2FactoryBlockBegin + (k+1)*blockBatch
				// log.Printf("begin find swap in block in %d to %d, the current block number is %d", beginBlock, endBlock, header.Number.Uint64())
				if k != 0 && k*blockBatch%3000000 == 0 {
					log.Printf("begin find swap in block in %d to %d, the current block number is %d", beginBlock, endBlock, header.Number.Uint64())
				}
				query := ethereum.FilterQuery{
					FromBlock: big.NewInt(beginBlock), // 指定uniswapV2Factory最早的区块为起始区块
					ToBlock:   big.NewInt(endBlock),   // 指定结束区块
					Addresses: addresses[j : j+addressBatch],
				}
				j += addressBatch
				k++

				// 获取事件日志
				logs, err := client.FilterLogs(context.Background(), query)
				if err != nil {
					log.Fatal(err)
				}

				// 解析合约 ABI
				contractAbi, err := abi.JSON(strings.NewReader(uniswapV2PairABI))
				if err != nil {
					log.Fatal(err)
				}

				for _, vLog := range logs {
					fmt.Printf("Log Block Number: %v\n", vLog.BlockNumber)
					fmt.Printf("Log Index: %v\n", vLog.Index)

					// 解析日志中的 Swap 事件
					var swapEvent struct {
						Sender     common.Address
						Amount0In  *big.Int
						Amount1In  *big.Int
						Amount0Out *big.Int
						Amount1Out *big.Int
						To         common.Address
					}

					err := contractAbi.UnpackIntoInterface(&swapEvent, "Swap", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Printf("Swap Event: %+v\n", swapEvent)
				}
			}
		}
	}

	/*
		// 指定 Uniswap 交易对合约地址
		pairAddress := common.HexToAddress("0xE4610c983877480a50eeA0D53E313eDa423eC678") // 替换为实际的 Uniswap Pair 地址

		// 创建查询参数
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(12345678), // 指定起始区块
			ToBlock:   big.NewInt(12345688), // 指定结束区块
			Addresses: []common.Address{pairAddress},
		}

		// 获取事件日志
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		// 解析合约 ABI
		contractAbi, err := abi.JSON(strings.NewReader(uniswapV2PairABI))
		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {
			fmt.Printf("Log Block Number: %v\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %v\n", vLog.Index)

			// 解析日志中的 Swap 事件
			var swapEvent struct {
				Sender     common.Address
				Amount0In  *big.Int
				Amount1In  *big.Int
				Amount0Out *big.Int
				Amount1Out *big.Int
				To         common.Address
			}

			err := contractAbi.UnpackIntoInterface(&swapEvent, "Swap", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Swap Event: %+v\n", swapEvent)
		}

		// 根据交易哈希获取交易详情
		txHash := common.HexToHash("0xYourTxHash") // 替换为实际的交易哈希
		tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
		if err != nil {
			log.Fatal(err)
		}

		if isPending {
			fmt.Println("Transaction is still pending")
		} else {
			fmt.Printf("Transaction details: %+v\n", tx)
		}
	*/
}
