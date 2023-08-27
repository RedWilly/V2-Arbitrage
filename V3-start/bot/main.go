package main

import (
	"bot/abi"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func main() {
	// 配置
	_, err := LoadBotConfig("./config.yaml")
	if err != nil {
		log.Println("load config err =>", err)
		return
	}
	// 节点
	InitRpcClient(GBotConfig.Rpc, GBotConfig.ChainId)
	// 钱包
	InitWallet(GBotConfig.PrivateKey)
	// 合约
	InitContract()

	// 加载需要监听的的 pair对

}

func arbitrageFunc(bot *abi.Bot, aPair *ArbitragePair) error {
	if len(aPair.Pairs) != 2 {
		return fmt.Errorf("pair lenght is not 2")
	}
	pair0, pair1 := aPair.Pairs[0], aPair.Pairs[1]

	profitInfo, err := bot.GetProfit(nil, common.HexToAddress(pair0), common.HexToAddress(pair1))
	if err != nil {
		return err
	}

	// profitInfo 的价值需要大于消耗本币的价值，才有利可图
	// 目前 profit 使用的计价是本币
	if profitInfo.Profit.Cmp(MaxGasCost()) != 1 {
		return fmt.Errorf("no profit")
	}
	// todo 是否是我们设定的最低利润值
	// 发起交易
	log.Printf("call flash arbitrage, profit(wei) => %v wei, token => %v\n", profitInfo.Profit, profitInfo.BaseToken)
	_, err = MakeTx(bot, pair0, pair1)
	if err != nil {
		return err
	}
	return nil
}

// MaxGasCost 套利交易消耗wei的上限
func MaxGasCost() *big.Int {
	gasPriceWei, _ := new(big.Int).SetString(GBotConfig.GasPrice, 10)
	gasLimit, _ := new(big.Int).SetString(GBotConfig.GasLimit, 10)
	gasCostWei := new(big.Int).Mul(gasPriceWei, gasLimit)
	return gasCostWei
}
