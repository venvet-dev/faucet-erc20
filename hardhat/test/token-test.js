const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("VevetCoin", function () {
  it("Check if token supply is correct", async function () {
    const [owner] = await ethers.getSigners();

    const tokenSupply = 100000;

    const factory = await ethers.getContractFactory("VenvetCoin");
    const token = await factory.deploy(tokenSupply);

    await token.deployed();

    const tokenDecimals = await token.decimals();

    // Check total supply
    const bigNumberDecimals = ethers.BigNumber.from(10).pow(ethers.BigNumber.from(tokenDecimals));
    const bigNumberSupply = ethers.BigNumber.from(tokenSupply).mul(ethers.BigNumber.from(bigNumberDecimals));
    expect(await token.totalSupply()).to.equal(bigNumberSupply);
  });

  it("Check if deployer supply is the max supply", async function () {
    const [owner] = await ethers.getSigners();

    const tokenSupply = 10000;

    const factory = await ethers.getContractFactory("VenvetCoin");
    const token = await factory.deploy(tokenSupply);

    await token.deployed();

    // Check deployer balance
    const balance = await token.balanceOf(owner.address);
    expect(await token.totalSupply()).to.equal(balance);
  });
});
