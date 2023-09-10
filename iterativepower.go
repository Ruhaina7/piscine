package piscine

func IterativePower(num int, power int) int {
	r := num
	if power < 0 {
		return 0
	} else {
		if power == 0 {
			return 1
		}
		for ru := power; ru > 1; ru-- {
			r = num * r
		}
		return r
	}
}
