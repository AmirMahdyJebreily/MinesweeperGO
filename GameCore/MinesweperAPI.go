package GameCore

import (
	"math/rand/v2"
)

type game interface {
	Init() bool
	Play(any) bool // At every moment of the game, ply() is called, and if it returns True, the game continues, otherwise, the player loses.
}

type mineGame interface {
	game
	SetFlag(pos TCpl) bool
	Unlock(pos TCpl) struct{}
}

type MineSweeper struct {
	mineGame
	Size       TCpl
	Score      int
	BombsCount int
	Bombs      SearchableCouple
	Start      TCpl
}

func (m *MineSweeper) NumberOfPoint(pos TCpl) int {
	res := 0
	neighbors := pos.AllNeighbors(m.Size)
	for i := range neighbors {
		if _, ok := m.Bombs[i]; ok {
			res++
		}
	}
	return res
}

func (m *MineSweeper) NumberpointChain(pos TCpl, res map[TCpl]int, firstNumber int) map[TCpl]int {
	neighbors := pos.AllNeighbors(m.Size)
	if firstNumber == 0 {
		for i := range neighbors {
			thisNum := m.NumberOfPoint(i)
			if _, ok := m.Bombs[i]; !ok {
				if _, isExists := res[i]; isExists {
					res[i] = thisNum
				}
				if thisNum == 0 {
					m.NumberpointChain(i, res, firstNumber)
				}
			}
		}
	}

	res[pos] = firstNumber
	return res
}

func randomiseBombs(bombsCount int, size TCpl, start TCpl) SearchableCouple {
	bombs := make(map[TCpl]struct{}, bombsCount)
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
			for _, ok := bombs[Cpl(randomI, randomJ)]; ok || Cpl(randomI, randomJ) == start; _, ok = bombs[Cpl(randomI, randomJ)] {
				randomI, randomJ = rand.IntN(size[0]), rand.IntN(size[1])
			}
		}
		bombs[Cpl(randomI, randomJ)] = struct{}{}

	}
	return bombs
}

func InitRand(size, start TCpl, bombsCount int) *MineSweeper {

	return &MineSweeper{
		BombsCount: bombsCount,
		Size:       size,
		Start:      start,
		Bombs:      randomiseBombs(bombsCount, size, start),
	}
}

func AsBmbMap(bombs []TCpl) map[TCpl]struct{} {
	bombsMap := make(map[TCpl]struct{}, 0)
	for _, v := range bombs {
		bombsMap[v] = struct{}{}
	}
	return bombsMap
}

func Init(size, start TCpl, bombsCount int, bombs []TCpl) *MineSweeper {
	return &MineSweeper{
		BombsCount: bombsCount,
		Size:       size,
		Start:      start,
		Bombs:      AsBmbMap(bombs),
	}
}

func (m *MineSweeper) Play(arg bool) bool {
	return arg
}
