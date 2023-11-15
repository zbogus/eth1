// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EthTransfer {

    // Belirtilen adrese belirtilen miktar ETH gönderilmesi.
    function transferEth(address payable to, uint256 amount) external payable {
        to.transfer(amount);
    }

    // Belirtilen hesabın bakiyesini döndüren fonksiyon.
    function getBalance(address account) external view returns (uint256) {
        return account.balance;
    }
    receive() external payable{}
}