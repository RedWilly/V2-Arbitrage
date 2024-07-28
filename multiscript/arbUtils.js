const { bnSqrt, saveJSONToFile, saveObjToCsv, readFromJSONFile } = require("./helpers");
const { ethers, network } = require("hardhat");
const hre = require("hardhat");
const { recoverAddress } = require("ethers/lib/utils");
const { PIVOT_TOKEN } = require("./config");

// // Modified to support multiple path lengths
const generatePaths = function (matchPairFile) {
  const matchPairList = readFromJSONFile(matchPairFile);
  const tokenPaths = [];
  for (let match of matchPairList) {
    let firstToken = Object.keys(match)[0];
    for (let i in match[firstToken]) {
      let pathObj = {};
      if (Array.isArray(match[firstToken][i])) {
        if (match[firstToken][i].length === 2) {
          // Quadrilateral path
          pathObj.path = [PIVOT_TOKEN, firstToken, ...match[firstToken][i], PIVOT_TOKEN];
          pathObj.pools = [match.startPool, match.middlePools[i], match.middlePools2[i], match.endPools[i]];
        } else if (match[firstToken][i].length === 3) {
          // Pentagonal path
          pathObj.path = [PIVOT_TOKEN, firstToken, ...match[firstToken][i], PIVOT_TOKEN];
          pathObj.pools = [match.startPool, match.middlePools[i], match.middlePools2[i], match.middlePools3[i], match.endPools[i]];
        } else if (match[firstToken][i].length === 4) {
          // Hexagonal path
          pathObj.path = [PIVOT_TOKEN, firstToken, ...match[firstToken][i], PIVOT_TOKEN];
          pathObj.pools = [match.startPool, match.middlePools[i], match.middlePools2[i], match.middlePools3[i], match.middlePools4[i], match.endPools[i]];
        }
      } else {
        // Triangular path
        pathObj.path = [PIVOT_TOKEN, firstToken, match[firstToken][i], PIVOT_TOKEN];
        pathObj.pools = [match.startPool, match.middlePools[i], match.endPools[i]];
      }
      tokenPaths.push(pathObj);
    }
  }
  return tokenPaths;
};

// Modified to work with variable path lengths
const addPairReserves = async function (paths, multiArbit, batchSize) {
  const pairAddrList = paths.map((e) => e.pools).flat();
  let pairReserveList = [];
  const numOfPairs = pairAddrList.length;
  const loopLim = Math.floor(numOfPairs / batchSize);
  let i = 0;
  let pairAddrBatch;
  while (i <= loopLim) {
    if (i != loopLim) {
      pairAddrBatch = pairAddrList.slice(i * batchSize, (i + 1) * batchSize);
    } else {
      pairAddrBatch = pairAddrList.slice(i * batchSize, i * batchSize + (numOfPairs % batchSize));
    }
    try {
      pairReserveList = pairReserveList.concat(await multiArbit.getBatchReserves(pairAddrBatch));
      i++;
    } catch (error) {
      console.log("Trying again step ", i);
    }
  }
  const pathsWithRes = [];
  let reserveIndex = 0;
  for (const path of paths) {
    let numLegs = path.pools.length;
    path.reserves = pairReserveList.slice(reserveIndex, reserveIndex + numLegs);
    pathsWithRes.push(path);
    reserveIndex += numLegs;
  }
  return pathsWithRes;
};

// params
const RATIO_SCALE_FACT = 100000;
const RATIO_SCALE_FACT_BN = ethers.BigNumber.from(RATIO_SCALE_FACT);
const MAX_RATIO_LIM = ethers.BigNumber.from(1.5 * RATIO_SCALE_FACT);
const R1 = ethers.BigNumber.from(0.9969 * RATIO_SCALE_FACT); // %0,3 input fee
const R2 = ethers.BigNumber.from(1 * RATIO_SCALE_FACT); // %0 output fee
const APPROX_GAS_FEE = ethers.BigNumber.from("46000000000000"); //("1250000000000000"); //250000 gas * 5 gwei per gas

// Modified to work with variable path lengths
const getAmountsOut = function (amountIn, reserves, r1, ratioScaleFact) {
  const amounts = [];
  let amountInTemp = amountIn;
  for (const reserve of reserves) {
    amountInTemp = getAmountOut(amountInTemp, reserve[0], reserve[1], r1, ratioScaleFact);
    amounts.push(amountInTemp);
  }
  return amounts;
};

const getAmountOut = function (amountIn, res0, res1, r1, ratioScaleFact) {
  return amountIn
    .mul(r1)
    .mul(res1)
    .div(res0.mul(ratioScaleFact).add(amountIn.mul(r1)));
};

