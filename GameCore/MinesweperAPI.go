package GameCore

import (
	"math/rand/v2"
)

type game interface {
	Init() bool
	Play(any) bool // At every moment of the game, ply() is called, and if it returns True, the game continues, otherwise, the player loses.
}

type NeatCouple [2]int

type mineGame interface {
	game
	SetFlag(pos NeatCouple) bool
	Unlock(pos NeatCouple) struct{}
}

type MineSweeper struct {
	mineGame
	Size       NeatCouple
	Score      int
	BombsCount int
	Bombs      []NeatCouple
}

func (zero *NeatCouple) AllNeighbors(max NeatCouple) map[NeatCouple]struct{} {
	var res = make(map[NeatCouple]struct{}, 8)
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			nearI := zero[0] + i
			nearJ := zero[1] + j
			if nearJ >= 0 && nearJ < max[1] && nearI >= 0 && nearI < max[0] {
				res[[2]int{nearI, nearJ}] = struct{}{}
			}
		}
	}
	return res
}

func (m *MineSweeper) NumberOfPoint(pos NeatCouple) int {
	res := 0
	neighbors := pos.AllNeighbors(m.Size)
	for i := range neighbors {
		for _, j := range m.Bombs {
			if i == j {
				res++
			}
		}
	}
	return res
}

func randomiseBombs(bombsCount int, size NeatCouple) []NeatCouple {
	bombs := make(map[NeatCouple]struct{}, bombsCount)
	res := make([]NeatCouple, bombsCount)
	for i := 0; i < bombsCount; i++ {
		randomI, randomJ := rand.IntN(size[0]), rand.IntN(size[1])
		if i < 4 {
			switch i {
			case 0:
				randomI = 0
			case 1:
				randomJ = 0
			case 2:
				randomI = size[0] - 1
			case 3:
				randomJ = size[1] - 1
			}
		} else {
			for _, ok := bombs[[2]int{randomI, randomJ}]; ok; _, ok = bombs[[2]int{randomI, randomJ}] {
				randomI, randomJ = rand.IntN(size[0]), rand.IntN(size[1])
			}
		}
		bombs[[2]int{randomI, randomJ}] = struct{}{}
		res[i] = [2]int{randomI, randomJ}
	}
	return res
}

func Init(size NeatCouple, bombsCount int) MineSweeper {
	m := MineSweeper{}
	m.BombsCount = bombsCount
	m.Size = size
	m.Bombs = randomiseBombs(m.BombsCount, m.Size)

	return m
}

func (m *MineSweeper) Play(arg bool) bool {
	return arg
}
