package piscine

func AppendRange(min, max int) []int {
	if max <= min {
		return nil
	}
	var result1 []int
	for r := min; r < max; r++ {
		result1 = append(result1, r)
	}
	return result1
}
