# Minesweeper GO  
A Go implementation of Minesweeper, currently available in terminal. In the future it will also compile for wasm and you can use it in your web projects

## How to play ?
The game is published for test on windows now. you can [download](https://github.com/AmirMahdyJebreily/MinesweeperGO/releases/tag/win-test) the .exe or you can build your exe :
```shell
git clone https://github.com/AmirMahdyJebreily/MinesweeperGO.git && cd MinesweeperGO
```
then you have to download modules
```shell
go mod download   
```
Now you can build it for any environment you want 👍
```shell
GOOS=linux GOARCH=amd64 go build ./cmd/main 
```
enjoy the game

but in the future you can play in Unix and Windows environments (under the terminal).

## Contribute
- If you can provide fast and efficient ways to implement game logic, fork the repository and get started 👍
- The game needs to have different themes. Design separate themes for the game 😊