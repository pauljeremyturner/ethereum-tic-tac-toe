# ethereum-tic-tac-toe
play the computer at tic tac toe via ethereum blockchain (geth/ golang)

## NOT MUCH TO SEE HERE, CHECK BACK LATER

## ReST api


## Building and Deploying the Smart Contract

### Compile Smart Contract

First you need to install a solc compiler.  Check the internet for options on how to do that depending on your operating system.

```
cd ./contracts
solc --abi Tictactoe.sol -o build
```

Check the directory <project root>/contracts/build for a `Tictactoe.abi` file
 
### Generate Go binding code
  
  There are other ways to do this, but generating the code to interact woth the Smart Contract via the gth api is extremely fiddly.  We can use `abigen` a geth utility to generate this code for us.
  
  Checkout github.com/ethereum/go-ethereum
  
  ```
  cd <project root>
  go run ./cmd/abigen/main.go 
  abigen	--abi=./build/Tictactoe.abi	--pkg=store	--out=Tictactoe-geth.go
  ```
  
### Compile the abi into EVM bytecode
   
   ```
   solc	--bin	./Tictactoe.sol	-o	build
   ```


## Tasks

 - ~~ETTT-1 Web service endpoints.  Define a restful api that will enable players to start and play a game.  The enpoints will only log to console at this stage~~ DONE
 
 Definition of Done: Use Insomnia to test this api.  Endpoints for start game and play turn.
 
 - ETTT-2 Smart contract. Define a smart contract that will store the state of the game and enable players to interact with it.
 
 Definition of Done: Smart contract with unit tests deployable and runnable via geth api.
 
 - ETTT-3 Plumb together web services and Smart Contract using geth
 
 Definition of Done: Web service endpoints trigger Smart Contract, check testnet for Contract state.
 
 
