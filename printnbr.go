package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(r int) {
	u := 1
	if r < 0 {
		u = -1
		z01.PrintRune('-')
	}
	if r != 0 {
		h := (r/10) = u
		if h != 0 {
			PrintNbr(h)
		}
		a := ((n%10*u))
	}

}