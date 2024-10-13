package GameCore

type TCpl [2]int
type SearchableCouple = map[TCpl]struct{}

func Cpl(i, j int) TCpl {
	return [2]int{i, j}
}

func (zero *TCpl) AllNeighbors(max TCpl) SearchableCouple {
	var res = make(map[TCpl]struct{}, 8)
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			nearI, nearJ := zero[0]+i, zero[1]+j
			if nearJ >= 0 && nearJ < max[1] && nearI >= 0 && nearI < max[0] {
				res[Cpl(nearI, nearJ)] = struct{}{}
			}
		}
	}
	return res
}
