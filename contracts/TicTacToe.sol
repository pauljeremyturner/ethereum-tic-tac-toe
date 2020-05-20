pragma solidity ^0.5.10;

contract Tictactoe {

  address internal owner;
  uint256 playPriceWei;

  struct Game {
    string circlePlayer;
    address circleAddress;
    string crossPlayer;
    address crossAddress;
    uint256 circles;
    uint256 crosses;
    bool circleTurn;
    bool inProgress;
  }

  mapping(address => Game) games;

    constructor(uint256 ppw) public {
        owner = msg.sender;
        playPriceWei = ppw;
    }

    function startGame(string memory name, bool circle, uint256 firstMove) public payable {
        require(msg.value == playPriceWei, "Player must transfer exactly the play Price in Wei to start a game");
        Tictactoe.Game storage game = games[msg.sender];
        game.inProgress = true;
        if (circle) {
            game.circlePlayer = name;
            game.circles = firstMove;
            game.circleTurn = false;
            game.circleAddress = msg.sender;
        } else {
            game.crossAddress = msg.sender;
            game.crossPlayer = name;
            game.crosses = firstMove;
            game.circleTurn = true;
        }
    }

}