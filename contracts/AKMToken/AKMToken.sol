// SPDX-License-Identifier: MIT
pragma solidity 0.8.27;

import "../@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../@openzeppelin/contracts/access/Ownable.sol";

contract AKMToken is ERC20, Ownable{
    constructor(address initialOwner) Ownable(initialOwner) ERC20("AKMToken", "AKM") {
        _mint(msg.sender, 1000 * 10 ** decimals());
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }
}
