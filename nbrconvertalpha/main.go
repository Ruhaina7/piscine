package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	x := 0
	for _, cifra := range s {
		if '0' <= cifra && cifra <= '9' {
			z := 0
			for k := '1'; k <= cifra; k++ {
				z++
			}
			x = x*10 + z
		} else {
			x = -1
			break
		}
	}
	if (x != -1) && !(1 <= x && x <= 26) {
		x = -1
	}
	return x
}

func baseAtoi(s string) int { // Ascii To Integer
	// '0' -> 48
	// '0' - '0' -> 0
	// 48 - 48 -> 0
	// '1' - '0' -> 1
	// 49 - 48 -> 1
	// to remove a 0 from a number divide by 10
	// y3ni, 100 / 10 = 10
	// divide result, modulo result
	// 100 % 10 = 0
	// 
}

func main() {
	arg := os.Args[1:]
	flagupper := false
	c := 0
	for range arg {
		c++
	}
	if c >= 2 && arg[1] == "--upper" {
		flagupper = true
	}
	for _, k := range arg {
		num := BasicAtoi(k)
		if num == -1 {
			z01.PrintRune(' ')
		} else {
			if !flagupper {
				z01.PrintRune(rune('a' + num - 1))
			} else {
				z01.PrintRune(rune('A' + num - 1))
			}
		}
	}
	z01.PrintRune(10)
}
