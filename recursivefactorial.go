package piscine

func RecursiveFactorial(num int) int {
	if num == 0 {
		return 1
	}
	if num < 0 || num > 100 {
		return 0
	} else {
		return num * RecursiveFactorial(num-1)
	}
}
