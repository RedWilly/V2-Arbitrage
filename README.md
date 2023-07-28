# Triangular DEX Arbitrage

## 1. Description

This bot is designed to make automated crypto token cyclic-arbitrage transactions on DEX platforms(Decentralized Exchange) for profit. The implemented form of the arbitrage is cyclic-arbitrage with three legs, so the name "Triangular". In triangular arbitrage, the aim is to start with an asset(here it is crypto tokens) and do 3 swap transactions to end with the start token. An example would be: WBNB->BUSD->Cake->WBNB The bot constantly searches for arbitrage opportunites on different DEX platforms and if the trade is profitable(end amount > start amount + transaction fees), then the trade is executed. It can be used on DEX platforms where Uniswap V2 AMM(Automated Market Maker) is utilized for swap calculations.

I used/tested this bot on Dogechain, CoreDao Gnosis & especially Binance Smart Chain, where significant number of DEX platforms can be found, which are all Uniswap V2 clones (Such as PancakeSwap, biSwap, MDEX etc.) yet it can be used on other EVM-compatible blockchains like Ethereum, Avalanche etc. with slight modification of parameters.

The algorithm(profitibility calculations, calculation of optimum input token amount etc.) used in this project is taken from [this paper](https://arxiv.org/pdf/2105.02784.pdf)

A smart contract is also written for batched data fetching from blockchain(to speed up the searching as well as for minimizing the number of request from RPC Node API) and batched static checking of swap transactions to see if they go through without executing actual transaction. In order to run the bot, the contract must be deployed on mainnet of the blockchain(for example on BSC/DOOGECHAIN) - CREDIT FOR THE CONTRACT DEVELOPMENT GOES TO UFUK

## 2. Run on local or Server

### 2.1 Requirements

After cloning this repo: (node.js must be already installed)

```bash
npm install
```

### 2.2. Usage

Before starting, change the name of the file ".env.example" to ".env" and update the information, which is then needed in hardhat.config.js

1. After installing the dependencies, first compile and deploy the contract

```bash
$ npx hardhat compile
$ npx hardhat run ./scripts/superArbitDeploy.js --network doge
```

2. After contract deployment, update the contract address in config.js(SUPER_ARBIT_ADDRESS)
   First we fetch all the swap pairs from available DEX platform pools.Only pairs from pools, that are active in last 7 days(you can change that if need), are fetched.
   This scripts outputs all the available pairs in a json file.(once it complete you can see the result in the pairsList.json)

```bash
npx hardhat run ./scripts/fetchPairs.js --network doge
```

3. Next step is to find all possible routes which starts with pivot token(WBNB/WDOGE) and ends also with the same token with two other tokens in between.
   After succesful run, this script outputs the result also in a json file.(see the result in the matchedPairs.json)

```bash
$ npx hardhat run ./scripts/findMatchedPairs.js --network doge
```

4. In the last step, run the main.js to check arbitrage opportunities as well as execute transactions if they are profitable.

```bash
$ npx hardhat run ./scripts/main.js --network doge
```

DONT MIND ME - HERE

5. But before you run the main.js i will advice you update arbUtils.js line 57 'APPROX_GAS_FEE'
   if input is 10 doge and profit found 0.6 doge. it will subtract the APPROX_GAS_FEE from the profit. i have currently set it to 0.32 doge
   calculate (0.6 - 0.32 = 0.28) so profit made is 0.28 doge

NOTE:the 0.32 i have sent is more than the require gas fee. the highest gas price the contract is 0.25 and the lowest is 0.04
the gas price goes up only when it trading accross 3 dexes and more than 5 tokens.

so for BSC change this to: 1250000000000000(0.00125) or 1300000000000000(0.0013)

### 2.3 Optional Suggestion

1. if you wish to run this local on you backgrond while the terminal is close i will advice you use pm2.
   Install Pm2:

```bash
npm install -g pm2
```

--name DOGE: Sets the name of the process to "DOGE" or any other name for easy identification. once it start to run you can now close the terminal you are using..

```bash
pm2 start --name DOGE npx -- hardhat run ./scripts/main.js --network doge
```

TIPS:

```bash
pm2 status
```

This command will display a table with information about all the running processes managed by pm2, including the process ID (ID), name, mode, restart counter (↺), status, CPU usage (cpu), and memory usage (memory).

if you are on a server and will like to set up pm2 to start automatically on system boot

```bash
pm2 startup
```

To stop a proess already running you can you

```bash
pm2 stop DOGE
```

To start a proess already running you can you

```bash
pm2 start DOGE
```

replace the DOGE with the actual name of the process you will like to use

## Disclaimer

Use this bot at your own risk!
This bot occasionally finds arbitrage opportunities and execute them. Sometimes it is possible that the transactions are fails/reverted, which can result from many reasons.(For example, a swap transaction is executed before ours, which changes the balances one of the swap pools, so the calculation is not valid anymore). So the bot must be improved in order to catch such situations.

Big Credits to Ufuk

```

```
