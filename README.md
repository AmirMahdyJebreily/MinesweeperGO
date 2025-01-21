# Minesweeper GO  
A Go implementation of Minesweeper, currently available in terminal. In the future it will also compile for wasm and you can use it in your web projects


![Minesweeper Game](./docs/img/MinesweeperGo.webp)
## How to play ? 🖥️
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

## Why did I write this code? 🤔
You might ask, "What unemployed programmer would implement such an old and boring game in Go?" But I have to tell you that first, I am not unemployed, and second, this game is not boring. I develop and maintain this codebase for this two reasons.

First, many programmers in the 90s had fun making games that ran in terminal or DOS environments. Today, there are many engines for making games that make the programmer's job easier and allow the game developer to focus on the game itself, not on how to present a graphical environment on a particular game console.

But making a game for environments with fewer features (such as terminal operating systems or DOS) can help you better understand the details of a programming language and the flow of a program. You have to manage memory, and you also have to make your game run with great performance on most devices. These are challenges that would be boring if we were to raise them in a problem other than "gaming". As a result, I prepared a lib so that my students could both make a game and have a better understanding of the process of the program they were writing. It was an enjoyable project for them. They had to do something that would change the text of the page quickly by moving the keyboard so that the user would not notice that a whole string had changed and... which they achieved as a team effort. Later, they all thanked me for defining such a project for them. It was a little difficult for them, but as they said, "it was worth it."

Secondly, I myself wanted to check where exactly Go could be run. I used other languages ​​for a long time, and this means it is incompatible with different platforms!
.NET runs wherever you can install its runtime, but the .NET runtime is much heavier than Go and cannot be run on every operating system, and for me, Go is a gateway to other operating systems and environments. I'm trying to install old and new operating systems on virtual machines (probably this one I made myself), and run this program on them to see where exactly go can be. I also have a Nano Pi 2 Fire at home, which does the same thing as a Raspberry Pi, but with a bit less capabilities! I've also ported it to different SBCs and maybe even different Micro Controllers to see exactly how far I can push Go. (TinyGo is a compiler that I hope will help me!)
This page will be updated gradually and I'll upload a YouTube video reporting on this exciting zero!

## Why Minesweeper? 💣
Why not?! It's a logical and engaging game that always tells you that something must have left a mark or this number shouldn't be here! If you like challenging your brain, this game will really entertain you!

## How do I make one for myself? 🧑‍💻🧑‍💻
Use the `minesweeperlib` to create your own Minesweeper (as an exercise or to help students, etc.). It's provides just **5 functions** that allow you to create your Minesweeper in any platform (such as a browser):

#### **1. The `GetBoard(cols, rows)` function**  
```go 
func GetBoard(cols, rows int) *[][]int {
      // look at the source for more information: 
      //github.com/AmirMahdyJebreily/MinesweeperGO/blob/main/pkg/minesweeperlib/minesweeperlib.go#L9
}
```
This is their simple, it's get `cols`(columns) and `rows` count, and returns you a pointer to a 2D slice (a matrix) that will be used as a board.
> It would have been better if I had used the name "table" instead of board! We'll fix that later.
---
#### **2. The `GetRandomBombs(cols, rows, x0, y0, count)` function**
```go
func GetRandomBombs(cols, rows, x0, y0 int, count int) *[][2]int {
      // look at the source for more information: 
      // github.com/AmirMahdyJebreily/MinesweeperGO/blob/main/pkg/minesweeperlib/minesweeperlib.go#L19
}

```
This function selects random points on a "board" to a specified `count` and places bombs in them. The algorithm used in this function is that a random point is selected. If no other point is attached there, that point is selected.  
In practice, it is better that when the user clicks on a cell for the first time, that cell is not a bomb. As a result, we receive `y0` and `x0` as the **first clicked point**, and we do not allow too many bombs around it up to a specified radius.

---

#### **3. The `GetCellNumbers(board, bombs)` function**
```go 
func GetCellNumbers(board *[][]int, bombs *[][2]int) *[][]int{ 
	// look at the source for more information: 
	// github.com/AmirMahdyJebreily/MinesweeperGO/blob/main/pkg/minesweeperlib/minesweeperlib.go#L62
}
```
This one helps you figure out what number each cell has. It actually gives you another board where each cell has a number based on the bombs. If you read the source, you'll see that the idea behind the numbers in each cell is simple, which is that each bomb adds 1 to the number of the cells around it.

We could get the cell numbers when the user clicks. But in that case, we would have to guess which cells should be displayed at the same time. And that means if we want to get the numbers with each click, we would have to get the numbers of all the cells for each click to determine which ones to display. Well, we're not crazy! We get all the numbers once and each time we just decide which cells should be displayed.

---
>So far, we have the basics of a great Minesweeper game. Now it’s time to figure out what happens when the user clicks!
>
>Well, we all know exactly what happens (if you don’t know, read up on the game or try playing it)

---
#### **4. The `GetOpeneds(board, selectedPoint)` function**
```go
func GetOpeneds(board *[][]int, selected [2]int) [][2]int { 
      // look at the source for more information: 
      // github.com/AmirMahdyJebreily/MinesweeperGO/blob/main/pkg/minesweeperlib/minesweeperlib.go#L104
}
```
This one will search all the surrounding houses for you, and if you click on a house that has no bombs around it, it will show you all the houses around it, and if there is another house among them that is not in contact with any bombs, it will also show the houses around it.

---
#### **5. The `GetState(board, bombsCount, flaggeds, point)` function`**
```go
func GetState(board *[][]int, bombsCount int, flaggeds map[[2]int]bool, point [2]int) int{
      // look at the source for more information: 
      // github.com/AmirMahdyJebreily/MinesweeperGO/blob/main/pkg/minesweeperlib/minesweeperlib.go#L124
}
```
The last function determines the win/loss status. If you have flagged all the bombs, then you win!

---

## Why do I write its documentation in detail? 📝
Because maybe someone will want to help make this repo's mission come true!

## Contribute ☕
- If you can provide fast and efficient ways to implement game logic, fork the repository and get started 👍
- The game needs to have different themes. Design separate themes for the game 😊
- If you can run the game on a platform that you don't think is common, I'd be happy to put the code you wrote for that platform in the cmd directory and submit a Pull Request so I can merge it 📝