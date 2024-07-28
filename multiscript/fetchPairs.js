//fetchpairs.js
const hre = require("hardhat");
const { saveJSONToFile, saveObjToCsv, readFromJSONFile } = require("./helpers");
const { MULTI_ARBIT_ADDRESS, FACTORY_ADDRESSES, PAIRLIST_OUTPUT_FILE } = require("./config");

const QUERY_STEP = 30; // Adjust the value as needed
const MAX_DAYS_OLD = 300; // The pool must be active latest 7 days ago in order to be added into results

const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

const retryQuery = async (fn, maxAttempts = 5, baseDelay = 1000) => {
    let attempt = 0;
    while (attempt < maxAttempts) {
        try {
            return await fn();
        } catch (error) {
            console.log(`Attempt ${attempt + 1} failed. Retrying...`);
            attempt++;
            await sleep(baseDelay * Math.pow(2, attempt));
        }
    }
    throw new Error(`Max attempts (${maxAttempts}) reached. Query failed.`);
};

const fetchData = async () => {
    const superArbit = await hre.ethers.getContractAt("SuperArbit", MULTI_ARBIT_ADDRESS);
    console.log(`Contract deployed at ${superArbit.address}`);
    const pairsArr = [];
    const timeLimit = Math.floor(Date.now() / 1000) - MAX_DAYS_OLD * 24 * 60 * 60;
    try {
        for (const [key, factoryAddr] of Object.entries(FACTORY_ADDRESSES)) {
            const swapFactory = await hre.ethers.getContractAt("IPancakeFactory", factoryAddr);
            const totalNumOfPairs = await swapFactory.allPairsLength();
            const loopLim = Math.floor(totalNumOfPairs.toNumber() / QUERY_STEP) + 1;
            console.log(`Factory: ${key}`);
            console.log(`Total Number of Pairs: ${totalNumOfPairs}`);
            console.log(`Loop limit: ${loopLim}\n`);
            for (let i = 0; i < loopLim; i++) {
                try {
                    console.log(`Querying pairs from index ${i * QUERY_STEP} to ${(i + 1) * QUERY_STEP}...`);
                    const data = await retryQuery(() =>
                        superArbit.retrievePairInfo(factoryAddr, i * QUERY_STEP, QUERY_STEP)
                    );
                    data.forEach((e) => {
                        if (e.lastBlockTimestamp >= timeLimit) {
                            pairsArr.push({
                                fromFactory: key,
                                pairAddress: e.pairAddr,
                                token0Address: e.token0Addr,
                                token1Address: e.token1Addr,
                                token0Symbol: e.token0Symbol,
                                token1Symbol: e.token1Symbol,
                                lastActivity: e.lastBlockTimestamp,
                                poolId: e.poolId,
                            });
                        }
                    });
                } catch (error) {
                    console.log("Timeout.Trying again...");
                }
            }
        }
    } catch (error) {
        console.log("Call Exception Error, aborting...");
        console.log(error);
    } finally {
        // write JSON string to a file
        saveJSONToFile(PAIRLIST_OUTPUT_FILE, pairsArr);
        console.log("JSON file is created.");
    }
};

fetchData()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
