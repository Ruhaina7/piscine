package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	p := &point{}
	setPoint(p)
	s1 := "x = "
	s2 := ", y = "
	for _, r := range s1 {
		z01.PrintRune(r)
	}
	printnbr(52)
	printnbr(50)
	for _, r := range s2 {
		z01.PrintRune(r)
	}
	printnbr(50)
	printnbr(49)
	z01.PrintRune('\n')
}
func printnbr(x rune) {
	z01.PrintRune(x)
}
