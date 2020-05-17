pragma solidity ^0.4.17;

contract Tictactoe {

    address internal ownerPlayerAddress
    address internal otherPlayerAddress

    bool internal playing = true;


    uint public gameIndex = 0;
    mapping(uint => Game) private gameMap;

    function Tictactoe() public {
        ownerPlayerAddress = msg.sender;
    }

    function startGame() {

    }

    function joinGame() {
        
    }

    function move(address player) returns (uint256, uint256) {

        require(msg.sender == ownerPlayerAddress || msg.sender == otherPlayerAddress, "Only 2 players allowed");
        require(playing, "This game has finished");
    }

}