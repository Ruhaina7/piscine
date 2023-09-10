package piscine

func Fibonacci(ind int) int {
	if ind < 0 {
		return -1
	}
	if ind == 0 {
		return 0
	}
	if ind == 1 {
		return 1
	} else {
		return Fibonacci(ind-1) + Fibonacci(ind-2)
	}
}
