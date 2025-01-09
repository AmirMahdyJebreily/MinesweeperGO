/*
The terminal interface for minesweeper pkg
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	gridSize     = 20 // اندازه ماتریس
	numParticles = 50 // تعداد ذرات
)

var grid [gridSize][gridSize]int // محیط ماتریس

// بررسی اینکه آیا نقطه خالی است
func isFree(x, y int) bool {
	return x >= 0 && x < gridSize && y >= 0 && y < gridSize && grid[x][y] == 0
}

// اتصال ذره به شبکه
func attachParticle(x, y int) {
	grid[x][y] = 1
}

// تولید ذرات و پراکنده کردن آنها
func randomWalk(x, y int) (int, int) {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // حرکت به اطراف
	for {
		dir := directions[rand.Intn(len(directions))]
		newX, newY := x+dir[0], y+dir[1]
		if !isFree(newX, newY) {
			return newX, newY
		}
		x, y = newX, newY
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// قرار دادن اولین نقطه (seed)
	center := gridSize / 2
	grid[center][center] = 1

	for i := 0; i < numParticles; i++ {
		// شروع از یک نقطه تصادفی
		x, y := rand.Intn(gridSize), rand.Intn(gridSize)
		x, y = randomWalk(x, y)
		attachParticle(x, y)
	}

	// نمایش ماتریس
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("* ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
