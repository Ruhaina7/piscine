package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args
	for _, r := range args[0] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
