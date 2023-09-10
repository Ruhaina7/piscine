package piscine

func RecursivePower(num int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		return num * RecursivePower(num, power-1)
	}
}
