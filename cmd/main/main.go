package main

import (
	"fmt"
	"github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"github.com/eiannone/keyboard"
	"slices"
	"strings"
)

const unoppend = "∙"
const zero = "O"
const bomb = "\u001b[41m\u001B[35mX\u001B[0m"

// use \x1b[1;31m ansi escape code for colorise
func coloriseNumber(n int) string {
	if n == 0 {
		return zero
	}
	return fmt.Sprintf("\x1b[1;3%vm%v\x1b[0m", n, n)
}

func SprintCell(data string, selected bool) string {
	if selected {
		return fmt.Sprintf("\u001B[5m[\u001B[25m%v\u001B[5m]\x1b[25m  ", data)
	}
	return fmt.Sprintf(" %v   ", data)
}

func Sprintgridf(board *[][]int, bombs *[][2]int, oppend *[][2]int, selected [2]int) *strings.Builder {
	rows, cols := len(*board), len((*board)[0])
	var res strings.Builder
	res.WriteString("\033[H\033[J") // Clear the old screen before print new board
	if bombs != nil {
		res.WriteString(fmt.Sprintf(" .:: [Size: %v×%v] [Bombs: %v]\n", cols, rows, len(*bombs)))
	}
	for i := rows - 1; i >= 0; i-- {
		res.WriteString(fmt.Sprintf(" %2d     ", i+1))
	lines:
		for j := 0; j < cols; j++ {
			isSelected := [2]int{j, i} == selected
			if oppend == nil {
				res.WriteString(SprintCell(unoppend, isSelected))
				continue lines
			}

			if !slices.Contains(*oppend, [2]int{j, i}) {
				res.WriteString(SprintCell(unoppend, isSelected))
				continue lines
			}

			if (*board)[i][j] == -1 {
				res.WriteString(SprintCell(bomb, isSelected))
				continue lines
			}

			res.WriteString(SprintCell(coloriseNumber((*board)[i][j]), isSelected))

		}
		res.WriteString("\n")
	}
	res.WriteString(fmt.Sprintf("\n      "))
	for j := 0; j < cols; j++ {
		res.WriteString(fmt.Sprintf("  %2d ", j+1))
	}
	res.WriteString(fmt.Sprintf("\n\n .:: \x1b[1;34m[Arrows: Move] [O: Open Cell] [F: Flag]\x1b[1;0m"))

	return &res
}

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var cols, rows int
	fmt.Println("Wellcome to CodeAgha's MineSweeper Game in terminal")
	for {
		fmt.Print("Enter The Columns,Rows: ")
		_, scanError := fmt.Scanf("%v,%v\n", &cols, &rows)
		if scanError != nil {
			fmt.Println("Please Input values in format: Columns,Rows")
			continue
		}
		break
	}
	var bombsCount int
	for {
		fmt.Print("Enter The count of bombs: ")
		_, scanError := fmt.Scanln(&bombsCount)
		if scanError != nil {
			fmt.Println("Please Input a number")
			continue
		}
		break
	}

	board := minesweeperlib.GetBoard(cols, rows)
	selected := [2]int{cols / 2, rows / 2}
	fmt.Println((*Sprintgridf(board, nil, nil, selected)).String())
	var x0, y0 int
	var bombs *[][2]int = nil
	var oppend [][2]int = nil

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyArrowLeft && selected[0] > 0 {
			selected[0]--
		}

		if key == keyboard.KeyArrowRight && selected[0] < cols-1 {
			selected[0]++
		}

		if key == keyboard.KeyArrowUp && selected[1] < rows-1 {
			selected[1]++
		}

		if key == keyboard.KeyArrowDown && selected[1] > 0 {
			selected[1]--
		}

		if (char == 'o' || char == 'O' || key == keyboard.KeyEnter) && selected[0] < cols {
			if bombs == nil {
				x0, y0 = selected[0], selected[1]
				oppend = make([][2]int, 0)
				bombs = minesweeperlib.GetRandomBombs(cols, rows, x0-1, y0-1, bombsCount)
				board = minesweeperlib.GetCellNumbers(board, bombs)
			}
			oppend = append((oppend), selected)
		}

		// update screen
		fmt.Println((*Sprintgridf(board, bombs, &oppend, selected)).String())
		fmt.Print()

	}
}
