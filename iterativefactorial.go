package piscine

func IterativeFactorial(num int) int {
	if num < 0 || num > 100 {
		return 0
	}
	result := 1
	for r := 1; r <= num; r++ {
		result = result * r
	}
	return result
}
