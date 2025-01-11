package main

import (
	"fmt"
	"github.com/AmirMahdyJebreily/MinesweeperGO/pkg/minesweeperlib"
	"github.com/eiannone/keyboard"
	"slices"
	"strings"
)

const unoppend = "\u001B[38;5;242m∙\u001B[1;0m"
const zero = "\u001B[38;5;248m■\u001B[1;0m"
const bomb = "\u001b[101m\u001B[31mX\u001B[0m"
const flag = "\u001b[102m\u001B[36mF\u001B[0m"

// use \x1b[1;31m ansi escape code for colorise
func coloriseNumber(n int) string {
	if n == 0 {
		return zero
	}
	return fmt.Sprintf("\x1b[1;3%vm%v\x1b[0m", n, n)
}

func SprintCell(data string, selected bool) string {
	if selected {
		return fmt.Sprintf("\u001B[5m[\u001B[25m%v\u001B[5m]\x1b[25m", data)
	}

	if data == flag {
		return fmt.Sprintf("\u001B[102m\u001B[36m[F]\u001B[0m")
	}
	if data == bomb {
		return fmt.Sprintf("\u001B[101m\u001B[31m[x]\u001B[0m")
	}

	return fmt.Sprintf(" %v ", data)
}

func Sprintgridf(board *[][]int, bombsCount int, flagged *map[[2]int]bool, oppend *[][2]int, selected [2]int, messages string) *strings.Builder {
	rows, cols := len(*board), len((*board)[0])
	var res strings.Builder
	res.WriteString("\033[H\033[J") // Clear the old screen before print new board
	res.WriteString(fmt.Sprintf(" [Size: %v×%v] [Bombs: %v] [Flags: %v]\n", cols, rows, bombsCount, bombsCount-len(*flagged)))
	for i := rows - 1; i >= 0; i-- {
	lines:
		for j := 0; j < cols; j++ {
			isSelected := [2]int{j, i} == selected

			if _, isflag := (*flagged)[[2]int{j, i}]; isflag {
				res.WriteString(SprintCell(flag, isSelected))
				continue lines
			}

			if (oppend) != nil {
				if !slices.Contains(*oppend, [2]int{j, i}) {
					res.WriteString(SprintCell(unoppend, isSelected))
					continue lines
				}
			}

			if (*board)[i][j] == -1 {
				res.WriteString(SprintCell(bomb, isSelected))
				continue lines
			}

			res.WriteString(SprintCell(coloriseNumber((*board)[i][j]), isSelected))
		}
		res.WriteString("\n")
	}
	res.WriteString(fmt.Sprintf("\n %v", messages))
	res.WriteString(fmt.Sprintf("\n .:: \x1b[1;34m[Arrows: Move] [O & Enter: Open Cell] [F: Flag] [Q & ESC: Quit]\x1b[1;0m"))

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
		suggested := int(float64(cols*rows)*0.21) - 1
		fmt.Printf("Enter The count of bombs: \u001b[s\n\u001b[90m(%v) bombs is recommended, Press the Enter for it ;)\u001B[1;0m\u001B[u ", suggested)
		_, scanError := fmt.Scanln(&bombsCount)
		if scanError != nil {
			bombsCount = suggested
			break
		}
		break
	}
	fmt.Println("\n\u001B[?1049h\u001B[H\u001B[J\u001B[?25l") // save screen
	board := minesweeperlib.GetBoard(cols, rows)
	flaggeds := make(map[[2]int]bool, bombsCount)
	selected := [2]int{cols / 2, rows / 2}
	fmt.Println((*Sprintgridf(board, bombsCount, &flaggeds, nil, selected, "select a cell to start")).String())
	var x0, y0 int
	var bombs *[][2]int = nil
	var oppend [][2]int = nil
	inGame := true
	for inGame {
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

		if char == 'q' || char == 'Q' || key == keyboard.KeyEsc {
			fmt.Print("\u001B[?1049l\u001B[?25h")
			break
		}

		message := ""

		if char == 'f' || char == 'F' || key == keyboard.KeySpace {
			if val, isflag := flaggeds[selected]; isflag || val {
				delete(flaggeds, selected)
			} else if len(flaggeds) < bombsCount {
				flaggeds[selected] = true
			}
		}
		if char == 'o' || char == 'O' || key == keyboard.KeyEnter {
			if bombs == nil {
				x0, y0 = selected[0], selected[1]
				oppend = make([][2]int, 0)
				bombs = minesweeperlib.GetRandomBombs(cols, rows, x0, y0, bombsCount)
				board = minesweeperlib.GetCellNumbers(board, bombs)
			}
			oppend = slices.Concat(oppend, minesweeperlib.GetOpeneds(board, selected))
		}

		state := minesweeperlib.GetState(board, bombsCount, flaggeds, selected)
		if state != 0 { // winner or loser
			if state == 1 {
				message = "\u001b[32mYou Win :)\u001b[1;0m"
				inGame = false
				fmt.Println((*Sprintgridf(board, bombsCount, &flaggeds, nil, selected, message)).String())
				fmt.Println("Press something to exit\n")
				fmt.Scanln()
				fmt.Print("\u001B[?1049l\u001B[?25h")
				break
			} else if state == 2 {
				message = "\u001b[31mGame Over :(\u001b[1;0m"
				inGame = false
				fmt.Println((*Sprintgridf(board, bombsCount, &flaggeds, nil, selected, message)).String())
				fmt.Println("Press something to exit\n")
				fmt.Scanln()
				fmt.Print("\u001B[?1049l\u001B[?25h")
				break
			}
		}
		// update screen
		fmt.Println((*Sprintgridf(board, bombsCount, &flaggeds, &oppend, selected, message)).String())
		fmt.Print()
	}
}
