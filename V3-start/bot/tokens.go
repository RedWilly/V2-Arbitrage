package main

var PolygonAmmFactory = map[string]string{
	"QuickSwap": "0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32",
	"MMFinance": "0x7cFB780010e9C861e03bCbC7AC12E013137D47A5",
	"ApeSwap":   "0xcf083be4164828f00cae704ec15a36d711491284 ",
}

type Tokens struct {
	Symbol  string
	Address string
}

var PolygonBaseTokens = []Tokens{
	{"WMATIC", "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"},
	{"USDC", "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"},
	{"USDT", "0xc2132D05D31c914a87C6611C10748AEb04B58e8F"},
}

var PolygonQuoteTokens = []Tokens{
	{"CORE", "0x2B03E8897840Dc2d2db6688F6E78A7Eae2e2f6A5"},
}

// ArbitragePair  需要监控的两个池子
type ArbitragePair struct {
	Symbols string
	Pairs   []string // pair0. pair1
}
