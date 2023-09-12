package main

import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	args := os.Args[1:]
	for r, u := range args {
		for j, h := range args {
			if u < h {
				args[r], args[j] = args[j], args[r]
			}
		}
	}
	for _, u := range args {
		for _, h := range u {
			z01.PrintRune(h)
			z01.PrintRune('\n')
		}
	}
}