// Modified to support multiple path lengths
const calculateProfit = function (pathsWithRes) {
  lucrPaths = [];
  for (const path of pathsWithRes) {
    let reserves = [];
    for (const [i, pool] of path.pools.entries()) {
      const resPool = path.reserves[i];
      if (path.path[i].toLowerCase() < path.path[i + 1].toLowerCase()) reserves.push([resPool.reserve0, resPool.reserve1]);
      else reserves.push([resPool.reserve1, resPool.reserve0]);
    }
    let res = calcRatio(reserves, R1, R2);
    if (res.ratio != undefined) {
      if (res.reverse) {
        path.path.reverse();
        path.pools.reverse();
        reserves.reverse();
        reserves.map((r) => r.reverse());
      }
      const { optAmountIn, amountOut } = calcOptiAmountIn(reserves, R1, R2);
      
      // Skip paths where calculations resulted in zero values
      if (optAmountIn.isZero() || amountOut.isZero()) {
        continue;
      }
      
      const expectedProfit = amountOut.sub(optAmountIn).sub(APPROX_GAS_FEE);
      
      // Only consider paths with positive profit
      if (expectedProfit.gt(0) && optAmountIn.gt(0)) {
        path.reserves = reserves.map((rs) => rs.map((r) => r.toString()));
        swapAmounts = getAmountsOut(optAmountIn, reserves, R1, RATIO_SCALE_FACT_BN);
        path.swapAmounts = swapAmounts.map((s) => s.toString());
        path.ratio = res.ratio.toNumber() / RATIO_SCALE_FACT;
        path.optimumAmountInBN = optAmountIn.toString();
        path.AmountOutBN = amountOut.toString();
        path.expectedProfitBN = expectedProfit.toString();
        path.realRatio = amountOut.sub(optAmountIn).mul(10000).div(optAmountIn).toNumber() / 10000;
        path.optimumAmountIn = parseFloat(ethers.utils.formatEther(optAmountIn));
        path.AmountOut = parseFloat(ethers.utils.formatEther(amountOut));
        path.expectedProfit = parseFloat(ethers.utils.formatEther(expectedProfit));
        lucrPaths.push(path);
      }
    }
  }
  return lucrPaths;
};

// Modified to support multiple path lengths
function calcOptiAmountIn(reserves, r1, r2) {
  const numReserves = reserves.length;
  let a = reserves[numReserves - 1][0];
  let a_ = reserves[numReserves - 1][1].mul(r1).mul(r2).div(RATIO_SCALE_FACT_BN.pow(2));

  // Check if initial values are valid
  if (a.isZero() || a_.isZero()) {
    return { optAmountIn: ethers.constants.Zero, amountOut: ethers.constants.Zero };
  }

  for (let i = numReserves - 2; i >= 0; i--) {
    const denominator = a.add(r1.mul(r2).mul(a_).div(RATIO_SCALE_FACT_BN.pow(2)));
    if (denominator.isZero()) {
      return { optAmountIn: ethers.constants.Zero, amountOut: ethers.constants.Zero };
    }
    const a_next = reserves[i][0].mul(a).div(denominator);
    const a_next_ = reserves[i][1].mul(a_).mul(r1).mul(r2).div(RATIO_SCALE_FACT_BN.pow(2).mul(denominator));
    a = a_next;
    a_ = a_next_;

    // Check if values are still valid
    if (a.isZero() || a_.isZero()) {
      return { optAmountIn: ethers.constants.Zero, amountOut: ethers.constants.Zero };
    }
  }

  const sqrtTerm = bnSqrt(r1.mul(r2).mul(a_).mul(a).div(RATIO_SCALE_FACT_BN.pow(2)));
  if (sqrtTerm.lte(a)) {
    return { optAmountIn: ethers.constants.Zero, amountOut: ethers.constants.Zero };
  }

  const optAmountIn = sqrtTerm.sub(a).mul(RATIO_SCALE_FACT_BN).div(r1);

  // calculate achievable amountOut
  let amountOut;
  let amountIn = optAmountIn;
  for (const r of reserves) {
    const denominator = r[0].add(r1.mul(amountIn).div(RATIO_SCALE_FACT_BN));
    if (denominator.isZero()) {
      return { optAmountIn: ethers.constants.Zero, amountOut: ethers.constants.Zero };
    }
    amountOut = r1.mul(r2).mul(r[1]).mul(amountIn).div(RATIO_SCALE_FACT_BN.pow(2)).div(denominator);
    amountIn = amountOut;
  }
  return { optAmountIn, amountOut };
}

// Modified to support multiple path lengths
function calcRatio(reserves, r1, r2) {
  let result = { ratio: undefined, reverse: undefined };
  try {
    const feeRatio = RATIO_SCALE_FACT_BN.pow(reserves.length * 2 + 1).div(r1.pow(reserves.length)).div(r2.pow(reserves.length));
    let num = RATIO_SCALE_FACT_BN;
    let den = RATIO_SCALE_FACT_BN;
    for (const reserve of reserves) {
      num = num.mul(reserve[1]);
      den = den.mul(reserve[0]);
    }
    const forwardRatio = num.mul(RATIO_SCALE_FACT).div(den);
    const reverseRatio = den.mul(RATIO_SCALE_FACT).div(num);
    if (forwardRatio.gt(RATIO_SCALE_FACT_BN) && forwardRatio.lt(MAX_RATIO_LIM) && forwardRatio.gt(feeRatio)) {
      result = { ratio: forwardRatio.sub(feeRatio), reverse: false };
    } else if (reverseRatio.gt(RATIO_SCALE_FACT_BN) && reverseRatio.lt(MAX_RATIO_LIM) && reverseRatio.gt(feeRatio)) {
      result = { ratio: reverseRatio.sub(feeRatio), reverse: true };
    }
  } catch (error) {}

  return result;
}

exports.generatePaths = generatePaths;
exports.addPairReserves = addPairReserves;
exports.calculateProfit = calculateProfit;
exports.APPROX_GAS_FEE = APPROX_GAS_FEE;