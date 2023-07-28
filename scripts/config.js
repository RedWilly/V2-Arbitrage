module.exports = {
  SUPER_ARBIT_ADDRESS: "0xc059F30196596b1e72B2299A909D8d232CAaE797", //arbitrage contract(SuperArbit.sol)
  FACTORY_ADDRESSES: {
    // Factory contract addresses of chosen DEX'es COREDAO
    icecreamswap: "0x9e6d21e759a7a288b80eef94e4737d313d31c13f",
    Fraxswap: "0x67b7DA7c0564c6aC080f0A6D9fB4675e52E6bF1d",
    quickswap: '0xc7c86B4f940Ff1C13c736b697e3FbA5a6Bc979F9',
    Dogeshrek: "0x7C10a3b7EcD42dd7D79C0b9d58dDB812f92B574A",
    Yoda: "0xAaA04462e35f3e40D798331657cA015169e005d7",
    Kibble: "0xF4bc79D32A7dEfd87c8A9C100FD83206bbF19Af5",

  },
  PIVOT_TOKEN: "0xB7ddC6414bf4F5515b52D8BdD69973Ae205ff101", // Wrapped Token
  PAIRLIST_OUTPUT_FILE: "./pairsList.json",
  MATCHED_PAIRS_OUTPUT_FILE: "./matchedPairs.json",
  MAX_GAS: 2000000,
  BSC_GAS_PRICE: 5000000000,
  MAX_TRADE_INPUT: 101 // WBNB/WDOGE/WCORE ( how much wrapped token you are willing to use and trade with)
};