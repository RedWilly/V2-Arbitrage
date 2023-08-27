package main

import (
	"bot/abi"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"time"
)

func MakeTx(bot *abi.Bot, pool0, pool1 string) (string, error) {
	// nonce
	log.Println("=== get nonce")
	nonce, err := RpcClient.NonceAt(context.TODO(), FROMADDRESS, nil)
	if err != nil {
		return "", err
	}

	// gas price
	log.Println("=== get gas price")
	gasPriceWei, _ := new(big.Int).SetString(GBotConfig.GasPrice, 10)
	log.Println("=== get gas price success, gas price is", gasPriceWei)

	// 签名
	auth, err := bind.NewKeyedTransactorWithChainID(PRIVATEKEY, new(big.Int).SetInt64(ChainId))
	if err != nil {
		log.Println("=== auth tx error:", err)
		return "", err
	}

	log.Println("=== auth success")
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPriceWei

	tx, err := bot.FlashArbitrage(auth, common.HexToAddress(pool0), common.HexToAddress(pool1))
	if err != nil {
		return "", err
	}

	// 设定时间(防止阻塞)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, RpcClient, tx)
	if err != nil {
		return "", err
	}
	log.Printf("flash arbitrage tx hash is: %s \n", receipt.TxHash.Hex())
	return receipt.TxHash.Hex(), nil
}
