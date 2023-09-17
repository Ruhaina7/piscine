package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(r []string) {
	for _, str1 := range r {
		for _, ch := range str1 {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
	
