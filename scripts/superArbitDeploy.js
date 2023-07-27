const { ethers } = require("hardhat");
const hre = require("hardhat");
const { PIVOT_TOKEN } = require("./config");

async function deploy() {
  // Get the contract to deploy
  const SuperArbit = await hre.ethers.getContractFactory("SuperArbit");
  const superArbit = await SuperArbit.deploy();

  await superArbit.deployed();
  console.log("SuperArbit deployed to:", superArbit.address);

  // Approve super arbit contract to swap Wrapped Token
  const wToken = await ethers.getContractAt("IERC20", PIVOT_TOKEN);
  await wToken.approve(superArbit.address, "0x1fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"); //infinite approval
}

deploy()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
