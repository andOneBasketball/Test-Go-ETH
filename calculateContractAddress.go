package main

// 计算合约部署的地址

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

// Keccak256 is a utility function to perform a Keccak-256 hash
func Keccak256(data []byte) []byte {
	hash := crypto.Keccak256Hash(data)
	return hash.Bytes()
}

// CreateAddress generates an Ethereum contract address from an account address and nonce
func CreateAddress(b common.Address, nonce uint64) common.Address {
	// RLP encode the account address and nonce
	data, err := rlp.EncodeToBytes([]interface{}{b, nonce})
	if err != nil {
		log.Fatalf("Failed to RLP encode: %v", err)
	}

	// Perform Keccak256 hashing and take the last 20 bytes (160-bit address)
	hash := Keccak256(data)
	return common.BytesToAddress(hash[12:]) // Take the last 20 bytes
}

func main() {
	// Example usage
	address := common.HexToAddress("f39fd6e51aad88f6f4ce6ab8827279cfffb92266")
	nonce := uint64(29) // Example nonce

	contractAddress := CreateAddress(address, nonce)
	fmt.Println("Contract Address:", contractAddress.Hex())
}
