package main

import (
	"bot/abi"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var RpcClient *ethclient.Client

var ChainId int64

// PRIVATEKEY 私钥(重要)
var PRIVATEKEY *ecdsa.PrivateKey

// FROMADDRESS 账号地址
var FROMADDRESS common.Address

// InitRpcClient 连接链网络
func InitRpcClient(rpc string, chainId int64) {
	// 连接rpc节点
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		panic("connect network err")
	}
	RpcClient = conn
	ChainId = chainId
}

// InitWallet 初始化钱包
func InitWallet(privateKey string) {
	// 私钥检查
	log.Println("=== check private key")
	MYPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic("check private key fail")
	}
	publicKey := MYPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	PRIVATEKEY = MYPrivateKey
	FROMADDRESS = crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Println("=== check private key success")
}

var FlashBot *abi.Bot

// InitContract 初始化合约
func InitContract() {
	flashBot, err := abi.NewBot(FROMADDRESS, RpcClient)
	if err != nil {
		panic("new flash bot contract err")
	}

	FlashBot = flashBot
}
