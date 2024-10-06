package main

import (
    "context"
    "fmt"
    "log"
    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("http://127.0.0.1:8545/")
    if err != nil {
        log.Fatal(err)
    }

    account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balance) // 25893180161173005034

    blockNumber := big.NewInt(37)
    balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balanceAt) // 25729324269165216042

    fbalance := new(big.Float)
    fbalance.SetString(balanceAt.String())
    ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
    fmt.Println(ethValue) // 25.729324269165216041

    // PendingBalanceAt 获取的是挂起的余额
    pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
    fmt.Println(pendingBalance) // 25729324269165216042
}