const hre = require("hardhat");
const totalSupply = 10000000;

async function main() {
  const factory = await hre.ethers.getContractFactory("VenvetCoin");
  const venvetToken = await factory.deploy(totalSupply);

  await venvetToken.deployed();

  console.log("VenvetCoin deployed to:", venvetToken.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
