require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-etherscan");
require("dotenv").config();

/**
 * @type import('hardhat/config').HardhatUserConfig
 */

module.exports = {
  defaultNetwork: "hardhat",
  networks: {
    hardhat: { //archive node
      forking: {
        url: `https://speedy-nodes-nyc.moralis.io/${process.env.MORALIS_SPEEDY_API_KEY}/bsc/mainnet/archive`,
        blockNumber: 17360588,
      },
    },
    doge: {
      url: `https://rpc.ankr.com/dogechain`,
      accounts: [process.env.PRIV_KEY],
    },
    bscmain_quicknode: {
      url: `https://weathered-damp-forest.bsc.quiknode.pro/${process.env.QUICKNODE_KEY}/`,
      accounts: [process.env.PRIV_KEY],
    },
  },
  /*etherscan: {
    apiKey: process.env.ETHERSCAN_TOKEN,
  },*/
  solidity: "0.8.4",
};
