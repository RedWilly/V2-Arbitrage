const { generatePaths, addPairReserves, calculateProfit, APPROX_GAS_FEE } = require("./arbUtils");
const hre = require("hardhat");
const { ethers, network } = require("hardhat");
const {
  PIVOT_TOKEN,
  MULTI_ARBIT_ADDRESS,
  MATCHED_PAIRS_OUTPUT_FILE,
  MAX_GAS,
  BSC_GAS_PRICE,
  MAX_TRADE_INPUT,
} = require("./config");

let execCount = 0;

const checkProfitAndExecute = async function (lucrPaths, router, signer) {
  console.log("Static batch check starts...");
  const startToken = PIVOT_TOKEN;
  const validPaths = [];

  for (const lucrPath of lucrPaths) {
    const pools = lucrPath.pools;
    amounts = [lucrPath.optimumAmountInBN];
    for (let i = 0; i < lucrPath.path.length - 1; i++) {
      if (lucrPath.path[i].toLowerCase() < lucrPath.path[i + 1].toLowerCase()) {
        amounts.push("0");
        amounts.push(lucrPath.swapAmounts[i]);
      } else {
        amounts.push(lucrPath.swapAmounts[i]);
        amounts.push("0");
      }
    }
    lucrPath.execAmounts = amounts;
    lucrPath.execPools = pools;

    // Validate amounts and pools
    const isValid = amounts.every(amount => ethers.BigNumber.from(amount).gte(0)) &&
                    pools.every(pool => ethers.utils.isAddress(pool));
    
    if (isValid) {
      validPaths.push(lucrPath);
      console.log("Valid path found:", JSON.stringify({
        path: lucrPath.path,
        pools: lucrPath.pools,
        amounts: lucrPath.execAmounts,
        expectedProfit: lucrPath.expectedProfit
      }, null, 2));
    } else {
      console.log("Invalid path detected:", JSON.stringify(lucrPath, null, 2));
    }
  }

  console.log(`Valid paths: ${validPaths.length} out of ${lucrPaths.length}`);

  if (validPaths.length === 0) {
    console.log("No valid paths to process");
    return [];
  }

  const amountsArr = validPaths.map((l) => l.execAmounts);
  const poolsArr = validPaths.map((l) => l.execPools);
  const pathLengths = validPaths.map((l) => l.path.length - 1);
  
  let result = [];
  try {
    result = await router.callStatic.multiSwapBatch(
      amountsArr, 
      poolsArr, 
      Array(validPaths.length).fill(startToken), 
      pathLengths, 
      { gasLimit: MAX_GAS * 10 }
    );
    console.log("Static multiSwapBatch check passed");
  } catch (error) {
    console.log(`MultiSwapBatch error: ${error.reason}`);
    console.log("Error details:", error);
    return [];
  }

  const lucrPathsPassed = validPaths.filter((l, index) => result[index]);
  console.log("Number of paths which passed static check: ", lucrPathsPassed.length);

  for (const path of lucrPathsPassed) {
    path.gas = "0";
    if (parseFloat(path.optimumAmountIn) < MAX_TRADE_INPUT) {
      console.log("Checking path:", JSON.stringify({
        path: path.path,
        pools: path.pools,
        // amounts: path.execAmounts,
        // expectedProfit: path.expectedProfit
      }, null, 2));
      try {
        console.log("Estimating gas...");
        let gas = await router.estimateGas.multiSwap(
          path.execAmounts, 
          path.execPools, 
          startToken, 
          path.path.length - 1
        );
        console.log("Gas(static) used: ", gas.toString());
        path.gas = gas.toString();
        const newProfit = ethers.BigNumber.from(path.expectedProfitBN)
          .sub(gas.mul(BSC_GAS_PRICE.toString()))
          .add(APPROX_GAS_FEE);
        console.log("New Profit", parseFloat(ethers.utils.formatEther(newProfit)));
        if (newProfit.gt(0)) {
          console.log("Profit is positive, attempting execution...");
          try {
            await router.callStatic.multiSwap(
              path.execAmounts, 
              path.execPools, 
              startToken, 
              path.path.length - 1, 
              { gasLimit: MAX_GAS }
            );
            console.log("Static multiSwap check passed, executing transaction...");
            await router.multiSwap(
              path.execAmounts, 
              path.execPools, 
              startToken, 
              path.path.length - 1, 
              { gasLimit: MAX_GAS }
            );
            console.log("!!!!EXECUTED!!!");
            execCount++;
          } catch (error) {
            console.log("Error during multiSwap execution:", error.reason);
          }
        } else {
          console.log("Profit after gas costs is not positive, skipping execution");
        }
      } catch (error) {
        console.log("Error during gas estimation:", error.reason);
      }
    } else {
      console.log("Optimum amount in exceeds MAX_TRADE_INPUT, skipping");
    }
  }
  return lucrPathsPassed;
}

const main = async () => {
  // Connect to router and other stuff
  const router = await ethers.getContractAt("MultiArbit", MULTI_ARBIT_ADDRESS);
  const signer = await ethers.getSigner();
  let paths = generatePaths(MATCHED_PAIRS_OUTPUT_FILE); // generate paths with pivot token -> WBNB
  let allLucrPathsPassed = [];
  
  while (true) {
    const stepSize = 50;
    const numOfPaths = paths.length;
    const loopLim = Math.floor(numOfPaths / stepSize);
    console.log(`\nNumber of Paths from JSON:${numOfPaths}, Total number of batches:${loopLim}\n`);
    let i = 0;
    let pathsSliced;

    while (i <= loopLim) {
      console.log(`Processing batch ${i + 1} of total ${loopLim}`);
      if (i != loopLim) {
        pathsSliced = paths.slice(i * stepSize, (i + 1) * stepSize);
      } else {
        pathsSliced = paths.slice(i * stepSize, i * stepSize + (numOfPaths % stepSize));
      }
      const pathsWithRes = await addPairReserves(pathsSliced, router, (batchSize = stepSize * 6)); // Increased batchSize to accommodate longer paths
      const lucrPaths = calculateProfit(pathsWithRes);
      console.log("Length of lucrative paths in current batch:", lucrPaths.length);
      
      // Check and execute profitable paths
      const lucrPathsPassed = await checkProfitAndExecute(lucrPaths, router, signer);
      if (lucrPathsPassed.length > 0) allLucrPathsPassed = allLucrPathsPassed.concat(lucrPathsPassed);
      console.log("Length all lucrative paths passed: ", allLucrPathsPassed.length);
      console.log(`-------Total number of executions: ${execCount}\n`);
      i++;
    }
  }
};

main().then();