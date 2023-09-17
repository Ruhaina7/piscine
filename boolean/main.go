package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if isEven(%nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	if isEven(len(os.Arg)) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
