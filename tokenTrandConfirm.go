package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// https://app.infura.io/ 站点账户的API Key
	apiKey := os.Getenv("INFURA_API_KEY")
	url := "https://sepolia.infura.io/v3/" + apiKey
	ec, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("could not connect to Infura with ethclient: %s", err)
	}
	ctx := context.Background()

	// 加载钱包账户的私钥
	pk, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("load private key error: %s", err)
	}

	to := common.HexToAddress("0xde05927035b51C5f6dE27b427e4649123723e141")
	amount := big.NewInt(100_000_000_000_000_00) // 0.01 ether

	// 获取账户余额
	address := crypto.PubkeyToAddress(pk.PublicKey)
	balance, err := ec.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(address, " get balance error: ", err)
	}
	log.Printf("%s balance: %s", address, balance)
	balance, err = ec.BalanceAt(context.Background(), to, nil)
	if err != nil {
		log.Fatal(to, " get balance error: ", err)
	}
	log.Printf("%s balance: %s", to, balance)

	// 发送交易
	hash, err := sendEth(ctx, ec, pk, to, amount)
	if err != nil {
		log.Fatalf("send Ether error: %s", err)
	}
	log.Printf("tx sent, hash: %s", hash)
	if err = waitConfirm(ctx, ec, *hash, time.Minute*10); err != nil {
		log.Fatalf("wait confirmation error, please check the tx by yourself: %s", err)
	}
	log.Printf("tx %s confirmed", hash)

	// 获取转账后的余额
	balance, err = ec.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(address, " get balance error: ", err)
	}
	log.Printf("after tx, %s balance: %s", address, balance)
	balance, err = ec.BalanceAt(context.Background(), to, nil)
	if err != nil {
		log.Fatal(to, " get balance error: ", err)
	}
	log.Printf("after tx, %s balance: %s", to, balance)
}

func sendEth(ctx context.Context, ec *ethclient.Client, pk *ecdsa.PrivateKey, to common.Address, amount *big.Int) (*common.Hash, error) {
	chainId, err := ec.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(pk.PublicKey)
	log.Printf("account: %s", address)

	nonce, err := ec.NonceAt(ctx, address, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("nonce: %d", nonce)
	header, err := ec.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("base fee: %s", header.BaseFee)
	gasTipCap, err := ec.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}
	gasPrice, err := ec.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("Suggested GasTipCap(maxPriorityFeePerGas): %s, gas price: %s", gasTipCap, gasPrice.String())
	txData := &types.DynamicFeeTx{
		ChainID: chainId,
		Nonce:   nonce,
		To:      &to,
		Value:   amount,
		Gas:     21000,

		GasFeeCap: header.BaseFee, // 可能小于实际gas消耗
		//GasTipCap: gasTipCap,
		GasTipCap: big.NewInt(1),
	}
	signedTx, err := types.SignNewTx(pk, types.LatestSignerForChainID(chainId), txData)
	log.Printf("tx hash: %s, chainID: %s, types: %d, gas limit: %d, gas price: %d, gas fee cap: %s, gas tip cap: %s", signedTx.Hash().Hex(), signedTx.ChainId(), signedTx.Type(), signedTx.Gas(), signedTx.GasPrice(), signedTx.GasFeeCap(), signedTx.GasTipCap())
	if err != nil {
		return nil, err
	}
	err = ec.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}
	hash := signedTx.Hash()
	return &hash, nil
}

func waitConfirm(ctx context.Context, ec *ethclient.Client, txHash common.Hash, timeout time.Duration) error {
	pending := true
	attempts := 0
	maxAttempts := 3
	for pending {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(timeout):
			return errors.New("timeout")
		case <-time.After(time.Second):
			// 仅确认交易是否完成
			_, isPending, err := ec.TransactionByHash(ctx, txHash)
			if err != nil {
				if err.Error() == "not found" && attempts < maxAttempts {
					attempts++
					continue
				}
				return err
			}
			if !isPending {
				pending = false // break `for`
			}
		}
	}
	// 获取交易详情
	receipt, err := ec.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		msg := fmt.Sprintf("transaction reverted, hash %s", receipt.TxHash.String())
		return errors.New(msg)
	}
	log.Printf("transfor confirmed, tx hash: %s, types: %d, gas used: %d, effective gas price: %s, blob gas used: %d, blob gas price: %s, contract address: %s", receipt.TxHash.Hex(), receipt.Type, receipt.GasUsed, receipt.EffectiveGasPrice.String(), receipt.BlobGasPrice, receipt.BlobGasPrice.String(), receipt.ContractAddress.Hex())
	return nil
}
