package piscine


func Iterativefactorial(num int) int {
	if num < 0 || num > 100 {
		return 0
	}
	res := 1
	for r := 1; r <= num; r++ {
		res = res * r
	}
	return res
}
