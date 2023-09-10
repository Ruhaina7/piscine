package piscine

func Sqrt(num int) int {
	if num == 1 {
		return 1
	}
	if num == 2 {
		return 0
	}
	if num > 0 {
		result := 1
		sqrtt := 0
		for r := 1; result <= num; r++ {
			result = r * r
			sqrtt++
			if result == num {
				return sqrtt
			}
		}
		return 0
	} else {
		return 0
	}
}
