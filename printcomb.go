package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for r := '0'; r <= '7'; r++ {
		for u := '1'; u <= '8'; u++ {
			for n := '2'; n <= '9'; n++ {

				if r < u && u < n {

					z01.PrintRune(r)
					z01.PrintRune(u)
					z01.PrintRune(n)

					if r != '7' || u != '8' || n != '9' {

						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
