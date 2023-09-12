package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	length := 0
	for index := range arg {
		length = index
	}
	for r := length; r >= 1; r-- {
		for _, u := range arg[r] {
			z01.PrintRune(u)
		}
		z01.PrintRune('\n')
	}
}
