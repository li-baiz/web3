// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract EtherStore {

    uint256 public withdrawalLimit = 1 ether;
    mapping(address => uint256) public lastWithdrawTime;
    mapping(address => uint256) public balances;

    function depositFunds() public payable {
        balances[msg.sender] += msg.value;
    }

    function withdrawFunds(uint256 _weiToWithdraw) public {
        require(balances[msg.sender] >= _weiToWithdraw, "Insufficient balance");
        // 限制提现金额
        require(_weiToWithdraw <= withdrawalLimit, "Exceeds withdrawal limit");
        // 限制允许提现的时间
        require(block.timestamp >= lastWithdrawTime[msg.sender] + 1 weeks, "Withdrawal not allowed yet");

        // 使用新的 call 语法，并传递空的 calldata
       (bool success, ) = msg.sender.call{value: address(this).balance}("");
require(success, "Transfer failed");


        balances[msg.sender] -= _weiToWithdraw;
        lastWithdrawTime[msg.sender] = block.timestamp;
    }
}
