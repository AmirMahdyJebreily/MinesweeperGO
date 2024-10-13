package main

import (
	"fmt"

	core "github.com/AmirMahdyJebreily/MinesweeperGO/GameCore"
	coreUtls "github.com/AmirMahdyJebreily/MinesweeperGO/GameCore/Utils"
)

func main() {
	g := core.InitRand(core.Cpl(8, 10), core.Cpl(5, 6), 11)

	fmt.Printf("%v", coreUtls.Sprintgridf(g))
	fmt.Printf("\n%#v", g.Bombs)
}
