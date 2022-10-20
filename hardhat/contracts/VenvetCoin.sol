// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.2;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract VenvetCoin is ERC20 {
  constructor(uint256 _supply) ERC20("VenvetCoin", "VVC") {
    _mint(msg.sender, _supply * 10 ** decimals());
  }
}
