package main

// 通过助记词生成私钥和账户地址，以及生成对应的keystore文件并保存到本地
// 分层确定性（HD）钱包，主私钥可以派生无限个子私钥，子私钥可以派生无限个子地址。恢复钱包后依次创建账户即可恢复所有账户

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// 助记词存在12（128位熵、主流）、15（160）、18（192）、21（224）、24（256位熵、安全性最强）个单词
	// 生成 BIP-39 助记词
	entropy, err := bip39.NewEntropy(128) // 128位熵，生成12个单词的助记词
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("助记词:", mnemonic)

	// 使用助记词生成种子
	seed := bip39.NewSeed(mnemonic, "")

	// 从种子生成主私钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("主私钥:", masterKey.String())

	// 从主私钥派生账户的私钥（根据BIP-44路径）
	// 以太坊 BIP-44 路径: m/44'/60'/0'/0/0
	purpose, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	coinType, _ := purpose.NewChildKey(bip32.FirstHardenedChild + 60)
	account, _ := coinType.NewChildKey(bip32.FirstHardenedChild + 0)
	change, _ := account.NewChildKey(0)
	addressIndex, _ := change.NewChildKey(0)

	privateKey, err := crypto.ToECDSA(addressIndex.Key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("私钥:", hex.EncodeToString(crypto.FromECDSA(privateKey)))

	// 生成公钥并获取地址
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKey).Hex()
	fmt.Println("地址:", address)

	store := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	key, err := store.ImportECDSA(privateKey, "your-password")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("keystore 文件已保存:", filepath.Join("./keystore", key.Address.Hex()))
}
