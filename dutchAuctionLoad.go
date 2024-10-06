package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    dutchAuction "github.com/Test-Go-ETH/Test-Go-ETH/contracts/dutchAuction" // for demo
)

func main() {
    client, err := ethclient.Dial("http://127.0.0.1:8545/")
    if err != nil {
        log.Fatal(err)
    }

    address := common.HexToAddress("0x9e545e3c0baab3e08cdfd552c960a1050f373042")
    instance, err := dutchAuction.NewDutchAuction(address, client)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("contract is loaded")
    tokenPrice, err := instance.GetTokenPrice()
	if err != nil {
        log.Fatal(err)
	}
	fmt.Println("token price is:", tokenPrice)
}