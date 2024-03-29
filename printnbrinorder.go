package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(r int) {
	if r < 0 {
		return
	}
	if r == 0 {
		z01.PrintRune('0')
		return
	}
	digitt := make([]int, 0)
	for r > 0 {
		digit := r % 10
		digitt = append(digitt, digit)
		r /= 10
	}
	for u := 0; u < 10; u++ {
		for _, digit := range digitt {
			if digit == u {
				z01.PrintRune(rune('0' + digit))
			}
		}
	}
}
