package main

import "github.com/01-edu/z01"

func printalphabet() {
	for r := 'a'; r <= 'z'; r++ {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
