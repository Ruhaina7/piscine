package main

import (
	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	length := 0
	for index := range arg {
		length = index
	}
	for r := 1; r <= length; r++ {
		for _, u := range arg[u] {
			z01.PrintRune(u)
		}
		z01.PrintRune('\n')
	}
}
