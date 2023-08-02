const fs = require('fs');
const Web3 = require('web3');

const nodeUrl = 'https://eth.public-rpc.com';

const uniswapV3FactoryAddress = '0x1F98431c8aD98523631AE4a59f267346ea31F984';

// --number of blocks to process in each batch
const blockBatchSize = 1000;

// deployment block of the factory contract to stop
const deploymentBlock = 12369621;

async function getAllPoolAddresses() {
    try {
        const web3 = new Web3(nodeUrl);
        const factoryContract = new web3.eth.Contract(UNISWAP_V3_FACTORY_ABI, uniswapV3FactoryAddress);

        // --get the latest block number
        const latestBlockNumber = await web3.eth.getBlockNumber();

        let startBlock = latestBlockNumber;

        const poolAddresses = [];
        while (startBlock > deploymentBlock) {
            const endBlock = Math.max(deploymentBlock, startBlock - blockBatchSize + 1);
            console.log(`Processing blocks from ${startBlock} to ${endBlock}`);

            // --let get past PoolCreated events within the specified block range
            const events = await factoryContract.getPastEvents('PoolCreated', {
                fromBlock: endBlock,
                toBlock: startBlock,
            });

            events.forEach((event) => {
                const { pool } = event.returnValues;
                poolAddresses.push(pool);
                console.log(`Pool Address: ${pool}`);

                const jsonContent = JSON.stringify(poolAddresses, null, 2);
                fs.writeFileSync('V3pairlist.json', jsonContent);
            });

            startBlock = endBlock - 1;
        }

        console.log('Pool addresses have been saved in V3pairlist.json');
    } catch (error) {
        console.error('error occurred:', error);
    }
}

const UNISWAP_V3_FACTORY_ABI = require('./zfactory.json');

getAllPoolAddresses();
