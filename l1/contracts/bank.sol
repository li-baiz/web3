// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.20;
//所有人都可以存钱
//只有合约owner才可以取钱
//只要取钱，合约就销毁selfdestruct
//扩展：支持主币之外的资产

contract Bank {
    //状态变量
    address public immutable owner;
    //事件
    event Deposit(address _ads, uint256 amount);
    event Withdraw(uint256 amount);
    //receive
    receive() external payable {
        emit Deposit(msg.sender,msg.value);
    }
    //构造函数
    constructor() payable{
        owner = msg.sender;
    }
    //方法
    function withdraw() external {
        require(msg.sender == owner,"Not owner");
        emit Withdraw(address(this).balance);
    }
    function getBalance() external view returns (uint256){
        return address(this).balance;
    }
}
