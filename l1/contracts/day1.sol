// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.20;

contract NumberStorage {
    uint public x;
    function setX(uint px) public {
        x = px;
    }
    function getX() public view returns(uint) {
        return x;
    }
    function add(uint a,uint b) private pure returns(uint){
        return a + b;
    }
}
