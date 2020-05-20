pragma solidity ^0.5.10;
contract Test {
event ItemSet(bytes32 key, bytes32 value);
string public version;
mapping (bytes32 => bytes32) public items;


function setItem(bytes32 key, bytes32 value) external {
items[key] = value;
emit ItemSet(key, value);
}
}