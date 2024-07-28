const { ethers } = require("hardhat");
const hre = require("hardhat");
const { PIVOT_TOKEN } = require("./config");

async function deploy() {
  // Get the contract to deploy
  const MultiArbit = await hre.ethers.getContractFactory("MultiArbit");
  const multiArbit = await MultiArbit.deploy();

  await multiArbit.deployed();
  console.log("MultiArbit deployed to:", multiArbit.address);

  // Approve multi arbit contract to swap Wrapped Token
  const wToken = await ethers.getContractAt("IERC20", PIVOT_TOKEN);
  await wToken.approve(multiArbit.address, "0x1fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"); //infinite approval
}

deploy()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
