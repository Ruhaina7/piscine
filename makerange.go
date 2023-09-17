package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	var result2 []int = make([]int, max-min)
	for r, u := min, 0; r < max; r++ {
		result2[u] = r
		u++
	}
}
