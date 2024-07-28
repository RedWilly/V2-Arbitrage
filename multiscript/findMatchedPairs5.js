const { bnSqrt, saveJSONToFile, saveObjToCsv, readFromJSONFile } = require("./helpers");
const hre = require("hardhat");
const { PIVOT_TOKEN, PAIRLIST_OUTPUT_FILE, MATCHED_PAIRS_OUTPUT_FILE } = require("./config");

// helpers
function getPairsOtherToken(pair, firstToken) {
  if (pair["token0Address"] == firstToken) return pair["token1Address"];
  else if (pair["token1Address"] == firstToken) return pair["token0Address"];
  else return "";
}

function isTokenInPair(pair, token) {
  return pair["token0Address"] == token || pair["token1Address"] == token;
}

function printProgress(curr, total) {
  process.stdout.clearLine();
  process.stdout.cursorTo(0);
  process.stdout.write(`${curr.toString()} of total ${total.toString()}`);
}

const pairsArr = readFromJSONFile(PAIRLIST_OUTPUT_FILE);
const pivotPairs = pairsArr.filter((pair) => isTokenInPair(pair, PIVOT_TOKEN));
const pivotPairTokens = new Set(pivotPairs.map((p) => getPairsOtherToken(p, PIVOT_TOKEN)));

const otherPairs = pairsArr.filter((pair) => !isTokenInPair(pair, PIVOT_TOKEN));
console.log(`Total number of pivot pairs: ${pivotPairs.length}`);
console.log(`Total number of other pairs: ${otherPairs.length}`);

const matchPairs = [];
const includedPaths = new Set();
const lenPivotPairs = pivotPairs.length;

for (let [index, pivotPair] of pivotPairs.entries()) {
  printProgress(index + 1, lenPivotPairs);
  let firstToken = getPairsOtherToken(pivotPair, PIVOT_TOKEN);
  let matchObj = {};
  matchObj[firstToken] = [];
  matchObj.startPool = pivotPair["pairAddress"];
  matchObj.middlePools = [];
  matchObj.middlePools2 = [];
  matchObj.middlePools3 = [];
  matchObj.middlePools4 = [];
  matchObj.endPools = [];
  for (let secondPair of otherPairs) {
    let secondToken = getPairsOtherToken(secondPair, firstToken);
    if (secondToken != "") {
      for (let thirdPair of otherPairs) {
        let thirdToken = getPairsOtherToken(thirdPair, secondToken);
        if (thirdToken != "" && thirdToken != firstToken) {
          for (let fourthPair of otherPairs) {
            let fourthToken = getPairsOtherToken(fourthPair, thirdToken);
            if (fourthToken != "" && fourthToken != secondToken && fourthToken != firstToken) {
              for (let endPair of pivotPairs) {
                let otherToken = getPairsOtherToken(endPair, PIVOT_TOKEN);
                if (otherToken == fourthToken) {
                  const pathPairs = [
                    pivotPair["pairAddress"].toLowerCase(),
                    secondPair["pairAddress"].toLowerCase(),
                    thirdPair["pairAddress"].toLowerCase(),
                    fourthPair["pairAddress"].toLowerCase(),
                    endPair["pairAddress"].toLowerCase(),
                  ];
                  const pathPairsJoined = pathPairs.sort().join();
                  if (!includedPaths.has(pathPairsJoined)) {
                    matchObj[firstToken].push([secondToken, thirdToken, fourthToken]);
                    matchObj.middlePools.push(secondPair["pairAddress"]);
                    matchObj.middlePools2.push(thirdPair["pairAddress"]);
                    matchObj.middlePools3.push(fourthPair["pairAddress"]);
                    matchObj.endPools.push(endPair["pairAddress"]);
                    includedPaths.add(pathPairsJoined);
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  if (matchObj[firstToken].length > 0) {
    matchPairs.push(matchObj);
  }
}

saveJSONToFile(MATCHED_PAIRS_OUTPUT_FILE, matchPairs);
console.log("\nPentagonal arbitrage paths saved to", MATCHED_PAIRS_OUTPUT_FILE);