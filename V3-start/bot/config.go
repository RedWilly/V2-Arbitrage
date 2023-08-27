package main

import (
	"errors"
	"github.com/spf13/viper"
)

var GBotConfig *BotConfig

type BotConfig struct {
	ChainId      int64  `yaml:"chainId"`
	Rpc          string `json:"rpc"`
	ContractAddr string `yaml:"contractAddr"`
	PrivateKey   string `yaml:"privateKey"`
	GasPrice     string `yaml:"gasPrice"` // wei
	GasLimit     string `yaml:"gasLimit"`
	MinProfit    uint64 `yaml:"minProfit"`
	Concurrency  uint64 `yaml:"concurrency"`
}

func LoadBotConfig(cfgFile string) (*BotConfig, error) {
	if cfgFile == "" {
		return nil, errors.New("config file path is empty")
	}
	cfg := &BotConfig{}
	viperObj := viper.New()
	viperObj.SetConfigFile(cfgFile)
	err := viperObj.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if err = viperObj.Unmarshal(cfg); err != nil {
		return nil, err
	}

	GBotConfig = cfg

	return cfg, nil
}
