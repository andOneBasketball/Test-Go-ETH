/*
ethereum.enable()
account = "0x5639Bc2D96c7bA37EECA625599B183241A2bBE6c"
message = "Hello, please sign this message to authenticate."
ethereum.request({method: "personal_sign", params: [account, message]})
0x956fb26bf2700ef02b584c348d98bc6948ba09966e6841fe9fe08ed0cb1f0a02789c014bfe6e9d5f9827260f9d983c1c2b046fe69e03ab7d7880c21bc2c50dd11b
*/

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func verifySignature(address string, message string, signature string) (bool, error) {
	// Step 1: Format the message with Ethereum's prefix
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)

	// Step 2: Hash the prefixed message
	messageHash := crypto.Keccak256([]byte(prefixedMessage))

	// Step 3: Decode the signature
	signatureBytes, err := hex.DecodeString(signature[2:]) // Remove "0x" prefix
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %v", err)
	}
	if len(signatureBytes) != 65 {
		return false, fmt.Errorf("invalid signature length")
	}

	// Extract r, s, and v from the signature
	//r := signatureBytes[:32]
	//s := signatureBytes[32:64]
	v := signatureBytes[64]

	// Adjust v value for recovery (27 or 28 are valid)
	if v < 27 {
		v += 27
	}

	// Append v back into the signature
	fullSignature := append(signatureBytes[:64], v-27) // Remove 27 offset for go-ethereum

	// Step 4: Recover the public key
	publicKey, err := crypto.SigToPub(messageHash, fullSignature)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %v", err)
	}

	// Step 5: Derive the address from the public key
	recoveredAddress := crypto.PubkeyToAddress(*publicKey)

	// Step 6: Compare the recovered address with the expected address
	log.Printf("address is %s", recoveredAddress.Hex())
	if bytes.Equal(recoveredAddress.Bytes(), common.HexToAddress(address).Bytes()) {
		return true, nil
	}
	return false, fmt.Errorf("address mismatch: got %s, expected %s", recoveredAddress.Hex(), address)
}

func main() {
	// Example data
	account := "0x5639Bc2D96c7bA37EECA625599B183241A2bBE6c"
	message := "127.0.0.10x5639Bc2D96c7bA37EECA625599B183241A2bBE6cSign in to Anime Petitionlogin59144696942531731825448"
	signature := "0xc43b720062b2db3b20204ddb07ecdfc384e7814d31b7351256f4261c7d311eec7588aaa708ae5cc80e9c024b9ace24b522f9f90e6675e0455fc1c96c6efa37e51c"

	// Verify the signature
	valid, err := verifySignature(account, message, signature)
	if err != nil {
		log.Fatalf("Error verifying signature: %v", err)
	}
	if valid {
		fmt.Println("Signature is valid!")
	} else {
		fmt.Println("Signature is invalid!")
	}
}
